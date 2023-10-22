package common

import (
	"fmt"
	"time"

	"enigmacamp.com/enigma-laundry-apps/config"
	"github.com/dgrijalva/jwt-go"
)

// Fungsi untuk membuat token JWT
func CreateJWTToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token kadaluarsa dalam 1 jam
	})
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = username
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token kadaluarsa dalam 1 jam

	tokenString, _ := token.SignedString(config.TokenCfg.JWTSignatureKey)
	return tokenString
}

// Fungsi untuk memverifikasi token JWT
func VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return config.TokenCfg.JWTSignatureKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return claims, nil
	} else {
		fmt.Println(err)
	}

	return nil, nil

}
