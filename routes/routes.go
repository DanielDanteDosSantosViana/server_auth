
package routes

import(
  "github.com/gorilla/mux"
  "server_auth/controllers"

)

func InitRoutes( router * mux.Router ){
    auth:= controllers.NewAuthenticatorController()
    router.HandleFunc("/mwebauth/authentication",auth.Authenticate).Methods("POST")
}




