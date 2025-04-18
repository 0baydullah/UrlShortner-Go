package handlers

import (
	"UrlShortner/database"
	"UrlShortner/models"
	"UrlShortner/utils"
	"encoding/json"
	"net/http"
)

const baseUrl = "https://loacalhost:8080/"

func ShortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.UrlReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	key := utils.GenerateKey(6)
	err = database.SaveUrl(key, req.URL)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := models.UrlRes{
		Key:      key,
		Url:      req.URL,
		ShortUrl: baseUrl + key,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	if key == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	originalUrl, err := database.GetUrl(key)
	if err != nil {
		http.Error(w, "Url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusFound)
}
