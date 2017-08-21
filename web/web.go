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
	"bytes"
	"context"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"

	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/rakyll/statik/fs"

	_ "github.com/sascha-andres/go-logsink/web/statik" // get access to data
)

var (
	statikFS      http.FileSystem
	jsTemplate    *template.Template
	numberOfLines = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "log_lines",
		Help: "Number of lines received",
	})
)

type templateData struct {
	Host   string
	Limit  int32
	Scheme string
}

func serveMainjs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/js/main.js" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	jsTemplate.Execute(w, templateData{Host: r.Host, Limit: int32(viper.GetInt("web.limit")), Scheme: getScheme(r)})
}

func getScheme(r *http.Request) string {
	var scheme string
	if len(r.Header["Referer"]) == 0 {
		scheme = "ws"
	} else {
		if strings.HasPrefix(r.Header["Referer"][0], "https") {
			scheme = "wss"
		} else {
			scheme = "ws"
		}
	}
	return scheme
}

// Start initializes the webserver and the server receving the lines
func Start() {
	log.Printf("Binding definition provided: %s", viper.GetString("web.bind"))
	log.Printf("Serving at: %s", viper.GetString("web.serve"))
	log.Printf("Line limit: %d", viper.GetInt("web.limit"))

	srv := &server{}
	go srv.run()

	prometheus.MustRegister(numberOfLines)

	r := mux.NewRouter()
	r.HandleFunc("/js/main.js", serveMainjs) // js template
	r.HandleFunc("/api/go-logsink/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(srv.hub, w, r)
	})
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix("/").Handler(handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(statikFS))) // static files
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
	h.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}

func init() {
	files, err := fs.New()
	if err != nil {
		log.Fatal("Error initializing filesystem: ", err)
	}
	statikFS = files
	file, err := statikFS.Open("/js/main.js")
	if err != nil {
		log.Fatal("Error reading js template: ", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	tmpl, err := template.New("jsTemplate").Parse(buf.String())
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}
	jsTemplate = tmpl
}
