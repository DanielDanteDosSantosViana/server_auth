package util

import "time"

type(

  Options struct{

        SigningMethod string
        PrivateKey    string
        PublicKey     string
        Expiration    time.Duration
  }

)

func NewOptions() *Options{
    return &Options{"HS256","darthvader","teste",60 * time.Minute}
}
