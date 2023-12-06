package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}


func (rt *_router) deleteImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")
	iId := ps.ByName("imageId")

		// non si possono eliminare immagini di altre persone
	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//nessun problema, elimino l'immagine
	err := rt.db.DeleteImage(uId.ToDatabase(),iId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}

	//non finito
}