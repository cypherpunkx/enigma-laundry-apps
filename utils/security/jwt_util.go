package security

import (
	"fmt"
	"time"

	"enigmacamp.com/enigma-laundry-apps/config"
	"enigmacamp.com/enigma-laundry-apps/model"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *model.UserCredential) (string, error) {

	now := time.Now().UTC()
	end := now.Add(config.TokenCfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.TokenCfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(config.TokenCfg.JWTSigningMethod, claims)
	ss, err := token.SignedString(config.TokenCfg.JWTSignatureKey)

	if err != nil {
		return "", fmt.Errorf("Failed to create access token : %s", err.Error())
	}

	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != config.TokenCfg.JWTSigningMethod {
			return nil, fmt.Errorf("Invalid token string method")
		}
		return config.TokenCfg.JWTSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Invalid parse token : %s", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid || claims["iss"] != config.TokenCfg.ApplicationName {
		return nil, fmt.Errorf("Invalid Token MapClaims")
	}

	return claims, nil
}
