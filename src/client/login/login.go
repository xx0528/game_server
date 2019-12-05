package login

import (
	"encoding/binary"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Login struct {
	conn *websocket.Conn
}

func (lg *Login)Connect() {

	// var addr = flag.String("addr", "127.0.0.1:3564", "http service address")

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:3564", Path: ""}
	log.Printf("connecting to %s", u.String())
	var err error
	lg.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer lg.conn.Close()
}

func (lg *Login)SendLogin() {

	data := []byte(`{
			"Hello": {
				"Name": "leaf websocket"
			}
		}`)
	
	m := make([]byte, 2+len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	log.Fatal("conn:", lg.conn)
	lg.conn.WriteMessage(websocket.TextMessage, data)
}