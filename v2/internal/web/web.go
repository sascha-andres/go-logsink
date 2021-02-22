// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"context"
	"embed"
	"github.com/arl/statsviz"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"
	"io/fs"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

var (
	//go:embed dist
	embededFiles  embed.FS
	jsTemplate    *template.Template
	numberOfLines = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "log_lines",
		Help: "Number of lines received",
	})
)

func getFileSystem() http.FileSystem {
	webDir := viper.GetString("web.directory")
	log.Printf("Directory: %s", webDir)

	if "" != webDir {
		log.Printf("serving static files from %s", webDir)
		log.Print("using live mode")
		return http.Dir(webDir)
	}

	log.Print("serving static files from binary")
	fsys, err := fs.Sub(embededFiles, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

// Start initializes the webserver and the server receving the lines
func Start() {
	log.Printf("Binding definition provided: %s", viper.GetString("web.bind"))
	log.Printf("Serving at: %s", viper.GetString("web.serve"))

	srv := &server{}
	go srv.run()

	prometheus.MustRegister(numberOfLines)

	r := mux.NewRouter()
	r.HandleFunc("/api/go-logsink/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(srv.hub, w, r)
	})
	if viper.GetBool("debug") {
		r.Methods("GET").Path("/debug/statsviz/ws").Name("GET /debug/statsviz/ws").HandlerFunc(statsviz.Ws)
		r.Methods("GET").PathPrefix("/debug/statsviz/").Name("GET /debug/statsviz/").Handler(statsviz.Index)
	}
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix("/").Handler(handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(getFileSystem())))
	http.Handle("/", r)
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	h := &http.Server{Addr: viper.GetString("web.serve"), Handler: handlers.CORS()(handlers.ProxyHeaders(r))}
	go func() {
		if err := h.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	<-stop

	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := h.Shutdown(ctx)
	if err != nil {
		log.Printf("error shutting down gracefully: %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
