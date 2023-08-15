package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerUserRetrieve(email string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbUser, err := cfg.DB.GetUser(email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve users")
			return
		}
		respondWithJSON(w, http.StatusOK, dbUser)
	}

}
