package controllers

import(
    "log"
    "fmt"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/dgrijalva/jwt-go"

    "server_auth/models"
    "server_auth/db"
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

func (auth * AuthenticatorController) Authenticate(respWriter http.ResponseWriter, req *http.Request) {

    decoder := json.NewDecoder(req.Body)

    user:= models.User{}
    err := decoder.Decode(&user)
    if err != nil {
        log.Fatal(err)
    }

    user.Id = bson.NewObjectId()

    if err := auth.session.DB("server_auth").C("users").Find(bson.M{"email": user.Email , "senha":user.Senha}).One(&user); err != nil {
        respWriter.WriteHeader(404)
        return
    }

    tokenString, err := generateToken()
    if err!=nil{
        panic(err)
    }

    user.Token = tokenString
    userJ,_:= json.Marshal(user)

    respWriter.Header().Set("Content-Type", "application/json")
    respWriter.WriteHeader(200)
    fmt.Fprintf(respWriter, "%s", userJ)
 }


func generateToken() (string,error){
    token := jwt.New(jwt.SigningMethodRS512)
    token.Claims["exp"] = time.Now()
    token.Claims["iat"] = time.Now()
    token.Claims["sub"] = "danielteste"
    tokenString, err := token.SignedString("vasco")
    if err != nil {
        panic(err)
        return "", err
    }

    return tokenString,nil

}

func (auth * AuthenticatorController) Find(respWriter http.ResponseWriter, req *http.Request) {

    decoder := json.NewDecoder(req.Body)

    user:= models.User{}
    err := decoder.Decode(&user)
    if err != nil {
        log.Fatal(err)
    }

    user.Id = bson.NewObjectId()

    if err := auth.session.DB("server_auth").C("users").Find(bson.M{"email": user.Email}).One(&user); err != nil {
        respWriter.WriteHeader(404)
        return
    }

    userJ,_:= json.Marshal(user)
    fmt.Println(user.Email)
    fmt.Println(user.Senha)
    fmt.Println(user.Token)
    fmt.Println(string(userJ))

    respWriter.Header().Set("Content-Type", "application/json")
    respWriter.WriteHeader(200)
    fmt.Fprintf(respWriter, "%s", userJ)
 }

