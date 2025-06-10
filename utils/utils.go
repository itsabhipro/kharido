package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/itsabhipro/kharido/models"
	"github.com/itsabhipro/kharido/types"
)

func GenerateAuthToken(user *models.User) (string, error) {
	screte := os.Getenv("SECRETE")
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserId,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	return authToken.SignedString([]byte(screte))
}

func VerifyAuthToken(jwtstring string) (*types.AuthToken, error) {
	parsedAuthToken, err := jwt.Parse(jwtstring, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknow signing method")
		}
		return []byte(os.Getenv("SECRETE")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedAuthToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("something went wrong")
	}
	var authObject = types.AuthToken{
		UserId: claims["user_id"].(int64),
		Email:  claims["email"].(string),
		Role:   claims["role"].(string),
	}
	return &authObject, nil

}
