package msg

import (
	"log"
	"reflect"
	"encoding/json"
	"net/url"
	"errors"
	"fmt"
	"encoding/binary"

	"github.com/gorilla/websocket"
)

type WSConn struct {
	conn		*websocket.Conn
	msgInfo 	map[string]*MsgInfo
}

type MsgInfo struct {
	msgType       reflect.Type
	msgHandler    MsgHandler
}

type MsgHandler func([]interface{})

func CreateConn() *WSConn {
	conn := new(WSConn)
	conn.msgInfo = make(map[string]*MsgInfo)
	return conn
}

func (ws *WSConn) Connect(host string) {
	u := url.URL{Scheme: "ws", Host: host, Path: ""}
	var err error
	ws.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := ws.conn.ReadMessage()
			if err != nil {
				return
			}
			msg, err := ws.Unmarshal(message)

			msgType := reflect.TypeOf(msg)
			if msgType == nil || msgType.Kind() != reflect.Ptr {
				log.Println("json message pointer required")
				return
			}
			msgID := msgType.Elem().Name()
			i, ok := ws.msgInfo[msgID]
			if !ok {
				log.Println("message not registered", msgID)
			}
			if i.msgHandler != nil {
				i.msgHandler([]interface{}{msg})
			}
		}
	}()
}

func (ws *WSConn) WriteMsg(data []byte) {
	m := make([]byte, 2+len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	ws.conn.WriteMessage(websocket.TextMessage, data)
}

func (ws *WSConn) Register(msg interface{}, msgHandler MsgHandler) string {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("json message pointer required")
	}
	msgID := msgType.Elem().Name()
	if msgID == "" {
		log.Fatal("unnamed json message")
	}
	if _, ok := ws.msgInfo[msgID]; ok {
		log.Fatal("message %v is already registered", msgID)
	}

	i := new(MsgInfo)
	i.msgType = msgType
	i.msgHandler = msgHandler
	ws.msgInfo[msgID] = i

	return msgID
}


// goroutine safe
func (ws *WSConn) Unmarshal(data []byte) (interface{}, error) {
	var m map[string]json.RawMessage
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	if len(m) != 1 {
		return nil, errors.New("invalid json data")
	}
	for msgID, data := range m {
		i, ok := ws.msgInfo[msgID]
		if !ok {
			return nil, fmt.Errorf("message %v not registered", msgID)
		}
		// msg
		msg := reflect.New(i.msgType.Elem()).Interface()
		return msg, json.Unmarshal(data, msg)
	}

	return nil, nil
}

func (ws *WSConn) Close() {
	ws.conn.Close()
}
