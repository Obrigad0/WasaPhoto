package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getUser() ritorna tutte le informazioni su un utente
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))

	if !verificaLogin(r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	//prelevo il token dell' user
	token := estrazioneToken(r.Header.Get("Authorization"))

	if uIdint != token {
		//se il richiedente del profilo NON e' il proprietario del profilo
		//verifico che l'Utente A (token) non sia stato bannato dall'Utente B (uId)
		//controllo se A e' stato bannato da B
		result, err := rt.db.IsBanned(User{uId: token}.ToDatabase(), User{uId: uIdint}.ToDatabase())
		if err != nil {
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result {
			http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
			return
		}
		//tutto ok, continuo
	}

	//prelevo il profilo

	var user User
	var follower []int
	var following []int
	var ban []int

	user, err := rt.db.GetUser(User{uId: uIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	follower, err2 := rt.db.GetFollowerList(User{uId: uIdint}.ToDatabase())
	if err2 != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	following, err3 := rt.db.GetFollowingList(User{uId: uIdint}.ToDatabase())
	if err3 != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	ban, err4 := rt.db.GetBanListVINT(User{uId: uIdint}.ToDatabase())
	if err4 != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	user.followers = follower
	user.following = following
	user.ban = ban

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
