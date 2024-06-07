package server

import (
	"log"
	"net/http"

	"github.com/guirialli/dolar-api/routers"
)

func ServerAndListen() {
	mux := http.NewServeMux()
	routers.AddDefaultRouters(mux)

	log.Println("Server listening on http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
