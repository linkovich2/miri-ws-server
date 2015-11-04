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
	router.HandleFunc("/forgot_password", controllers.ForgotPassword).Methods("POST")
	router.HandleFunc("/reset_password", controllers.ResetPassword).Methods("POST")

	router.Handle("/refresh-token",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")

	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	// @ todo update account actions

	return router
}
