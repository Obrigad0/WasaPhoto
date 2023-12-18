package api

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))

	// non si possono caricare immagini per altre persone
	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	var image Image
	image.Author = strconv.Itoa(uIdint)
	image.Data = time.Now().UTC()

	// ParseMultipartForm analizza il corpo di una richiesta multipart/form-data
	// 30 << 20 e' la dimensione max in byte della rchiesta multipart/form-data
	err := r.ParseMultipartForm(30 << 20) // << e' lo shift a sinistra per la moltiplicazione
	if err != nil {
		http.Error(w, "Impossibile analizzare la richiesta multipart", http.StatusBadRequest)
		return
	}

	// prendo il file dalla richiesta
	// "file" e' il file vero e proprio in formato binario
	// "metadata" sono tutti le informazioni associate all'immagine, come il nome, che non ci interessa visto che e' sostituito dall'id dell'immagine creato dal dbms
	file, _, _ := r.FormFile("file")
	defer file.Close() //chiudo il file appena preso

	descrizione := r.FormValue("descrizione")

	// prendo la descrizione
	image.Descrizione = descrizione

	id, err := rt.db.PostImage(image.ToDatabase())
	if err != nil {
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	// /user/{idUser}/images/{imageId}:
	// creo il file vuoto per l'immagine
	out, _ := os.Create(filepath.Join("/user/", strconv.Itoa(uIdint), "/images/", strconv.Itoa(id)))
	_, _ = io.Copy(out, file)
	// id e' il nome del file

	out.Close()
	w.WriteHeader(http.StatusAccepted)

}
