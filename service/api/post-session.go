package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"
	"strconv"
	"os"


}

func (rt *_router) postSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	//dobbiamo tornare il token, l'id che rappresenta l'utente
	w.Header().Set("Content-Type", "application/json")

	//nome inserito dall'utente
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)

	if err != nil{
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return 
	}

	//controllo se l'utente esiste
	var uId integer
	uId,err := rt.db.Access(username.ToDatabase())

	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	if uId < 0{
		//Errore, l'utente non esiste va creato!
		uId,err := rt.db.CreateUser(username.ToDatabase())
		// la funzione mi torna anche l'id dell'utente creato
		if err != nil || uId < 0{
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		// uso l'id per creare la dir dell'utente nel server
		// /user/id
		// /user/id/image
		path := "/user/" + strconv.Itoa(uId) + "/image"
		err := os.MkdirAll(path, os.ModePerm)


	}
	//tutto fatto!
	// invio l'id all'utente
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(uId)
	if err != nil {
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return 
	}

}
