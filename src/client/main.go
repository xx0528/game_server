package main

import (
	"log"
	"bufio"
	"os"
	
	// "client/msg"
	"client/login"

)

func main(){

	// log.Printf("-----------------test msg-----------------")
	// msgTest := new(msg.Msg)
	// msgTest.TestTCP()
	// msgTest.TestWebsocket()

	log.Printf("-----------------test login-----------------")
	loginTest := new(login.Login)
	loginTest.OnLogin()

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