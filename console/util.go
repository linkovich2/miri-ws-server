package main

import (
	"fmt"
	"os"

	"github.com/jonathonharrell/miri-ws-server/engine"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	cmd := os.Args[1]
	args := os.Args[2:]
	fmt.Printf("%v\n", cmd)
	fmt.Printf("%v\n", args)

	if cmd == "admin" {
		Admin(args[0])
	}
}

func Admin(email string) {
	engine.LoadEnv()
	env := engine.GetEnv()
	engine.ConnectToDatabase(env.DBHost, env.DBName)
	defer engine.CloseDatabaseConnection()

	db := engine.GetDB()
	c := db.C("users")

	q := bson.M{"email": email}
	change := bson.M{"$set": bson.M{"isadmin": true}}
	err := c.Update(q, change)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully set user to admin role.")
}
