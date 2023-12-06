package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//deleteLike() elimina il like inserito da un'immagine
func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")	// proprietario immagine
	iId := ps.ByName("imageId")	// id foto 
	lId := ps.ByName("likeUserId")	// il like da levare
	//lId == token

	//non puoi levare il like di un altra persona
	if !verificaToken(lId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//nessun problema, levo il like
	err := rt.db.UnlikePhoto(lId.ToDatabase(),iId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore nella comunicazione con il db o like non presente", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
