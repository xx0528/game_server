package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.HeartBeat{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginMsg{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.GameMsg{}, game.ChanRPC)
}
