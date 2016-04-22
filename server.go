package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"

    "server_auth/routes"

)

func main() {

    router:= mux.NewRouter().StrictSlash(true)
    routes.InitRoutes(router)
    log.Fatal(http.ListenAndServe(":8080", router))
}



