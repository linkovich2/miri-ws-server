package routers

import (
	"github.com/jonathonharrell/miri-ws-server/engine/controllers"
	"github.com/jonathonharrell/miri-ws-server/engine/core/authentication"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetCharacterRoutes(router *mux.Router) *mux.Router {
	router.Handle("/characters/create",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.CreateCharacter),
		)).Methods("POST")

	router.Handle("/characters/list",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.ListCharacters),
		)).Methods("GET")

	return router
}
