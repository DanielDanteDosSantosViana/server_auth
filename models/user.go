package models

import "gopkg.in/mgo.v2/bson"

type(

   User struct {
      Id bson.ObjectId
      Email string
      Senha string
      Token string
  }

)
