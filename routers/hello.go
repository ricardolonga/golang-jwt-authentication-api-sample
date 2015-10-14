package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/ricardolonga/golang-jwt-authentication-api-sample/core/authentication"
	"github.com/ricardolonga/golang-jwt-authentication-api-sample/controllers"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
