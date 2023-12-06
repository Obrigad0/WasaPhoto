package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//getStream() preleva l'ultima immagine caricata per ogni utente seguito dal richiedente
func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")

	//solo il proprietario dell'account puo' vedere il proprio stream
	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	var image []Image
	var user  []User
	// db-user.go
	image,user, err = rt.db.GetStream(uId.ToDatabase())
	if err != nil {
		// Errore nell'esecuzione della query
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	
}
