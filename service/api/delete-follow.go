package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//deleteFollow() permette all'utente di levare il follow su un altro utente
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")
	fId := ps.ByName("followedUserId")
	//uId == token 

	// non si puo smettere di seguire user per altre persone
	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//Non gestisco il fatto che gli utenti si siano bannati tra loro, tanto si stanno defollowando

	//nessun problema, elimino il follow
	err := rt.db.UnfollowUser(uId.ToDatabase(),fId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}
