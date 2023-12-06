package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}


//getBan() ritorna tutta la lista degli utenti bannati dall'utente
func (rt *_router) getBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")

	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	var user  []User
	user, err = rt.db.GetBanList(uId.ToDatabase())

	if err != nil {
		// Errore nell'esecuzione della query
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
