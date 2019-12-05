package main

import (
	"log"

	"client/login"
	"client/msg"
)
type ClientConf struct {

}

func main(){
	log.Printf("-----------------test msg-----------------")
	msgTest := new(msg.Msg)
	msgTest.TestTCP()
	msgTest.TestWebsocket()

	log.Printf("-----------------test login-----------------")
	loginTest := new(login.Login)
	loginTest.Connect()
}