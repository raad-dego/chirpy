package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDString := chi.URLParam(r, "chirpID")
	chirpID, err := strconv.Atoi(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID")
		return
	}

	dbChirp, err := cfg.DB.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:       dbChirp.ID,
		Body:     dbChirp.Body,
		AuthorId: dbChirp.AuthorId,
	})
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	authorID := -1
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = strconv.Atoi(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID")
			return
		}
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		if authorID != -1 && dbChirp.AuthorId != authorID {
			continue
		}

		chirps = append(chirps, Chirp{
			ID:       dbChirp.ID,
			Body:     dbChirp.Body,
			AuthorId: dbChirp.AuthorId,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

// NOT NEEDED - still works
// func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
// 	authorIDString := r.URL.Query().Get("author_id")
// 	if len(authorIDString) != 0 {
// 		authorID, err := strconv.Atoi(authorIDString)
// 		if err != nil {
// 			respondWithError(w, http.StatusBadRequest, "Invalid author ID")
// 			return
// 		}
// 		chirpsByAuthor, err := cfg.DB.GetChirpsByAuthor(authorID)
// 		if err != nil {
// 			respondWithError(w, http.StatusInternalServerError, "Invalid author")
// 			return
// 		} else if len(chirpsByAuthor) == 0 {
// 			respondWithError(w, http.StatusNoContent, "No Chirps yet")
// 			return
// 		}
// 		sort.Slice(chirpsByAuthor, func(i, j int) bool {
// 			return chirpsByAuthor[i].ID < chirpsByAuthor[j].ID
// 		})
// 		respondWithJSON(w, http.StatusOK, chirpsByAuthor)
// 		return
// 	}

// 	dbChirps, err := cfg.DB.GetChirps()
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
// 		return
// 	}

// 	chirps := []Chirp{}
// 	for _, dbChirp := range dbChirps {
// 		chirps = append(chirps, Chirp{
// 			ID:       dbChirp.ID,
// 			Body:     dbChirp.Body,
// 			AuthorId: dbChirp.AuthorId,
// 		})
// 	}

// 	sort.Slice(chirps, func(i, j int) bool {
// 		return chirps[i].ID < chirps[j].ID
// 	})

// 	respondWithJSON(w, http.StatusOK, chirps)
// }
