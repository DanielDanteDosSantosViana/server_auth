
package services

import(

  "net/http"
  "fmt"
  "io/ioutil"
  "bytes"
)

type(

  SenderHTTP struct{}
)

func NewSenderHTTP()*SenderHTTP {
    return &SenderHTTP{}
}


func (send *SenderHTTP) Send(url string, method string, data []byte){
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
  req.Header.Set("X-Custom-Header", "myvalue")
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
        panic(err)
  }
  defer resp.Body.Close()

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))

}

