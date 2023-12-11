package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//"encoding/json"

func (rt *_router) deleteImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	iIdint, _ := strconv.Atoi(ps.ByName("imageId")) // id foto

	// non si possono eliminare immagini di altre persone
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	//nessun problema, elimino l'immagine
	err := rt.db.DeleteImage(User{uId: uIdint}.ToDatabase(), Image{iId: iIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore, errore del db", http.StatusInternalServerError)
	}

	//non finito
}
