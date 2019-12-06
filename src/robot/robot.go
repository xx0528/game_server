package main

import (
	"net/url"
	"log"
	"encoding/json"
	"encoding/binary"
	"errors"
	"robot/msg"
	"reflect"

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
	msgInfo 	map[string]*MsgInfo
}

type MsgInfo struct {
	msgType       reflect.Type
	msgHandler    MsgHandler
}

type MsgHandler func([]interface{})

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
		return
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			log.Printf("recv --------- ")
			_, message, err := robot.conn.ReadMessage()
			if err != nil {
				log.Println("read:---", err)
				return
			}
			msg, err := robot.Unmarshal(message)

			msgType := reflect.TypeOf(msg)
			if msgType == nil || msgType.Kind() != reflect.Ptr {
				log.Println("json message pointer required")
				return
			}
			msgID := msgType.Elem().Name()
			i, ok := robot.msgInfo[msgID]
			if !ok {
				log.Println("message not registered", msgID)
			}
			if i.msgHandler != nil {
				i.msgHandler([]interface{}{msg})
			}
			// aa := message.(*msg.PlayerMsg)
			log.Printf("recv:--- %s", msgID)

		}
	}()
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (robot *Robot) Register(msg interface{}, msgHandler MsgHandler) string {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("json message pointer required")
	}
	msgID := msgType.Elem().Name()
	if msgID == "" {
		log.Fatal("unnamed json message")
	}
	if _, ok := robot.msgInfo[msgID]; ok {
		log.Fatal("message %v is already registered", msgID)
	}

	i := new(MsgInfo)
	i.msgType = msgType
	i.msgHandler = msgHandler
	robot.msgInfo[msgID] = i

	return msgID
}

func (robot *Robot) Init() {

	robot.Register(&msg.PlayerMsg{}, robot.OnPlayerMsg)

	robot.objid = bson.NewObjectId().Hex()
	robot.connect()
	// robot.Run()
}

func (robot *Robot) OnPlayerMsg(args []interface{}) {
	m := args[0].(*msg.PlayerMsg)
	log.Printf("recv OnPlayerMsg msg --- ", m)
}

func (robot *Robot) Run() {
	for {
		_, data, err := robot.conn.ReadMessage()
		if err != nil {
			log.Printf("read message: %v", err)
			break
		}

		msg, err := robot.Unmarshal(data)
		if err != nil {
			log.Printf("unmarshal message error: %v", err)
			break
		}
		log.Printf("recv msg --- ", msg)
	}
}

// goroutine safe
func (robot *Robot) Unmarshal(data []byte) (interface{}, error) {
	var m map[string]json.RawMessage
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	if len(m) != 1 {
		return nil, errors.New("invalid json data")
	}

	for msgID, data := range m {
		log.Printf("msgID : ", msgID, " data - ", data )
	}

	panic("bug")
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
