package grpcserver

import (
	"github.com/arl/statsviz"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func startDebug() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/debug/statsviz/ws").Name("GET /debug/statsviz/ws").HandlerFunc(statsviz.Ws)
	r.Methods("GET").PathPrefix("/debug/statsviz/").Name("GET /debug/statsviz/").Handler(statsviz.Index)

	h := &http.Server{Addr: viper.GetString("web.serve"), Handler: handlers.CORS()(handlers.ProxyHeaders(r))}
	if err := h.ListenAndServe(); err != nil {
		logrus.Fatal("ListenAndServe: ", err)
	}
}
