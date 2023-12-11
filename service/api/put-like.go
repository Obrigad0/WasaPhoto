package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// putLike() inserisce il like su un'immagine
func (rt *_router) putLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))     // proprietario immagine
	iIdint, _ := strconv.Atoi(ps.ByName("imageId"))    // id foto
	lIdint, _ := strconv.Atoi(ps.ByName("likeUserId")) // chi mette like, cioe' io

	if uIdint == lIdint {
		//Non puoi metterti like da sol*, e' triste!!!
		http.Error(w, "Non puoi metterti like da sol* !!", http.StatusBadRequest)
		return
	}

	if !verificaToken(lIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	//verifico se gli utenti si sono bannati tra di loro prima di aggiungere il like
	result, err := rt.db.IsBanned(User{UId: uIdint}.ToDatabase(), User{UId: lIdint}.ToDatabase()) // A e' stato bannato da B
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result {
		http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
		return
	}
	result2, err := rt.db.IsBanned(User{UId: lIdint}.ToDatabase(), User{UId: uIdint}.ToDatabase()) // A e' stato bannato da B
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result2 {
		http.Error(w, "L'utente e' stato bannato dall'utente attuale", http.StatusUnauthorized)
		return
	}

	//nessun problema, aggiungo il like
	err2 := rt.db.LikePhoto(User{UId: lIdint}.ToDatabase(), Image{IId: iIdint}.ToDatabase())
	if err2 != nil {
		http.Error(w, "Errore nella comunicazione con il db o like gia messo", http.StatusInternalServerError)

	}

	w.WriteHeader(http.StatusNoContent)

}
