package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetAuthenticationRoutes(router)
	router = SetCharacterRoutes(router)
	router = SetPlayRoutes(router)
	return router
}
