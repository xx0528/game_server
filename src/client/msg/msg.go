package msg

import (
	"encoding/binary"
	"net"
	"flag"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Msg struct {
}

func (msg *Msg)TestTCP() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	log.Printf("connecting to tcp ")
	if err != nil {
		panic(err)
	}

	data := []byte(`{
			"Hello": {
				"Name": "leaf tcp"
			}
		}`)
	
	m := make([]byte, 2+len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	conn.Write(m)
}

func (msg *Msg)TestWebsocket() {
	var addr = flag.String("addr", "127.0.0.1:3564", "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: ""}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	data := []byte(`{
			"Hello": {
				"Name": "leaf websocket"
			}
		}`)
	
	m := make([]byte, 2+len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	c.WriteMessage(websocket.TextMessage, data)
}