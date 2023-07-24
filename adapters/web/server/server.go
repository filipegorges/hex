package web

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/filipegorges/hex/adapters/web/handler"
	"github.com/filipegorges/hex/application"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w *WebServer) SetService(service application.ProductServiceInterface) {
	w.Service = service
}

func (w WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "LOG: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
