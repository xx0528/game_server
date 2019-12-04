package main

import (
	"encoding/binary"
	"net"
	"flag"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main(){
	TestTCP()
	TestWebsocket()
}

func TestTCP() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
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

var addr = flag.String("addr", "127.0.0.1:3564", "http service address")

func TestWebsocket() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
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