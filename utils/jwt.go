package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	Encode(data interface{}) (string, error)
	Decode(token string) (*jwt.Token, error)
}

type jwtService struct {
	secret string
	issuer string
}

func JWTAuthService() JWTService {
	conf := GetConfig()
	return &jwtService{
		secret: conf.Jwt_Secret,
		issuer: conf.Jwt_Issuer,
	}
}

func (service *jwtService) Encode(data interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  service.issuer,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"data": data,
	})
	return token.SignedString([]byte(service.secret))
}

func (service *jwtService) Decode(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %d", token.Header["alg"])
		}
		return []byte(service.secret), nil
	})
}
