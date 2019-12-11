package internal

func NewGame(cfg *GameCfg) *Game {
	game := new(Game)
	game.Init(cfg)
	return game
}

type Game struct {
	ID			int
	Name		string
	rooms		map[string]*Room
}

func (game *Game) Init(cfg *GameCfg) {
	game.rooms = make(map[string]*Room)
	game.ID = cfg.gameID
	game.Name = cfg.name
}

func (game *Game) GetID() int {
	return ID
}

func (game *Game) GetRoom(roomID string) *Room {
	return game.rooms[roomID]
}