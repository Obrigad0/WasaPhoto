package api

import{

	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"

}
//deleteComments() elimina il proprio commento da una foto
func (rt *_router) deleteComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	uId := ps.ByName("idUser")		// proprietario foto
	iId := ps.ByName("imageId")		// id immagine di proprieta' di uId
	cId := ps.ByName("commentId")	// id commento
	// token : chi effettua l'operazione

	if !verificaLogin(r.Header.Get("Authorization")){
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return 
	}

	// verifico che chi sta provando ad elimnare il commento, lo abbia anche scritto
	// non puoi eliminare commenti di altre persone
	
	token := estrazioneToken(r.Header.Get("Authorization"))
	commentatore,err := rt.db.GetTheCommenter(cId.ToDatabase())
	if err != nil{
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}
	if token != commentatore{
		http.Error(w, "Non puoi eliminare commenti non scritti da te!", http.StatusUnauthorized)
		return
	}

	err := rt.db.UncommentPhoto(cId.ToDatabase())
	if(err != nil){
		http.Error(w, "Errore nella comunicazione con il db o commento non presente", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)

}