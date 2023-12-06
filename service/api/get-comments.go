package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}

//getComments() ritorna tutta la lista dei commenti di un'immagine
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")
	iId := ps.ByName("imageId")


	if !verificaLogin(r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	// tutti possono vedere i commenti di una foto

	//prelevo il token dell' user
	token := estrazioneToken(r.Header.Get("Authorization"))

	if(uId != token){
		//se la persona che richiede i commenti NON e' il proprietario del profilo
		//verifico che l'Utente A (token) non sia stato bannato dall'Utente B (uId) proprietario foto
		//e viceversa (non puoi vedere i commenti  di una persona che hai bannato)
		//controllo se A e' stato bannato da B
		result,err := rt.db.IsBanned(uId.ToDatabase(),token.ToDatabase())	// A e' stato bannato da B
		if err != nil{
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result{
			http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
			return 
		}
		result,err := rt.db.IsBanned(token.ToDatabase(),uId.ToDatabase())// A e' stato bannato da B
		if err != nil{
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result{
			http.Error(w, "L'utente e' stato bannato dall'utente attuale", http.StatusUnauthorized)
			return 
		}		
		
		//tutto ok, continuo
	}

	var comments []Comment
	comments,err := rt.db.GetComments(iId.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}



