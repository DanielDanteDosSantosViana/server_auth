
package routes

import(
  "github.com/gorilla/mux"
  "server_auth/controllers"

)

func InitRoutes( router * mux.Router ){
    auth:= controllers.NewAuthenticatorController()
    user:= controllers.NewUserController()


    //generate tooken
    router.HandleFunc("/mwebauth/authentication",auth.GenerateAuthenticate).Methods("POST")

    //create users
    router.HandleFunc("/mwebauth/users/{email}",user.Create).Methods("POST")

}




