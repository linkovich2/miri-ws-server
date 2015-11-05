package services

import (
	"encoding/json"
	"net/http"

	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/core/authentication"
	db "github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"github.com/jonathonharrell/miri-ws-server/engine/models"

	"github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type errorResponse struct {
	Error string `json:"error"`
}

func Login(requestUser *parameters.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	if user, err := authBackend.Authenticate(requestUser); err == nil {
		token, err := authBackend.GenerateToken(user.ID.Hex())
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func CreateUser(requestUser *parameters.User) (int, []byte) {
	session, dbName := db.GetSession() // connect
	database := session.DB(dbName)
	defer session.Close()

	authBackend := authentication.InitJWTAuthenticationBackend()

	existing := &models.User{}
	err := database.C("users").Find(bson.M{"email": requestUser.Email}).One(&existing)

	if err == nil { // checking for existing user
		response, _ := json.Marshal(errorResponse{"A user already exists with that email."})
		return http.StatusBadRequest, response
	}

	if !govalidator.IsEmail(requestUser.Email) {
		response, _ := json.Marshal(errorResponse{"Not a valid email"})
		return http.StatusBadRequest, response
	}

	if len(requestUser.Password) < 6 {
		response, _ := json.Marshal(errorResponse{"Password must be at least 6 characters long."})
		return http.StatusBadRequest, response
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(requestUser.Password), 10)
	i := bson.NewObjectId()
	user := &models.User{ID: i, Email: requestUser.Email, HashedPassword: string(hashed), IsAdmin: false}
	database.C("users").Insert(user)
	logger.Write.Info("New User Created: %s", user.ID.Hex())

	token, err := authBackend.GenerateToken(user.ID.Hex())
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	} else {
		response, _ := json.Marshal(parameters.TokenAuthentication{token})
		return http.StatusOK, response
	}
}

func RefreshToken(requestUser *parameters.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.Email)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.Key, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}

func ForgotPassword(email string) (int, []byte) {
	return http.StatusTeapot, []byte("")
}

func ResetPassword(pwd string) (int, []byte) {
	return http.StatusTeapot, []byte("")
}
