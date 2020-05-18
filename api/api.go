package api

import (
	"log"
	"net/http"
	"strconv"
)

type Api struct {
	logger *log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Api {
	return &Api{logger: logger}
}

func (api *Api) Start(port int) error {
	p := strconv.Itoa(port)
	s := &http.Server{
		Addr:    ":" + p,
		Handler: api.bootRouter(),
	}
	api.server = s

	return s.ListenAndServe()
}