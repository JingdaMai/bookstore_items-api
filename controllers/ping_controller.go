package controllers

import "net/http"

const pong = "pong"

var PingController pingControllerInterface = &pingController{}

type pingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(pong))
	if err != nil {
		panic(err)
	}
}
