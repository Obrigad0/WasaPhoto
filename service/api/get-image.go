package api

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Obrigad0/WasaPhoto/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uIdint, _ := strconv.Atoi(ps.ByName("idUser"))
	iIdint, _ := strconv.Atoi(ps.ByName("imageId")) // id foto

	// w.Header().Set("Content-Type", "application/json")

	// w.WriteHeader(http.StatusOK)

	path := filepath.Join(cartellaPrincipale, strconv.Itoa(uIdint), "/imgUs/", strconv.Itoa(iIdint))
	http.ServeFile(w, r, path)

}
