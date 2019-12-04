package internal

type playerMgr struct {
	PlayerMap map[string]*Player
}

func NewPlayerMgr() *playerMgr {
	mgr := new(playerMgr)
	mgr.Init()
	return mgr
}

func (mgr *playerMgr) Init() {
	mgr.PlayerMap = make(map[string]*Player)
}

func (mgr *playerMgr) Get(id string) *Player {
	mgr.PlayerMap[player.objid] = player
}

func (mgr *playerMgr) AddPlayer(player *Player) {
	mgr.PlayerMap[player.objid] = player
}

func (mgr *playerMgr) DelPlayer(objid string) {
	delete(mgr.PlayerMap, objid)
}

func (mgr *playerMgr) Close() {
	for _, player := range mgr.PlayerMap {
		player.SaveSync()
	}
}