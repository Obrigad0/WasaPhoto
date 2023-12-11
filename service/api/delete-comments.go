package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// deleteComments() elimina il proprio commento da una foto
func (rt *_router) deleteComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	cIdint, _ := strconv.Atoi(ps.ByName("commentId")) // id commento
	// token : chi effettua l'operazione

	if !verificaLogin(r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	// verifico che chi sta provando ad elimnare il commento, lo abbia anche scritto
	// non puoi eliminare commenti di altre persone

	token := estrazioneToken(r.Header.Get("Authorization"))
	commentatore, err := rt.db.GetTheCommenter(Comment{cId: cIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if token != commentatore {
		http.Error(w, "Non puoi eliminare commenti non scritti da te!", http.StatusUnauthorized)
		return
	}

	err2 := rt.db.UncommentPhoto(Comment{cId: cIdint}.ToDatabase())
	if err2 != nil {
		http.Error(w, "Errore nella comunicazione con il db o commento non presente", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)

}
