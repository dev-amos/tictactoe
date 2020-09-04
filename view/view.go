// Package view defines the "frontend" for the game tic tac toe
package view

import (
	"github.com/dev-amos/tictactoe/board"
)

type GetUserToSelectBoxParams struct {
	PlayerName   string
	PlayerSymbol board.BoxContent
}

// The model that is responsible for taking inputs from the players
type View interface {
	DeclareDraw()
	DeclareWinner(playerName string)
	PrintBoard(b board.Board)
	GetDimensions() (int, error)
	GetUserName(playerCount int) (string, error)
	GetUserToSelectBox(p GetUserToSelectBoxParams) (int, error)
}
