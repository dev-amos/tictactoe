// Package real contains a real-human-controlled tic tac toe player
package real

import (
	"github.com/dev-amos/tictactoe/board"
	"github.com/dev-amos/tictactoe/player"
)

type realPlayer struct {
	name   string
	symbol board.BoxContent
}

type NewPlayerParams struct {
	Name   string
	Symbol board.BoxContent
}

// NewPlayer creates a human-controlled player.
func NewPlayer(params NewPlayerParams) player.Player {
	return realPlayer{
		name:   params.Name,
		symbol: params.Symbol,
	}
}

// GetName returns name of player.
func (rp realPlayer) GetName() string {
	return rp.name
}

// GetSymbol returns symbol used by player to fill the boxes in tic tac toe
func (rp realPlayer) GetSymbol() board.BoxContent {
	return rp.symbol
}
