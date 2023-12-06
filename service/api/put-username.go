package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//putUsername() cambia nome utente dell'utente, se autorizzato
func (rt *_router) putUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")

	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	// variabile utilizzta per memorizzare l'username inviato dal clienet
	var newUserName string
	err := json.NewDecoder(r.Body).Decode(&newUserName)
	if err != nil{
		//errore nella richiesta json
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return 
	}
	err = rt.db.ChangeUserName(uId.ToDatabase(),newUserName.ToDatabase())
	if err != nil {
		// Errore nell'esecuzione della query
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	//tutto ok, l'operazione di modifica e' stata effettuata con successo 
	w.WriteHeader(http.StatusNoContent)
}