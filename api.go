package main

import (
	"encoding/json"
	"net/http"
)

func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err :=  f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJson(w, e.Status, e)
				return
			}
			writeJson(w, http.StatusInternalServerError, apiError{ Err: "internal server error", Status: http.StatusInternalServerError })
		}
	}
}

func handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{ Err: "invalid method", Status: http.StatusMethodNotAllowed}; 
	}

	user := User {}
	if !user.Valid{
		return apiError{ Err: "User is not logged in", Status: http.StatusUnauthorized}
	}

	return writeJson(w, http.StatusOK, User{})
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}