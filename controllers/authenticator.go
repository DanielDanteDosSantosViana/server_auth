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

    options := util.NewOptions()

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

func (auth * AuthenticatorController) ValidateAuthentication(respWriter http.ResponseWriter, req *http.Request) {

    decoder := json.NewDecoder(req.Body)
    token:= models.Token{}
    err := decoder.Decode(&token)
    if err != nil {
        log.Fatal(err)
        respWriter.Header().Set("Content-Type", "application/json")
        respWriter.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintln(respWriter, "Error while Parsing Token!")
     }

    err = services.ValidateJWTToken(token.Value)
    if err!=nil{
        respWriter.Header().Set("Content-Type", "application/json")
        respWriter.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintln(respWriter, "WHAT? Invalid Token!")
    }else{

        respWriter.Header().Set("Content-Type", "application/json")
        respWriter.WriteHeader(200)
        fmt.Fprintf(respWriter, "%s", "OK")
    }
 }
