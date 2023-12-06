package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"
	"time"
	"io"

}


func (rt *_router) postImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")

	// non si possono caricare immagini per altre persone
	if !verificaToken(uId,r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	var image Image
	image.author = uId
	image.data = time.Now().UTC()

	err := json.NewDecoder(r.Body).Decode(&image.file,&image.descrizione)

	//controllo tipo di file?

	id,err := rt.db.PostImage(image.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	// /user/{idUser}/images/{imageId}:
	// creo il file vuoto per l'immagine
	out, _ := os.Create(filepath.Join("/user/",uId,"/images/",id))
	_ , _ = io.Copy(out, image.file)
	//id e' il nome del file

	out.Close()
	w.WriteHeader(http.StatusCreated)

}
