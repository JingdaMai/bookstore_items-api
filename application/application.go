package application

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var router = mux.NewRouter()

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8082",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
