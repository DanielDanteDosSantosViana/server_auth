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
