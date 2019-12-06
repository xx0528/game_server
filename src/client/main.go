package main

import (
	"log"
	"bufio"
	"os"
	"time"
	"sync"
	"encoding/binary"
	
	// "client/msg"
	"client/login"

	"github.com/gorilla/websocket"
)

var (
	client *login.Login
)
func main(){

	// log.Printf("-----------------test msg-----------------")
	// msgTest := new(msg.Msg)
	// msgTest.TestTCP()
	// msgTest.TestWebsocket()

	log.Printf("-----------------test login-----------------")
	client := new(login.Login)
	client.OnLogin()

	// InputCmd()

	// HeartBeat(client)
}

func InputCmd() {

	// 创建一个map 指定key为string类型 val为int类型
    counts := make(map[string]int)
    // 从标准输入流中接收输入数据
    input := bufio.NewScanner(os.Stdin)

    log.Printf("Please type in something:\n")

    // 逐行扫描
    for input.Scan() {
        line := input.Text()

        // 输入bye时 结束
        if line == "bye" {
            break
        }

        // 更新key对应的val 新key对应的val是默认0值
        counts[line]++
    }

    // 遍历map统计数据
    for line, n := range counts {
        log.Printf("%d : %s\n", n, line)
    }
}

func HeartBeat(client *login.Login) {
	
	data := []byte(`{
		"HeartBeat": {
			"PID": 1
		}
	}`)
	
	m := make([]byte, 2+len(data))

	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	
	client.Conn.WriteMessage(websocket.TextMessage, data)

	// go func() {
	// 	for {
	// 		client.Conn.WriteMessage(websocket.TextMessage, data)
	// 		log.Printf("发送心跳包---")

	// 		time.Sleep(1)
	// 	}
	// }()

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 20; i = i + 1 {
		wg.Add(1)
		go func() {
			client.Conn.WriteMessage(websocket.TextMessage, data)
			time.Sleep(3)
			log.Printf("send -- -")
		}()
		wg.Done()
	}
	log.Printf("心跳包 发送---- ")
    wg.Wait()
}