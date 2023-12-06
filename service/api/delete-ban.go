package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//deleteBan() elimina il ban aggiunto dall'utente su un altro utente
func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")		
	bId := ps.ByName("bannedUserId")
	//uId == Token 


	// non si possono sbannare user per altre persone
	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//nessun problema, elimino il ban
	err := rt.db.UnbanUser(uId.ToDatabase(),bId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}
