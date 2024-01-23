package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/Obrigad0/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

// getStream() preleva l'ultima immagine caricata per ogni utente seguito dal richiedente
func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))

	// solo il proprietario dell'account puo' vedere il proprio stream
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	var image []database.Image
	var user []database.User
	var err error
	// db-user.go
	image, user, err = rt.db.GetStream(User{UId: uIdint}.ToDatabase())
	if err != nil {
		// Errore nell'esecuzione della query
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	// Prelevo tutti i commenti di tutte le immagini
	for i := 0; i < len(image); i++ {
		image[i].Comments, err = rt.db.GetComments(Image{IId: image[i].IId}.ToDatabase())
	}

	combinedArray := map[string]interface{}{
		"utente":         user,
		"ultimaImmagine": image,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(combinedArray)

}
