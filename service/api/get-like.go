package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getLike() preleva tutti i like di un'immagine
func (rt *_router) getLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	iIdint, _ := strconv.Atoi(ps.ByName("imageId")) // id foto

	//solo il proprietario dell'account puo' vedere i like di una sua foto
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	var user []User
	var err error
	user, err = rt.db.GetLikesList(Image{iId: iIdint}.ToDatabase())

	if err != nil {
		// Errore nell'esecuzione della query
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
