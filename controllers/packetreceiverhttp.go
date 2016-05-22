package controllers

import(
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "server_auth/models"
    "server_auth/services"

)

type(
    PacketReceiverHTTPController struct{
    }
)

func NewPacketReceiverHTTPController() *PacketReceiverHTTPController{
    return &PacketReceiverHTTPController{}
}

func (receiver * PacketReceiverHTTPController) ReceivedPacket(respWriter http.ResponseWriter, req *http.Request) {

    body, _ := ioutil.ReadAll(req.Body)
    bytes := []byte(body)
    var requestData []models.RequestData
    json.Unmarshal(bytes, &requestData)

    token:=requestData[len(requestData)-1].Value
    if token!=""{
        err := services.ValidateJWTToken(token)
        if err!=nil{
            respWriter.Header().Set("Content-Type", "application/json")
            respWriter.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintln(respWriter, "WHAT? Invalid Token!")
        }else{

            sender:=services.NewSenderHTTP()

            go sender.Send("http://localhost:5000/receiver_server/receiver","POST",bytes)

            respWriter.Header().Set("Content-Type", "application/json")
            respWriter.WriteHeader(200)
            fmt.Fprintf(respWriter, "%s", "OK")

        }
    }else{
        respWriter.Header().Set("Content-Type", "application/json")
        respWriter.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintf(respWriter, "%s", "Has not sent token")
    }

 }
