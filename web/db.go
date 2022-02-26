package web

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/mgo.v2"
)

func ConnectWithMgo() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	fmt.Printf("USERNAME: %s, PASSWORD: %s\n",
		username, password)
	cred := &mgo.Credential{
		Username: username,
		Password: password,
	}
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Println("err: connect failed")
	}

	err = session.Login(cred)
	if err != nil {
		fmt.Println("err: login failed")
	}

}
