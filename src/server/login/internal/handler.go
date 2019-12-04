package internal

import (
	"reflect"
	"server/base"
	"server/game"
	"server/msg"
	"server/hall"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.LoginMsg{}, onLogin)
}

func onLogin(args []interface{}) {
	m := args[0].(*msg.LoginMsg)
	a := args[1].(gate.Agent)

	log.Debug("on recv login msg %v", m)
	mgodb.Get(base.DBTask{req.Account, base.DBNAME, base.ACCOUNTSET, "account", m.Account, &base.AccountInfo{}, func(param interface{}, err error) {
		info := param.(*base.AccountInfo)
		if "" == info.Account {
			info.Account = req.Account
			info.Password = req.Password
			info.ObjID = bson.NewObjectId().Hex()
			mgodb.Set(base.DBTask{info.Account, base.DBNAME, base.ACCOUNTSET, "account", req.Account, info, nil})
		}

		if info.Password != req.Password {
			a.WriteMsg(&msg.RetMsg{1, "", "login", nil})
			return
		}

		agent.SetUserData(info)
		skeleton.AsynCall(game.ChanRPC, "OnLogin", agent, func(err error) {
			if nil != err {
				log.Error("login failed: ", info.ObjID, " ", err.Error())
				a.WriteMsg(&msg.RetMsg{-1, "", "login", nil})
				return
			}
		})
	}})
	return 0
}