package api

import (
	"net/http"
	"regexp"
)

func (api *Api) bootRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.router)
	return mux
}

func (api *Api) router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	api.logger.Printf("Request to path: %s", path)

	convertRoute, err := regexp.Compile("/convert/?.*")
	if err != nil {
		api.logger.Fatal("Failure compiling convert route")
	}

	switch r.Method {
	case http.MethodGet:
		switch {
		case path == "/":
			api.handleStatus(w, r)
		case convertRoute.MatchString(path):
			api.handleConversion(w, r)
		default:
			api.jsonResponse(w, "Not found, please check and try again.", nil, http.StatusNotFound)
		}
	default:
		api.jsonResponse(w, "Not found, please check and try again.", nil, http.StatusNotFound)
	}
}
