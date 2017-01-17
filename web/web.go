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
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

//var jsTemplate = template.Must(template.ParseFiles("www/js/main.js"))

func serveMainjs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/js/main.js" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/javascripthtml; charset=utf-8")
	jsTemplate := template.Must(template.ParseFiles("www/js/main.js"))
	jsTemplate.Execute(w, r.Host)
}

// Start initializes the webserver and the server receving the lines
func Start() {
	fmt.Printf("Binding definition provided: %s\n", viper.GetString("web.bind"))
	fmt.Printf("Serving at: %s\n", viper.GetString("web.serve"))

	srv := &server{}
	go srv.run()

	r := mux.NewRouter()
	r.HandleFunc("/js/main.js", serveMainjs) // js template
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(srv.hub, w, r)
	})
	r.PathPrefix("/").Handler(handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir("./www")))) // static files
	http.Handle("/", r)
	if err := http.ListenAndServe(viper.GetString("web.serve"), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}