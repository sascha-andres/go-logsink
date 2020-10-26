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
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/rakyll/statik/fs"

	_ "github.com/sascha-andres/go-logsink/v2/web/statik" // get access to data
)

var (
	statikFS      http.FileSystem
	jsTemplate    *template.Template
	numberOfLines = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "log_lines",
		Help: "Number of lines received",
	})
)

// Start initializes the webserver and the server receving the lines
func Start() {
	log.Printf("Binding definition provided: %s", viper.GetString("web.bind"))
	log.Printf("Serving at: %s", viper.GetString("web.serve"))
	webDir := viper.GetString("web.directory")
	log.Printf("Directory: %s", webDir)

	srv := &server{}
	go srv.run()

	prometheus.MustRegister(numberOfLines)

	r := mux.NewRouter()
	r.HandleFunc("/api/go-logsink/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(srv.hub, w, r)
	})
	r.Handle("/metrics", promhttp.Handler())
	if "" == webDir {
		log.Print("serving static files from binary")                                                    // js template
		r.PathPrefix("/").Handler(handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(statikFS))) // static files
	} else {
		log.Printf("serving static files from %s", webDir)
		r.PathPrefix("/").Handler(handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(webDir)))) // static files
	}
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

func init() {
	files, err := fs.New()
	if err != nil {
		log.Fatal("Error initializing filesystem: ", err)
	}
	statikFS = files
}
