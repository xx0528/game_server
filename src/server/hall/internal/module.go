package internal

import (
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/log"

	"server/base"
	"server/conf"
	"server/db"
)

var (
	skeleton 	= base.NewSkeleton()
	ChanRPC  	= skeleton.ChanRPCServer
	mgodb		= new(db.Mongodb)
	PlayerMgr	= NewPlayerMgr()
	TimerMgr 	= NewTimerMgr()
	uidbuilder	= new(UidBuilder)
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	var err error
	mgodb, err = db.Dial(conf.Server.MgodbAddr, conf.Server.GameMgoConnNum, skeleton)
	if nil == mgodb {
		log.Error("dial mongodb failed:", conf.Server.MgodbAddr, " ", err.Error())
		return
	}

	mgodb.EnsureUniqueIndex(base.DBNAME, base.PLAYERSET, []string{"uid"})

	uidbuilder.Init()
}

func (m *Module) OnDestroy() {
	PlayerMgr.Close()
	mgodb.Close()
	log.Release("closed")
}
