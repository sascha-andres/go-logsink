package web

import (
	"context"
	"github.com/arl/statsviz"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//Start initializes the webserver and the server receving the lines
func Start() {
	logrus.Printf("Serving at: %s", viper.GetString("web.serve"))

	pipe := make(chan *logsink.LineMessage)
	srv := &server{
		hub: newHub(),
	}
	go srv.processOutput(pipe)
	go srv.run(pipe)

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
			logrus.Fatal("ListenAndServe: ", err)
		}
	}()
	<-stop

	logrus.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := h.Shutdown(ctx)
	if err != nil {
		logrus.Printf("error shutting down gracefully: %s", err)
	} else {
		logrus.Println("Server gracefully stopped")
	}
}
