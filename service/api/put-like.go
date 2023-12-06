package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//putLike() inserisce il like su un'immagine
func (rt *_router) putLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")		// proprietario immagine
	iId := ps.ByName("imageId")		// id foto 
	lId := ps.ByName("likeUserId")	// chi mette like, cioe' io

	if uId == lId{
		//Non puoi metterti like da sol*, e' triste!!!
		http.Error(w, "Non puoi metterti like da sol* !!", http.StatusBadRequest)
		return 
	}

	if !verificaToken(lId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//verifico se gli utenti si sono bannati tra di loro prima di aggiungere il like
	result,err := rt.db.IsBanned(uId.ToDatabase(),lId.ToDatabase())	// A e' stato bannato da B
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result{
		http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
		return 
	}
	result,err := rt.db.IsBanned(lId.ToDatabase(),uId.ToDatabase())// A e' stato bannato da B
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if result{
		http.Error(w, "L'utente e' stato bannato dall'utente attuale", http.StatusUnauthorized)
		return 
	}

	//nessun problema, aggiungo il like
	err := rt.db.LikePhoto(lId.ToDatabase(),iId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore nella comunicazione con il db o like gia messo", http.StatusInternalServerError)

	}

	w.WriteHeader(http.StatusNoContent)

}