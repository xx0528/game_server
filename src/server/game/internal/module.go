package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	gameMgr	 = NewGameMgr()
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	m.gameMgr.Init()
}

func (m *Module) OnDestroy() {

}
