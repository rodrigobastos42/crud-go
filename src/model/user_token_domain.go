package model

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *internalerrors.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.GetID(),
		"email": ud.GetEmail(),
		"name":  ud.GetName(),
		"age":   ud.GetAge(),
		"ext":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", internalerrors.NewInternalServerError(fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}
