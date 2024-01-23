package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/Obrigad0/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getAllImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))

	if !verificaLogin(r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	var image []database.Image

	image, err := rt.db.GetAllImage(User{UId: uIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	// Prelevo tutti i commenti di tutte le immagini, poco efficiente ma ci piace cosi
	for i := 0; i < len(image); i++ {
		image[i].Comments, err = rt.db.GetComments(Image{IId: image[i].IId}.ToDatabase())
	}

	w.Header().Set("Content-Type", "application/json")

	// invio le informazioni dell'immagine
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(image)
}
