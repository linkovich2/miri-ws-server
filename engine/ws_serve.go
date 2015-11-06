package engine

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jonathonharrell/miri-ws-server/engine/database"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"github.com/jonathonharrell/miri-ws-server/engine/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"stablelib.com/v1/uniuri"
	"strconv"
)

type ConnectionData struct {
	User      *models.User
	Character *models.Character
}

func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			} else {
				return []byte(env.JWTSecretKey), nil
			}
		})

		if err == nil && token.Valid {
			if !bson.IsObjectIdHex(token.Claims["_id"].(string)) {
				logger.Write.Error("Invalid hex value for user ID found in auth token.")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			u, err := getUser(token.Claims["_id"].(string))
			if err != nil {
				logger.Write.Error("User not found for ID", token.Claims["_id"].(string))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ws, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				logger.Write.Error(err.Error())
				return
			}

			c := &Connection{send: make(chan []byte, 256), webSocket: ws, ID: uniuri.New(), IsAdmin: u.IsAdmin}
			Hub.register <- c
			go c.writePump()
			c.readPump()
		} else {
			logger.Write.Error("%v", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	addr := ":" + strconv.Itoa(env.Port)
	go http.ListenAndServe(addr, nil)
}

func getUser(userId string) (*models.User, error) {
	session := database.GetSession() // connect
	db := session.DB(env.DBName)
	defer session.Close()

	u := models.User{}
	err := db.C("users").Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&u)
	return &u, err
}
