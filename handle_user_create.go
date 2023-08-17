package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/raad-dego/chirpy/internal/auth"
	"github.com/raad-dego/chirpy/internal/database"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	email := strings.ToLower(params.Email)
	_, err = cfg.DB.GetUserByEmail(email)
	if err != nil && err != database.ErrNotExist {
		respondWithError(w, http.StatusBadRequest, "Email already used.")
		return
	}
	hashedPassword, err := auth.HashPassword(params.Password)

	user, err := cfg.DB.CreateUser(email, hashedPassword)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, User{
		ID:    user.ID,
		Email: user.Email,
	})
}
