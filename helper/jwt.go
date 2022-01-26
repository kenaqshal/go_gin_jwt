package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte("jwt_go")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

//func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
//	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
//		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
//			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
//
//		}
//		return []byte(service.secretKey), nil
//	})
//
//}
