package msg

import (
	"github.com/name5566/leaf/network/json"
	//"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = json.NewProcessor()
	//Processor = protobuf.NewProcessor()
)

func init() {
	Processor.Register(&HeartBeat{})
	Processor.Register(&LoginMsg{})
	Processor.Register(&LoginRet{})
	Processor.Register(&GameMsg{})
	Processor.Register(&PlayerMsg{})
}

type HeartBeat struct {
	PID		int		`json:"pid"`
}

type LoginMsg struct {
	Cmd 		string	`json:"cmd"`
	Account		string	`json:"account"`
	Password	string	`json:"password"`
}

type LoginRet struct {
	Code		int			`json:"code"`
	ErrorMsg	string		`json:"errormsg"`
	Cmd 		string		`json:"cmd"`
	Data 		interface{}	`json:"data"`
}

type LoginAns struct {
	UserCheck string	`json:"usercheck"`
}

type PlayerMsg struct {
	Code		int			`json:"code"`
	ErrorMsg	string		`json:"errormsg"`
	Cmd			string		`json:"cmd"`
	Ans			interface{}	`json:"data"`
}

type GameMsg struct {
	Cmd string 		`json:"cmd"`
	Req interface{}	`json:"req"`
}
