package routers

import (
	"github.com/jonathonharrell/miri-ws-server/engine/controllers"
	"github.com/jonathonharrell/miri-ws-server/engine/core/authentication"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetPlayRoutes(router *mux.Router) *mux.Router {
	router.Handle("/play",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Play),
		)).Methods("GET")

	return router
}
