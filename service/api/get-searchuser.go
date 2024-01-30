package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/Obrigad0/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

// getSearchUser() ritorna tuna lista di utenti
func (rt *_router) getSearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	nome := r.URL.Query().Get("queryNomeUtente")
	fmt.Println("Ecco il nome arrivato: ", nome)
	if !verificaLogin(r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	token := estrazioneToken(r.Header.Get("Authorization"))
	fmt.Println("Ecco il token ", token)

	var users []database.User
	users, err := rt.db.GetSearchUser(User{UId: token, Name: nome}.ToDatabase())

	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// se non ci sono utenti da tornare
	if len(users) == 0 {
		_ = json.NewEncoder(w).Encode([]database.User{})
		return
	}
	_ = json.NewEncoder(w).Encode(users)
}
