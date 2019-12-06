package msg

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
