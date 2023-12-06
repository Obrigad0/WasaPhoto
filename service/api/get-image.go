package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}


func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")
	iId := ps.ByName("imageId")

	if !verificaLogin(r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	//prelevo il token dell' user
	token := estrazioneToken(r.Header.Get("Authorization"))

	if(uId != token){
		//se il richiedente del profilo NON e' il proprietario del profilo
		//verifico che l'Utente A (token) non sia stato bannato dall'Utente B (uId)
		//controllo se A e' stato bannato da B
		result,err := rt.db.IsBanned(token.ToDatabase(),uId.ToDatabase())
		if err != nil{
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result{
			http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
			return 
		}
		//tutto ok, continuo
	}

	//prelevo il profilo
	// dove inserisco l'utente
	var image Image
	// db-image.go
	image,err := rt.db.GetImage(token.ToDatabase(),iId.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}


	//non finito 
	
	w.WriteHeader(http.StatusOK)

}