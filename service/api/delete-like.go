package api

import (
	"net/http"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// deleteLike() elimina il like inserito da un'immagine
func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	iIdint, _ := strconv.Atoi(ps.ByName("imageId"))    // id foto
	lIdint, _ := strconv.Atoi(ps.ByName("likeUserId")) // il like da levare
	// lId == token

	// non puoi levare il like di un altra persona
	if !verificaToken(lIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	// nessun problema, levo il like
	err := rt.db.UnlikePhoto(User{UId: lIdint}.ToDatabase(), Image{IId: iIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db o like non presente", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
