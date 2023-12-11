package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// putFollow()
func (rt *_router) putFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	fIdint, _ := strconv.Atoi(ps.ByName("followedUserId"))

	if uIdint == fIdint {
		//Non puoi seguire te stess* !!!
		http.Error(w, "Non ti puoi seguire da sol* !!", http.StatusBadRequest)
		return
	}

	// non si possono seguire user per altre persone
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	//verifico se gli utenti si sono bannati tra di loro prima di aggiungere il follow
	//Non posso seguire una persona che mi ha bannato
	//Non posso seguire una persona che ho bannato
	result, err := rt.db.IsBanned(User{uId: uIdint}.ToDatabase(), User{uId: fIdint}.ToDatabase()) // A e' stato bannato da B
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result {
		http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
		return
	}
	result2, err2 := rt.db.IsBanned(User{uId: fIdint}.ToDatabase(), User{uId: uIdint}.ToDatabase()) // A e' stato bannato da B
	if err2 != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result2 {
		http.Error(w, "L'utente e' stato bannato dall'utente attuale, quindi non puo seguirlo!", http.StatusUnauthorized)
		return
	}

	//nessun problema, aggiungo il follow
	err3 := rt.db.FollowUser(User{uId: uIdint}.ToDatabase(), User{uId: fIdint}.ToDatabase())
	if err3 != nil {
		http.Error(w, "Errore nella comunicazione con il db o utente gia seguito", http.StatusInternalServerError)

	}

	w.WriteHeader(http.StatusNoContent)

}
