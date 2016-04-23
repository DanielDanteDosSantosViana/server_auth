package controllers

import(
    "log"
    "fmt"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "server_auth/models"
    "server_auth/db"
    "server_auth/services"
    "server_auth/util"

    "time"
)

type(
    AuthenticatorController struct{
            session *mgo.Session
    }
)

func NewAuthenticatorController() *AuthenticatorController{
    return &AuthenticatorController{db.NewSession()}
}

func (auth * AuthenticatorController) GenerateAuthenticate(respWriter http.ResponseWriter, req *http.Request) {

    decoder := json.NewDecoder(req.Body)

    user:= models.User{}
    err := decoder.Decode(&user)
    if err != nil {
        log.Fatal(err)
    }

    if err := auth.session.DB("server_auth").C("users").Find(bson.M{"email": user.Email , "senha":user.Senha}).One(&user); err != nil {
        respWriter.WriteHeader(404)
        return
    }

    options := util.Options{
        SigningMethod: "HS256",
        PrivateKey:    "darthvader",
        PublicKey:     "teste",
        Expiration:    60 * time.Minute,
    }

    tokenString, err := services.GenerateJWTToken(string(user.Id),options)
    if err!=nil{
        panic(err)
    }

    user.Token = tokenString
    userJ,_:= json.Marshal(user)

    respWriter.Header().Set("Content-Type", "application/json")
    respWriter.WriteHeader(200)
    fmt.Fprintf(respWriter, "%s", userJ)
 }


