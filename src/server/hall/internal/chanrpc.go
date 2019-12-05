package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("OnLogin", rpcOnLogin)
}

func rpcOnLogin(args []interface{}) {
	a := args[0].(game.Agent)
	log.Debug("close agent")

	userdata := a.UserData()
	if nil == userdata {
		return
	}

	info := userdata.(*base.AccountInfo)
	player := PlayerManager.Get(info.ObjID)
	if nil != player {
		player.agent.Close()
		player.agent = a
		return
	}
	
	mgodb.Get(base.DBTask{info.ObjID, base.DBNAME, base.PLAYERSET, "_id", base.BsonObjectID(info.ObjID), CreatePlayer(), func(param interface{}, err error) {		
		player := param.(*Player)
		player.objid = info.ObjID
		player.agent = a
		if "" == player.Account {
			player.InitData(info.Account)
			player.Save()
		}

		player.OnLogin()
		PlayerManager.AddPlayer(player)

		player.SendMsg(0, "login", &msg.LoginAns{info.ObjID})
	}})
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("hall new agent")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("hall close agent")
	userdata := a.UserData()
	if nil == userdata {
		return
	}

	info := userdata.(*base.AccountInfo)
	player := PlayerManager.Get(info.ObjID)
	if nil != player {
		player.OnLogout()
	}

	PlayerManager.DelPlayer(info.ObjID)
}
