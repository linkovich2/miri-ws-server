package controllers

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jonathonharrell/miri-ws-server/engine/core/authentication"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func getUserId(r *http.Request) (string, error) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return authBackend.Key, nil
	})
	if err != nil {
		return "", err
	}

	if !bson.IsObjectIdHex(token.Claims["sub"].(string)) {
		return "", errors.New("Invalid hex value for user ID found. Hacking attempt?")
	}

	return token.Claims["sub"].(string), nil
}
