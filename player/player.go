// Package player defines a tic tac toe player
package player

import (
	"github.com/dev-amos/tictactoe/board"
)

// Player of tic-tac-toe.
type Player interface {
	GetName() string
	GetSymbol() board.BoxContent
}
