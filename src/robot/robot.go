package main

import (
	"net/url"
	"log"

	"encoding/binary"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/websocket"
)

type Robot struct {
	objid 		string
	GameID		int
	SeatID		int
	RoomID		int
	NickName	string
	Money		int64
	conn		*websocket.Conn
}

func CreateRobot() *Robot {
	robot := new(Robot)
	return robot
}

func (robot *Robot) connect() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:3564", Path: ""}
	log.Printf("connecting to %s", u.String())
	var err error
	robot.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
}

func (robot *Robot) Init() {
	robot.objid = bson.NewObjectId().Hex()
	robot.connect()
}

func (robot *Robot) Login() {
	data := []byte(`{
		"LoginMsg": {
			"Cmd": "login",
			"Account": "xx",
			"Password": "123456"
		}
	}`)
	robot.SendMsg(data)
}

func (robot *Robot) GetMoney() int64{
	return robot.Money
}

func (robot *Robot) SendMsg(data []byte) {
	m := make([]byte, 2+len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	robot.conn.WriteMessage(websocket.TextMessage, data)
}

func (robot *Robot) Logout() {
	robot.conn.Close()
}
