package internal

import (
	"github.com/name5566/leaf/util"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

	"server/base"
	"server/msg"
)

const (
	INTATTR_COIN  		= 0 //金币
	INTATTR_POWER 		= 1 //体力

	INTATTR_MAX 		= 10
)

const (
	STRATTR_NICK = 0 //昵称
	STRATTR_ICON = 1 //头像

	STRATTR_MAX = 2
)


type Player struct {
	objid   		string
	agent   		gate.Agent
	UID 			int64				
	Account 		string						
	IntAttr 		[]int    //整型属性
	StrAttr 		[]string //字符串属性
}

func CreatePlayer() *Player {
	player := new(Player)
	player.IntAttr = make([]int, INTATTR_MAX)
	player.StrAttr = make([]string, STRATTR_MAX)
	return player
}

func (p *Player) GetIntAttr(index int) int {
	if index < 0 || index >= INTATTR_MAX {
		return 0
	}
	return p.IntAttr[index]
}

func (p *Player) SetIntAttr(index, val int) {
	if index < 0 || index >= INTATTR_MAX {
		return
	}

	p.IntAttr[index] = val
}

func (p *Player) GetStrAttr(index int) string {
	if index < 0 || index >= STRATTR_MAX {
		return ""
	}
	return p.StrAttr[index]
}

func (p *Player) SetStrAttr(index int, val string) {
	if index < 0 || index >= STRATTR_MAX {
		return
	}

	p.StrAttr[index] = val
}

func (p *Player) InitData(account string) {
	p.Account = account
	p.UID     = uidbuilder.GenerateUID()
}

func (p *Player) SendMsg(ret int, cmd string, ans interface{}) {
	errmsg := ""

	message := &msg.PlayerMsg{ret, errmsg, cmd, ans}
	p.agent.WriteMsg(message)
}

//登陆
func (p *Player) OnLogin() {
}

func (p *Player) OnLogout() {
	p.Save()

	TimerMgr.RmvAllTimer(p)
}

//保存玩家数据
func (p *Player) Save() {
	mgodb.Set(base.DBTask{p.objid, base.DBNAME, base.PLAYERSET, "_id", base.BsonObjectID(p.objid), util.DeepClone(p), func(param interface{}, err error) {
		if nil != err {
			log.Error("save playerdata failed:", p.objid)
		}
	} })
}

//同步保存玩家数据
func (p *Player) SaveSync() {
	if nil != mgodb.SetSync(base.DBNAME, base.PLAYERSET, "_id", base.BsonObjectID(p.objid), p) {
		log.Error("save playerdata failed:", p.objid)
	}
}