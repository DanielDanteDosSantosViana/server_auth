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
)

type(
    UserController struct{
            session *mgo.Session
    }
)

func NewUserController() *UserController{
    return &UserController{db.NewSession()}
}

func (userC * UserController) Create(respWriter http.ResponseWriter, req *http.Request) {

    decoder := json.NewDecoder(req.Body)

    user:= models.User{}
    err := decoder.Decode(&user)
    if err != nil {
        log.Fatal(err)
    }

    user.Id = bson.NewObjectId()

    userC.session.DB("server_auth").C("users").Insert(user)
    userJ,_:= json.Marshal(user)
    respWriter.Header().Set("Content-Type", "application/json")
    respWriter.WriteHeader(200)
    fmt.Fprintf(respWriter, "%s", userJ)
 }


