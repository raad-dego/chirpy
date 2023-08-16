package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/raad-dego/chirpy/internal/auth"
	"github.com/raad-dego/chirpy/internal/database"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrSignatureInvalid = errors.New("signature is invalid")
)

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}
	type response struct {
		User
		Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	email := strings.ToLower(params.Email)
	user, err := cfg.DB.GetUserByEmail(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	if err != nil && err != database.ErrNotExist {
		respondWithError(w, http.StatusUnauthorized, "Email not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(params.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not log in")
		return
	}

	if params.ExpiresInSeconds > 86400 || params.ExpiresInSeconds == 0 {
		params.ExpiresInSeconds = 60 * 60 * 24 // 24 hours in seconds
	}

	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Duration(params.ExpiresInSeconds)*time.Second)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create JWT")
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:    user.ID,
			Email: user.Email,
		},
		Token: token,
	})
}
