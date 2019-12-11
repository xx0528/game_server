package internal

import (
	"server/conf"
)

const (
	GAME_DDZ		= 1
	GAME_ZJH		= 2
	GAME_FISH		= 3
	GAME_DZPK		= 4

)

type GameMgr struct {
	games		map[string]*Game
}

func NewGameMgr() *GameMgr {
	gameMgr := new(GameMgr)
	gameMgr.Init()
	return gameMgr
}

func (mgr *GameMgr) Init() {
	mgr.games = make(map[string]*Game)

	mgr.CreateGames()
}

func (mgr *GameMgr) CreateGames(){
	for _, cfg := range conf.GameCfgs {
		switch cfg.gameID {
		case GAME_DDZ:
			g := NewGame(cfg)
		case GAME_ZJH:
		case GAME_FISH:  
		case GAME_DZPK:
		default:
			g := NewGame(cfg)
		}
		mgr.games[g.GetID()] = g
	}

}