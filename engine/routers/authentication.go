package routers

import (
	"github.com/jonathonharrell/miri-ws-server/engine/controllers"
	"github.com/jonathonharrell/miri-ws-server/engine/core/authentication"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")

	router.Handle("/refresh-token", // @todo do we need this route? probably not
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")

	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	// @ todo forgot password, password reset, update account actions

	return router
}
