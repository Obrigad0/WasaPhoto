package api

import (
	"encoding/json"
	"net/http"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/Obrigad0/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

// getSearchUser() ritorna tuna lista di utenti
func (rt *_router) getSearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	nome := r.URL.Query().Get("queryNomeUtente")
	if !verificaLogin(r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	token := estrazioneToken(r.Header.Get("Authorization"))

	var users []database.User
	users, err := rt.db.GetSearchUser(User{UId: token, Name: nome}.ToDatabase())

	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// se non ci sono utenti da tornare

	if len(users) == 0 {
		_ = json.NewEncoder(w).Encode([]database.User{})
		return

	} else {
		for i := 0; i < len(users); i++ {

			var follower []int
			var following []int

			follower, err2 := rt.db.GetFollowerList(User{UId: users[i].UId}.ToDatabase())
			if err2 != nil {
				http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
				return
			}

			following, err3 := rt.db.GetFollowingList(User{UId: users[i].UId}.ToDatabase())
			if err3 != nil {
				http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
				return
			}
			users[i].Followers = follower
			users[i].Following = following

		}

	}
	_ = json.NewEncoder(w).Encode(users)

}
