package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// putBan() inserisce il ban su un utente
func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	bIdint, _ := strconv.Atoi(ps.ByName("bannedUserId"))

	if uIdint == bIdint {
		//Non ci si puo' autobannare !!!
		http.Error(w, "Non ti puoi bannare da sol* !!", http.StatusBadRequest)
		return
	}

	// non si possono bannare user per altre persone
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	err := rt.db.BanUser(User{uId: uIdint}.ToDatabase(), User{uId: bIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore, errore del db o utente gia bannato", http.StatusInternalServerError)
	}

	//il ban implica che gli utenti non si vogliano piu seguire
	err2 := rt.db.UnfollowUser(User{uId: bIdint}.ToDatabase(), User{uId: uIdint}.ToDatabase())
	if err2 != nil {
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}
	err3 := rt.db.UnfollowUser(User{uId: uIdint}.ToDatabase(), User{uId: bIdint}.ToDatabase())
	if err3 != nil {
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}

	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
