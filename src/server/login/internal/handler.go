package internal

import (
	"reflect"
	"server/base"
	"server/msg"
	"server/hall"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
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
	mgodb.Get(base.DBTask{m.Account, base.DBNAME, base.ACCOUNTSET, "account", m.Account, &base.AccountInfo{}, func(param interface{}, err error) {
		info := param.(*base.AccountInfo)
		if "" == info.Account {
			info.Account = m.Account
			info.Password = m.Password
			info.ObjID = bson.NewObjectId().Hex()
			mgodb.Set(base.DBTask{info.Account, base.DBNAME, base.ACCOUNTSET, "account", m.Account, info, nil})
		}

		if info.Password != m.Password {
			a.WriteMsg(&msg.LoginRet{1, "", "login", nil})
			return
		}

		a.SetUserData(info)
		skeleton.AsynCall(hall.ChanRPC, "OnLogin", a, func(err error) {
			if nil != err {
				log.Error("login failed: ", info.ObjID, " ", err.Error())
				a.WriteMsg(&msg.LoginRet{-1, "", "login", nil})
				return
			}
		})
	}})
	return
}