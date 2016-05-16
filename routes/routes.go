
package routes

import(
  "github.com/gorilla/mux"
  "server_auth/controllers"

)

func InitRoutes( router * mux.Router ){
    auth:= controllers.NewAuthenticatorController()
    user:= controllers.NewUserController()
    receiver:= controllers.NewPacketReceiverHTTPController()

    //generate tooken
    router.HandleFunc("/mwebauth/authentication",auth.GenerateAuthenticate).Methods("POST")

    //create users
    router.HandleFunc("/mwebauth/user",user.Create).Methods("POST")

    //validate
    router.HandleFunc("/mwebauth/validateToken",auth.ValidateAuthentication).Methods("POST")

    //receiver HTTP
    router.HandleFunc("/mwebauth/receiver",receiver.ReceivedPacket).Methods("POST")


}




