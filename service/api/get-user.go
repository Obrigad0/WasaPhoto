package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//getUser() ritorna tutte le informazioni su un utente
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")

	if !verificaLogin(r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//prelevo il token dell' user
	token := estrazioneToken(r.Header.Get("Authorization"))

	if(uId != token){
		//se il richiedente del profilo NON e' il proprietario del profilo
		//verifico che l'Utente A (token) non sia stato bannato dall'Utente B (uId)
		//controllo se A e' stato bannato da B
		result,err := rt.db.IsBanned(token.ToDatabase(),uId.ToDatabase())
		if err != nil{
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result{
			http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
			return 
		}
		//tutto ok, continuo
	}

	//prelevo il profilo

	var user User
	var follower []integer
	var following []integer
	var ban []integer

	user,err := rt.db.GetUser(uId.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	follower,err := rt.db.GetFollowerList(uId.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	
	following,err := rt.db.GetFollowingList(uId.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	ban,err := rt.db.GetBanListVINT(uId.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	user.followers = follower
	user.following = following
	user.ban = ban 

	w.WriteHeader(http.StatusOK)

}
