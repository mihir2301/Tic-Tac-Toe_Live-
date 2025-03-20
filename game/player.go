package game

type Player struct {
	Id     string
	Symbol string
	Score  int
}

//Creating a newPlayer instance

func NewPlayer(id string, symbol string) *Player {
	return &Player{
		Id:     id,
		Symbol: symbol,
		Score:  0,
	}
}
