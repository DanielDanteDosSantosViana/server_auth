
package services

import(

  "time"
  "log"

  "server_auth/util"
  "server_auth/crypto"
  jwt "github.com/dgrijalva/jwt-go"
  //"github.com/gorilla/context"


)


func GenerateJWTToken(userId string, op util.Options) (string, error) {
    t := jwt.New(jwt.GetSigningMethod(op.SigningMethod))

    now := time.Now()
    t.Claims["iat"] = now.Unix()
    t.Claims["exp"] = now.Add(op.Expiration).Unix()
    t.Claims["sub"] = userId
    t.Claims["jti"] = crypto.GenerateRandomKey(32)

    tokenString, err := t.SignedString([]byte(op.PrivateKey))
    if err != nil {
        logError("ERROR: GenerateJWTToken: %v\n", err)
    }
    return tokenString, err

}

/*
func ValidateJWTToken() (string, error) {


}
*/

func logError(format string, err interface{}) {
  if err != nil {
    log.Printf(format, err)
  }
}


