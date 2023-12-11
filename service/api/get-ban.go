package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getBan() ritorna tutta la lista degli utenti bannati dall'utente
func (rt *_router) getBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))

	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	var user []User
	var err error
	user, err = rt.db.GetBanList(User{uId: uIdint}.ToDatabase())

	if err != nil {
		// Errore nell'esecuzione della query
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
