package game

import (
	"errors"
	"fmt"
)

type Board struct {
	Grid [3][3]string
}

func NewBoard() *Board {
	return &Board{}
}
func (b *Board) MakeAMove(row, col int, symbol string) error {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return errors.New("invalid cell position")
	}
	if b.Grid[row][col] != "" {
		return errors.New("cell already occupied")
	}
	b.Grid[row][col] = symbol
	return nil
}

//Check if there is a Winner

func (b *Board) CheckWinner() string {
	WinningPattern := [][][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, pattern := range WinningPattern {
		a, d, c := pattern[0], pattern[1], pattern[2]

		if b.Grid[a[0]][a[1]] != "" && b.Grid[a[0]][a[1]] == b.Grid[d[0]][d[1]] &&
			b.Grid[d[0]][d[1]] == b.Grid[c[0]][c[1]] {
			return b.Grid[a[0]][a[1]]
		}
	}
	return ""
}

//check if there is a draw

func (b *Board) CheckDraw() bool {
	for _, row := range b.Grid {
		for _, col := range row {
			if col == "" {
				return false
			}
		}
	}
	return true
}

//displayboard prints the board in console for debugging

func (b *Board) DisplayBoard() {
	for _, row := range b.Grid {
		fmt.Println(row)
	}
	fmt.Println()
}
