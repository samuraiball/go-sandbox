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
		"b5548e70-732d-11eb-971c-57ee55c52577",
		jwt.StandardClaims{
			Issuer:    "issuer",
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	var privateKeyRaw = []byte(`
-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----
`)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyRaw)
	if err != nil {
		panic(err)
	}

	ss, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n")
	fmt.Printf("%v\n", ss)

	time.Sleep(time.Second * 3)

	fmt.Printf("\n\n")

	var keyPublicRaw = []byte(`
-----BEGIN PUBLIC KEY-----
-----END PUBLIC KEY-----
`)

	var tokenString = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9." +
		"dXNlcklkOmI1NTQ4ZTcwLTczMmQtMTFlYi05NzFjLTU3ZWU1NWM1MjU3OSBleHA6MTYxMzgwMzgyNyBpc3M6aXNzdWVyCg==." +
		"OROYDYIdEouUWfUEgzXXPf3Awhx1QUcTN8banzg6ZinC-5d3wsJbFpOZEM2zzlLEVE02P5EPjNChmaQ2vFcO-DUoH4e8rUXpInP317dmRIk-WWMmsh72OrcN42b9jYUpMFI6V6TYA6o1Xu86MSPWBAQY-4qD-bUeYA9HTVwV5K9zPK8lGk-XYcno0xmEi1Uhu0YkLXBTYy4Kmj1ql7Ccg4NFBbvD3DmkxP0rRNFd-htNtfoUuB_J-d5diPp8-_5cL35gEQrsQK-IARke5dpGzq9zbdVk_gZYdgLislx5M2c4d44cbgWth9cFXdaOtWn6Nz7uy5-3eRvusXWMI68dBw"

	decodedToken, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyPublicRaw)
		if err != nil {
			panic(err)
		}
		return publicKey, nil
	},
	)

	if err != nil {
		panic(err)
	} else {
		decodedClaims := decodedToken.Claims.(*MyClaims)
		fmt.Println("userId:", decodedClaims.UserId)
	}

}
