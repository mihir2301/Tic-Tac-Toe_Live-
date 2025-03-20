package game

import "errors"

type Game struct {
	ID       string
	Board    *Board
	Players  []string
	NextTurn string
}

func NewGame(id string) *Game {
	return &Game{
		ID:       id,
		NextTurn: "X",
	}
}

//Add palyer to the Game

func (g *Game) AddPlayer(playerID string) error {
	if len(g.Players) >= 2 {
		return errors.New("Game is full")
	}
	g.Players = append(g.Players, playerID)
	return nil
}

//Make a move

func (g *Game) MakeAMove(row, col int, symbol string) error {
	if g.Board.Grid[row][col] != "" {
		return errors.New("cell already occupied")
	}

	if symbol != g.NextTurn {
		return errors.New("not your turn")
	}

	g.Board.Grid[row][col] = symbol
	g.SwitchTurn()

	return nil
}

//switch turns

func (g *Game) SwitchTurn() {
	if g.NextTurn == "X" {
		g.NextTurn = "O"
	} else {
		g.NextTurn = "X"
	}
}
