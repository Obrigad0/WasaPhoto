package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// dobbiamo tornare il token, l'id che rappresenta l'utente
	w.Header().Set("Content-Type", "application/json")

	// nome inserito dall'utente
	var username User
	err := json.NewDecoder(r.Body).Decode(&username)

	if err != nil {
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return
	}

	// controllo se l'utente esiste
	var uId int
	var contenitore int
	uId, err2 := rt.db.Access(User{Name: username.Name}.ToDatabase())

	if err2 != nil {
		a := 1
		b := a
		a = b
	}

	if uId < 0 {
		// Errore, l'utente non esiste va creato!
		uId, err := rt.db.CreateUser(User{Name: username.Name}.ToDatabase())
		// la funzione mi torna anche l'id dell'utente creato
		if err != nil || uId < 0 {
			http.Error(w, "Errore nella comunicazione con il db 2", http.StatusInternalServerError)
			return
		}
		// uso l'id per creare la dir dell'utente nel server
		// /wp/id
		// /wp/id/image
		path := filepath.Join(cartellaPrincipale, strconv.Itoa(uId))
		err2 := os.MkdirAll(filepath.Join(path, "imgUs"), os.ModePerm)
		if err2 != nil {
			http.Error(w, "Errore nella creazione della dir dell'utente", http.StatusInternalServerError)
			return
		}
		contenitore = uId

	}
	// tutto fatto!
	// invio l'id all'utente
	if uId == -1 {
		uId = contenitore
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(uId)
	if err != nil {
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return
	}

}
