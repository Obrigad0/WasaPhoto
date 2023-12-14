package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/Obrigad0/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

// getComments() ritorna tutta la lista dei commenti di un'immagine
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	iIdint, _ := strconv.Atoi(ps.ByName("imageId")) // id foto

	if !verificaLogin(r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	// tutti possono vedere i commenti di una foto

	// prelevo il token dell' user
	token := estrazioneToken(r.Header.Get("Authorization"))

	if uIdint != token {
		// se la persona che richiede i commenti NON e' il proprietario del profilo
		// verifico che l'Utente A (token) non sia stato bannato dall'Utente B (uId) proprietario foto
		// e viceversa (non puoi vedere i commenti  di una persona che hai bannato)
		// controllo se A e' stato bannato da B
		result, err := rt.db.IsBanned(User{UId: uIdint}.ToDatabase(), User{UId: token}.ToDatabase()) // A e' stato bannato da B
		if err != nil {
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result {
			http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
			return
		}
		result2, err2 := rt.db.IsBanned(User{UId: token}.ToDatabase(), User{UId: uIdint}.ToDatabase()) // A e' stato bannato da B
		if err2 != nil {
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result2 {
			http.Error(w, "L'utente e' stato bannato dall'utente attuale", http.StatusUnauthorized)
			return
		}

		// tutto ok, continuo
	}

	var comments []database.Comment
	comments, err := rt.db.GetComments(Image{IId: iIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comments)

}
