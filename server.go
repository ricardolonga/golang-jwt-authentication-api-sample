package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/ricardolonga/golang-jwt-authentication-api-sample/settings"
	"github.com/ricardolonga/golang-jwt-authentication-api-sample/routers"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
