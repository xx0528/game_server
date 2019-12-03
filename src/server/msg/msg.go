package msg

import (
	"github.com/name5566/leaf/network/json"
	//"github.com/name5566/leaf/network/protobuf"
)

var (
	Processor = json.NewProcessor()
	//Processor = json.NewProcessor()
)
// var Processor network.Processor

func init() {
	Processor.Register(&Hello{})
}

type Hello struct {
	Name string
}