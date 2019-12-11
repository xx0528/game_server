package internal

import (
	"math/rand"
)

const (
	ROOM_STATE_WAIT		= 0//等待状态				
	ROOM_STATE_GAME		= 1//游戏状态
	ROOM_STATE_DELETE	= 2//删除状态
	ROOM_STATE_UNINIT	= 3//未初始化	
)
var(
	S_ID = 1
)
type Room struct {
	ID				int64
	stateID			int
	enterMoney		int64
	baseScore		int
	topScore		int
	configID		int
	round			int

	seatList		map[string]*Seat
	playerList		map[string]*Player
}

func (room *Room) Init() {
	rand.Seed(time.Now().UnixNano())

	room.ID = S_ID
	S_ID = S_ID + 1
	round = 1
	baseScore = 0
	topScore = 0
	stateID = ROOM_STATE_UNINIT
}

func (room *Room) GetID() int64 {
	return room.ID
}

func (room *Room) GetSeat(seatID string) *Seat {
	return room.seatList[seatID]
}

func (room *Room) AddSeat(seatID string, seat *Seat) {
	room.seatList[seatID] = seat
}

func (room *Room) GetPlayer(playerID string) *Player {
	return room.playerList[playerID]
}

func (room *Room) AddPlayer(playerID string, player *Player) {
	room.playerList[playerID] = player
}

func (room *Room) Quit() {
	
}