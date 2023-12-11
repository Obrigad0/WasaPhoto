package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// deleteBan() elimina il ban aggiunto dall'utente su un altro utente
func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	bIdint, _ := strconv.Atoi(ps.ByName("bannedUserId"))
	//uId == Token

	// non si possono sbannare user per altre persone
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	//nessun problema, elimino il ban
	err := rt.db.UnbanUser(User{uId: uIdint}.ToDatabase(), User{uId: bIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}
