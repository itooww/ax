package main

import (
	"fmt"
	"os"
	"encoding/json"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("please input JWT")
	}

	fmt.Println(os.Args)
	tokenString := os.Args[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	fmt.Println("Header")
	jsonHeader, err := json.MarshalIndent(token.Header, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonHeader))

	fmt.Printf("\nPayload\n")
	jsonClaims, err := json.MarshalIndent(token.Claims, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonClaims))

	fmt.Printf("\nSignature\n")
	fmt.Println(token.Signature)
}