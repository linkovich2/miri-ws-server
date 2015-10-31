package authentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/models"
	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/settings"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"errors"
)

type JWTAuthenticationBackend struct {
	Key []byte
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			Key: []byte("i23k8jnTghdfadGGrt32hgSGH42zSD53HaraaR48990A"), // @todo move to env var
		}
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
		panic(err) // @todo this should probably just call logger
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

func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	// redisConn := redis.Connect()
	// return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims["exp"]))
	// @todo remove token?
	return nil
}

func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
	// @todo check if token exists in mongo ?
	// redisConn := redis.Connect()
	// redisToken, _ := redisConn.GetValue(token)
	//
	// if redisToken == nil {
	// 	return false
	// }
	//
	// return true

	return false
}
