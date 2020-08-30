package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8082",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
