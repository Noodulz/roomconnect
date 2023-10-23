package token

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lesterfernandez/roommate-finder/server/data"
	"net/http"
	"time"
)

var JWTKey = []byte("secret")

//=====================================================================================
// Encode JWT into JSON, send it as response
//=====================================================================================

func SendJWT(w http.ResponseWriter, jwt string) {
	w.Header().Set("Application-Type", "application/json")
	json.NewEncoder(w).Encode(data.TokenMessage{Token: jwt})
}

//=====================================================================================
// Create JWT
//=====================================================================================

func CreateJWT(username string) (string, error) {

	expirationTime := time.Now().Add(time.Minute * 1)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)

	return tokenString, err
}

//========================================================================================================
// Verify that a JWT is valid, and return the registeredClaims
//========================================================================================================

func VerifyJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {

	//Verify the correct signing method is used (HS256)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, err
	}

	// Verify if token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return token, nil
}