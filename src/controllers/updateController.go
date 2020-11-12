package controllers

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"services"
	"strings"
	"structs/shared"
)

func UpdateController(w http.ResponseWriter, r *http.Request) {
	//handle preflight
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Vary", "Origin")
	w.Header().Add("Vary", "Access-Control-Request-Method")
	w.Header().Add("Vary", "Access-Control-Request-Headers")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "POST,OPTIONS")

	if r.Method == "OPTIONS" {
		return
	}

	_, p, ok := r.BasicAuth()
	if !ok || len(strings.TrimSpace(p)) < 1 {
		http.Error(w, "CachetHQ Token required.", http.StatusUnauthorized)
		return
	}

	var request shared.UpdateRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	code, msg := services.UpdateService(p, request)
	if code != 0 {
		http.Error(w, msg, code)
	}
}