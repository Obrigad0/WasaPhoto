package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//putFollow() 
func (rt *_router) putFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")
	fId := ps.ByName("followedUserId")

	if uId == fId{
		//Non puoi seguire te stess* !!!
		http.Error(w, "Non ti puoi seguire da sol* !!", http.StatusBadRequest)
		return 
	}

	// non si possono seguire user per altre persone
	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//verifico se gli utenti si sono bannati tra di loro prima di aggiungere il follow
	//Non posso seguire una persona che mi ha bannato
	//Non posso seguire una persona che ho bannato
	result,err := rt.db.IsBanned(uId.ToDatabase(),fId.ToDatabase())	// A e' stato bannato da B
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result{
		http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
		return 
	}
	result,err := rt.db.IsBanned(fId.ToDatabase(),uId.ToDatabase())// A e' stato bannato da B
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result{
		http.Error(w, "L'utente e' stato bannato dall'utente attuale, quindi non puo seguirlo!", http.StatusUnauthorized)
		return 
	}

	//nessun problema, aggiungo il follow
	err := rt.db.FollowUser(uId.ToDatabase(),fId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore nella comunicazione con il db o utente gia seguito", http.StatusInternalServerError)

	}

	w.WriteHeader(http.StatusNoContent)

}
