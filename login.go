package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/raad-dego/chirpy/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) handlerLogIn(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
		ExpiresSec int `json:"expires_in_sec"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	email := strings.ToLower(params.Email)
	dbUser, err := cfg.DB.GetUser(email)
	if err != nil && err != database.ErrNotExist {
		respondWithError(w, http.StatusUnauthorized, "Email not found")
		return
	}
	// hashedPasswordAttempt, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid password format.")
	// }

	err = bcrypt.CompareHashAndPassword(dbUser.HashedPassword, []byte(params.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not log in")
		return
	}
	respondWithJSON(w, http.StatusOK, dbUser)
}
