package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//dobbiamo tornare il token, l'id che rappresenta l'utente
	w.Header().Set("Content-Type", "application/json")

	//nome inserito dall'utente
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)

	if err != nil {
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return
	}

	//controllo se l'utente esiste
	var uId int
	uId, err2 := rt.db.Access(User{name: username}.ToDatabase())

	if err2 != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	if uId < 0 {
		//Errore, l'utente non esiste va creato!
		uId, err := rt.db.CreateUser(User{name: username}.ToDatabase())
		// la funzione mi torna anche l'id dell'utente creato
		if err != nil || uId < 0 {
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		// uso l'id per creare la dir dell'utente nel server
		// /user/id
		// /user/id/image
		//	FORSE CAMBIARE
		path := "/user/" + strconv.Itoa(uId) + "/image"
		os.MkdirAll(path, os.ModePerm)

	}
	//tutto fatto!
	// invio l'id all'utente
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(uId)
	if err != nil {
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return
	}

}
