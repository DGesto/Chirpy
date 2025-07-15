package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	type chirpBody struct {
		Body string `json:"body"`
	}

	type clearedChirp struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	chirp := chirpBody{}
	err := decoder.Decode(&chirp)
	if err != nil {
		msg := "Something went wrong"
		respondWithError(w, 400, msg)
		return
	}

	if len(chirp.Body) > 140 {
		msg := "Chirp is too long"
		respondWithError(w, 400, msg)
		return
	}

	cleanedChirp := cleanChirp(chirp.Body)
	respondWithJSON(w, 200, clearedChirp{cleanedChirp})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorResponse struct {
		ErrorResp string `json:"error"`
	}

	errResp := errorResponse{msg}
	dat, err := json.Marshal(errResp)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func cleanChirp(chirp string) string {
	words := strings.Split(chirp, " ")
	badWords := []string{"kerfuffle", "sharbert", "fornax"}

	for idx, word := range words {
		if slices.Contains(badWords, strings.ToLower(word)) {
			words[idx] = "****"
		}
	}

	return strings.Join(words, " ")
}
