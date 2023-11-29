package main

import "net/http"

func readinessEndpoint(w http.ResponseWriter, r *http.Request) {
	type statusResponse struct {
		Status string `json:"status"`
	}
	responseWithJSON(w, http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func errorEndpoint(w http.ResponseWriter, r *http.Request) {
	type statusResponse struct {
		Error string `json:"error"`
	}

	responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
}