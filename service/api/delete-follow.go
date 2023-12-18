package api

import (
	"net/http"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// deleteFollow() permette all'utente di levare il follow su un altro utente
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	fIdint, _ := strconv.Atoi(ps.ByName("followedUserId"))
	// uId == token

	// non si puo smettere di seguire user per altre persone
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	// Non gestisco il fatto che gli utenti si siano bannati tra loro, tanto si stanno defollowando

	// nessun problema, elimino il follow
	err := rt.db.UnfollowUser(User{UId: uIdint}.ToDatabase(), User{UId: fIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}
