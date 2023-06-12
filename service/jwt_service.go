package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/katerji/UserAuthKit/model"
	"os"
)

type JWTService struct{}

func (JWTService) VerifyToken(token string) (model.User, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return model.User{}, errors.New("error parsing token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		jsonClaims, err := json.Marshal(claims)
		if err != nil {
			return model.User{}, errors.New("error parsing token")
		}
		var user model.User
		if err := json.Unmarshal(jsonClaims, &user); err != nil {
			return model.User{}, errors.New("error parsing token")
		}
		return user, nil
	}
	return model.User{}, errors.New("invalid token")

}

func (JWTService) CreateJwt(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
