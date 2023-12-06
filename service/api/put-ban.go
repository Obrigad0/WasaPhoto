package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}
//putBan() inserisce il ban su un utente
func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")
	bId := ps.ByName("bannedUserId")

	if uId == bId{
		//Non ci si puo' autobannare !!!
		http.Error(w, "Non ti puoi bannare da sol* !!", http.StatusBadRequest)
		return 
	}

	// non si possono bannare user per altre persone
	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}


	err := rt.db.BanUser(uId.ToDatabase(),bId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore, errore del db o utente gia bannato", http.StatusInternalServerError)
	}

	//il ban implica che gli utenti non si vogliano piu seguire
	err := rt.db.UnfollowUser(bId.ToDatabase(),uId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}
	err := rt.db.UnfollowUser(uId.ToDatabase(),bId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}


	// Respond with 204 http status
	w.WriteHeader(http.StatusNoContent)
}
