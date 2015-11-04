package authentication

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"github.com/jonathonharrell/miri-ws-server/engine/models"
	"github.com/jonathonharrell/miri-ws-server/engine/settings"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type JWTAuthenticationBackend struct {
	Key []byte
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var (
	authBackendInstance *JWTAuthenticationBackend = nil
	blacklist                                     = make(map[string]bool)
)

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{[]byte(settings.GetEnv().JWTSecretKey)}
	}

	return authBackendInstance
}

func (backend *JWTAuthenticationBackend) GenerateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(settings.GetEnv().JWTExpirationDelta)).Unix()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = userId
	tokenString, err := token.SignedString(backend.Key)
	if err != nil {
		logger.Write.Error(err.Error())
		return "", err
	}
	return tokenString, nil
}

func (backend *JWTAuthenticationBackend) Authenticate(user *parameters.User) (models.User, error) {
	existing := models.User{}
	err := database.GetDB().C("users").Find(bson.M{"email": user.Email}).One(&existing)

	if err != nil { // no existing user
		return models.User{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(existing.HashedPassword), []byte(user.Password)) == nil {
		return existing, nil
	} else {
		return models.User{}, errors.New("Authentication Error: Passwords did not match")
	}
}

func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	blacklist[tokenString] = true
	return nil
}

func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
	if _, ok := blacklist[token]; ok {
		return true
	}

	return false
}
