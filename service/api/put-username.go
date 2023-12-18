package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

// putUsername() cambia nome utente dell'utente, se autorizzato
func (rt *_router) putUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))

	if !verificaToken(uIdint, r.Header.Get("Authorization")) {
		// Token non valido, ritorno Errore 401 al client
		http.Error(w, "Token non valido", http.StatusUnauthorized)
		return
	}

	// variabile utilizzta per memorizzare l'username inviato dal clienet
	var newUserName string
	err := json.NewDecoder(r.Body).Decode(&newUserName)
	if err != nil {
		//errore nella richiesta json
		http.Error(w, "Errore nella richiesta json", http.StatusBadRequest)
		return
	}
	err = rt.db.ChangeUserName(User{UId: uIdint}.ToDatabase(), User{Name: newUserName}.ToDatabase())
	if err != nil {
		// Errore nell'esecuzione della query
		http.Error(w, "Errore nella comunicazione con il db", http.StatusInternalServerError)
		return
	}

	// tutto ok, l'operazione di modifica e' stata effettuata con successo
	w.WriteHeader(http.StatusNoContent)
}
