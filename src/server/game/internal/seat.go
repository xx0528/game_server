package internal

const (
	SEAT_STATE_READY 	= 0		//准备
	SEAT_STATE_GAME		= 1		//游戏中
	SEAT_STATE_OFFLINE	= 2		//离线 托管
	SEAT_STATE_LEAVE 	= 3		//暂离
	SEAT_STATE_FOLD		= 4		//弃牌
)

type Seat struct {
	ID			int64
	roomID		int64
	room		*Room
	state		int
	playerID	int64
	coin		int64
}

func (seat *Seat)Init(room *Room) {
	rand.Seed(time.Now().UnixNano())
	seat.coin = 0	
	seat.ID = hInternal.GenerateUID()
	seat.room = room
	seat.roomID = room.GetID()
}