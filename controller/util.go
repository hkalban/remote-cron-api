package controller

import "net/http"

func setSuccessHeaders(w http.ResponseWriter) {
	setDefaultHeaders(w)
	w.WriteHeader(http.StatusOK)
}

func setAcceptedHeaders(w http.ResponseWriter) {
	setDefaultHeaders(w)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte{})
}

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
