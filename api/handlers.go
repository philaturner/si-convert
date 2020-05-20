package api

import (
	"encoding/json"
	"net/http"
	unit "si-convert/convert"
	"strconv"
	"strings"
)

type response struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Status int `json:"status"`
}

func (api *Api) jsonResponse(w http.ResponseWriter, msg string, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;utf-8")

	b, err := json.Marshal(response{Message: msg, Data: data, Status: status})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"message": "internal error", "status": "500"}`))
		api.logger.Fatalf("Failed to marshal JSON: %v", err)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		api.logger.Fatalf("Failed to marshal JSON: %v", err)
	}
}

func (api *Api) handleStatus(w http.ResponseWriter, r *http.Request) {
	message := "All systems are operational"
	api.jsonResponse(w, message, nil, http.StatusOK)
}

func (api *Api) handleConversion(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	v, err := strconv.ParseFloat(parts[4], 64)
	siUnit := unit.SI{Value: v, Option: parts[2]}
	res, err := siUnit.Convert(parts[3])
	if err != nil {
		api.jsonResponse(w, "We don't support this conversion", nil, http.StatusBadRequest)
		return
	}
	api.jsonResponse(w, "Conversion successful", res, http.StatusOK)
}