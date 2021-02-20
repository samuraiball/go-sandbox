package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func main() {

	claims := MyClaims{
		"userid",
		jwt.StandardClaims{
			Issuer:    "issuer",
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}

	fmt.Println(time.Now().Add(time.Hour * 3).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var key = []byte(``)

	ss, err := token.SignedString(key)
	if err != nil {
		println(err.Error())
	}

	fmt.Printf("%v\n", ss)

}
