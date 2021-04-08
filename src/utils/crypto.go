package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	return err == nil
}

func CreateToken(email string) (string, error) {
	secret := "SETASECRETHERESETASECRETHERESETASECRETHERESETASECRETHERESETASECRETHERE"
	if secret == "" {
		log.Fatal("$SECRET must be set")
	}

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour / 2).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(secret)) // SECRET
	if err != nil {
		return "token creation error", err
	}
	return token, nil
}

func ValidateToken(tokenString string) bool {
	secret := "SETASECRETHERESETASECRETHERESETASECRETHERESETASECRETHERESETASECRETHERE"
	if secret == "" {
		log.Fatal("$SECRET must be set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
