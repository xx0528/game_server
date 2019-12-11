package internal

type PlayerMgr struct {
	PlayerMap map[string]*Player
}

func NewPlayerMgr() *PlayerMgr {
	mgr := new(PlayerMgr)
	mgr.Init()
	return mgr
}

func (mgr *PlayerMgr) Init() {
	mgr.PlayerMap = make(map[string]*Player)
}

func (mgr *PlayerMgr) Get(id string) *Player {
	return mgr.PlayerMap[id]
}

func (mgr *PlayerMgr) AddPlayer(player *Player) {
	mgr.PlayerMap[player.objid] = player
}

func (mgr *PlayerMgr) DelPlayer(objid string) {
	delete(mgr.PlayerMap, objid)
}

func (mgr *PlayerMgr) Close() {
	for _, player := range mgr.PlayerMap {
		player.SaveSync()
	}
}