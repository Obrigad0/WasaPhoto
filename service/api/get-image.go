package api

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/Obrigad0/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	iIdint, _ := strconv.Atoi(ps.ByName("imageId")) // id foto

	if !verificaLogin(r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	// prelevo il token dell' user
	token := estrazioneToken(r.Header.Get("Authorization"))

	if uIdint != token {
		// se il richiedente del profilo NON e' il proprietario del profilo
		// verifico che l'Utente A (token) non sia stato bannato dall'Utente B (uId)
		// controllo se A e' stato bannato da B
		result, err := rt.db.IsBanned(User{UId: token}.ToDatabase(), User{UId: uIdint}.ToDatabase())
		if err != nil {
			http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
			return
		}
		if result {
			http.Error(w, "L'utente e' stato bannato, non e' autorizzato!", http.StatusUnauthorized)
			return
		}
		// tutto ok, continuo
	}

	// prelevo il profilo
	// dove inserisco l'utente
	var image database.Image
	// db-image.go
	image, err := rt.db.GetImage(User{UId: token}.ToDatabase(), Image{IId: iIdint}.ToDatabase())
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	// Prelevo tutti i commenti di tutte le immagini
	image.Comments, err = rt.db.GetComments(Image{IId: image.IId}.ToDatabase())

	w.Header().Set("Content-Type", "application/json")

	// invio le informazioni dell'immagine
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(image)

	// invio l'immagine

	path := filepath.Join(cartellaPrincipale, strconv.Itoa(uIdint), "/imgUs/", strconv.Itoa(iIdint))
	http.ServeFile(w, r, path)
}
