package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/raad-dego/chirpy/internal/auth"
	"github.com/raad-dego/chirpy/internal/database"
)

func (cfg *apiConfig) handlerRedVerification(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID int `json:"user_id"`
		} `json:"data"`
	}

	requestApi, err := auth.GetPolkaApi(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't retrieve Polka API")
		return
	}
	if requestApi != cfg.polkaApi {
		respondWithError(w, http.StatusUnauthorized, "Invalid Polka API")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	if params.Event != "user.upgraded" {
		respondWithError(w, http.StatusOK, "Wrong event")
		return
	}

	err = cfg.DB.MarkUserAsRedVerified(params.Data.UserID)
	if err != nil {
		if errors.Is(err, database.ErrNotExist) {
			respondWithError(w, http.StatusNotFound, "Couldn't find user")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user")
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
