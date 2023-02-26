package helper

import (
	"errors"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user models.User, secret string) (string, error) {
  token := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, jwt.MapClaims{
    "foo": "bar",
  })

  tokenString, err := token.SignedString([]byte(secret))
  if err != nil {
    return "", err
  }

  return tokenString, nil
}
func ValidateToken(tokenString string, secret string) (*jwt.Token, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("Invalid Signing Method")
    }

    return secret, nil
  })

  if err != nil {
    return nil, err
  }

  return token, nil
}
