package main

import (
	"log"
	"os"
	"os/signal"
)

var (
	robotMgr = NewRobotMgr()
)
func main() {
	log.Printf("robot mgr init------")
	robotMgr.Init()

	robot := CreateRobot()
	log.Printf("create robot")
	robot.Init()
	log.Printf("init robot")
	robot.Login()
	log.Printf("robot login")
	robotMgr.Add(robot)
	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Printf("close robot (signal: %v)", sig)

	robotMgr.Close()
}