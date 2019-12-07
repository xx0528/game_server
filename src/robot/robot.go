package main

import (
	"log"
	// "time"
	// "bytes"
	// "math/rand"
	// "strconv"
	// "encoding/binary"
	// "reflect"
	// "unsafe"
	// "encoding/json"
	"robot/msg"
	"robot/utils"

	"gopkg.in/mgo.v2/bson"
)

type Robot struct {
	objid 		string
	GameID		int
	SeatID		int
	RoomID		int
	NickName	string
	Money		int64
	wsConn		*msg.WSConn
}

func CreateRobot() *Robot {
	robot := new(Robot)
	return robot
}

func (robot *Robot) Init() {

	robot.wsConn = msg.CreateConn()
	robot.wsConn.Connect("127.0.0.1:3564")
	robot.wsConn.Register(&msg.PlayerMsg{}, robot.OnPlayerMsg)
	robot.wsConn.Register(&msg.LoginRet{}, robot.OnLogin)
	robot.wsConn.Register(&msg.HeartBeat{}, robot.OnHeartBeat)

	robot.objid = bson.NewObjectId().Hex()
	robot.NickName = utils.GetFullName()
}

func (robot *Robot) GetID() string {
	return robot.objid
}

func (robot *Robot) GetName() string {
	return robot.NickName
}

func (robot *Robot) OnPlayerMsg(args []interface{}) {
	m := args[0].(*msg.PlayerMsg)
	robot.Debug("recv OnPlayerMsg msg --- ", m.Cmd)
}

func (robot *Robot) HeartBeat() {
	// for {
	// 	time.Sleep(time.Second*4)
		data := []byte(`{
			"HeartBeat": {
				"PID": 1
			}
		}`)
		robot.SendMsg(data)
	// }
}
func (robot *Robot) OnHeartBeat(args []interface{}) {
	m := args[0].(*msg.HeartBeat)
	robot.Debug("recv OnHeartBeat msg --- ", m.PID)
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

func (robot *Robot) OnLogin(args []interface{}) {
	m := args[0].(*msg.LoginRet)
	robot.Debug("recv LoginRet msg --- ", m.Cmd)
}

func (robot *Robot) GetMoney() int64{
	return robot.Money
}

func (robot *Robot) SendMsg(data []byte) {
	robot.wsConn.WriteMsg(data)
}

func (robot *Robot) Logout() {
	robot.wsConn.Close()
}

func (robot *Robot) Debug(format string, a ...interface{}) {
	log.Printf(robot.GetName() + " : " + format, a... )
}