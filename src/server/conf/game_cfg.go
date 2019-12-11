package conf

import (
	"github.com/name5566/leaf/log"
	"encoding/json"
	"io/ioutil"
)

var (
	GameCfgs []*GameCfg
) 

type GameCfg struct {
	gameID		int
	name		string
	rooms		[]*RoomCfg
}

type RoomCfg struct {
	configID 		int
    baseScore 		int
    topScore 		int
    tax 			int
    enterMoney 		int
    desc 			string
}

func init() {
	data, err := ioutil.ReadFile("conf/game.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &GameCfgs)
	if err != nil {
		log.Fatal("%v", err)
	}
}