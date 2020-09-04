// Package board represents the tic tac toe board with its boxes
package board

import (
	"errors"
)

var (
	ErrBoxOccupied         = errors.New("box is occupied and cannot be filled")
	ErrInvalidDimension    = errors.New("dimensions cannot be negative or 0")
	ErrInvalidWinCondition = errors.New("number of boxes to fill to win cannot be negative or 0")
)

// BoxContent is the state of a tic tac toe box
type BoxContent int

// States of a tic tac toe box
const (
	E BoxContent = iota // Empty box
	X                   // Box with a x symbol
	O                   // Box with a o symbol
)

// Board is a nxn matrix defined by user input
type Board struct {
	WinCount           int
	Boxes              [][]BoxContent
	WinConditionChecks []winConditionCheck
}

// direction determines how to traverse in the tic tac toe board when checking if a player has won
type direction int

const (
	right direction = 1  // go rightwards in a row
	down  direction = 1  // go downwards in a column
	left  direction = -1 // go leftwards in a row
	up    direction = -1 // go upwards in a column
	stay  direction = 0  // stay put in a row/column
)

// check defines a particular check path when checking if a player has won
type check struct {
	rowDirection direction
	colDirection direction
}

// winConditionCheck defines a specific condition to determine if a player has won
// for example, checking the row involves checking both right and left direction from a specific box position
// this should always contain a two element slice check
type winConditionCheck struct {
	checks [2]check
}

// NewBoardParams defines the structure for the parameters needed to create a new board
type NewBoardParams struct {
	WinCount   int
	Dimensions int
}

// CheckForWinnerParams defines the structure for the parameters needed to check for a winner
type CheckForWinnerParams struct {
	PlayerSymbol BoxContent
	RowIdx       int
	ColIdx       int
}

// GetBoxContentParams defines the structure for the parameters needed to get a box's content on a particular row and col idx
type GetBoxContentParams struct {
	RowIdx int
	ColIdx int
}

// InsertBoxWithContentParams defines the structure for the parameters needed to insert a player's symbol into a box on a particular row and col idx
type InsertBoxWithContentParams struct {
	RowIdx  int
	ColIdx  int
	Content BoxContent
}

// NewBoard creates a new tic tac toe board
func NewBoard(p NewBoardParams) (*Board, error) {

	if p.WinCount <= 0 {
		return nil, ErrInvalidWinCondition
	} else if p.Dimensions <= 0 {
		return nil, ErrInvalidDimension
	}

	winConditionChecks := generateChecks()

	b := &Board{
		p.WinCount,
		make([][]BoxContent, p.Dimensions),
		winConditionChecks,
	}

	// fill box with empty content
	for row := 0; row < p.Dimensions; row++ {
		b.Boxes[row] = make([]BoxContent, p.Dimensions)
		for col := 0; col < p.Dimensions; col++ {
			b.Boxes[row][col] = E
		}
	}

	return b, nil
}

// SelectBox inserts a player's symbol into a box on a particular row and col idx
func (b Board) SelectBox(p InsertBoxWithContentParams) error {
	if b.Boxes[p.RowIdx][p.ColIdx] != E {
		return ErrBoxOccupied
	}

	b.Boxes[p.RowIdx][p.ColIdx] = p.Content

	return nil
}

// CheckForWinner checks for all possible win conditions from a player's position in a box of a specific row and col index
func (b Board) CheckForWinner(p CheckForWinnerParams) bool {

	for _, winConditionCheck := range b.WinConditionChecks {
		consecutivePlayerSymbolFound := 1

		for _, check := range winConditionCheck.checks {
			for i := 1; i < b.WinCount; i++ {

				checkRowIdx := p.RowIdx + (int(check.rowDirection) * i)
				checkColIdx := p.ColIdx + (int(check.colDirection) * i)

				// skip any checks that has exceeded the tic tac toe board's dimensions
				if checkRowIdx >= len(b.Boxes) || checkRowIdx < 0 || checkColIdx >= len(b.Boxes[0]) || checkColIdx < 0 {
					break
				}

				getBoxContentParams := GetBoxContentParams{
					RowIdx: checkRowIdx,
					ColIdx: checkColIdx,
				}

				if b.getBoxContent(getBoxContentParams) != p.PlayerSymbol {
					break
				}

				consecutivePlayerSymbolFound++
				if consecutivePlayerSymbolFound == b.WinCount {
					return true
				}
			}

		}

	}

	return false
}

// getBoxContent returns the symbol contained within a box of a particular row and col index
func (b Board) getBoxContent(p GetBoxContentParams) BoxContent {
	return b.Boxes[p.RowIdx][p.ColIdx]
}

// generateChecks returns all possible win conditions to check for to determine a player's victory
// this involves checking horixontally, vertically, diagonally and reverse diagonally
func generateChecks() []winConditionCheck {

	winViaRowChecks := winConditionCheck{
		checks: [2]check{
			check{stay, right},
			check{stay, left},
		},
	}
	winViaColChecks := winConditionCheck{
		checks: [2]check{
			check{down, stay},
			check{up, stay},
		},
	}
	winViaDiagonalChecks := winConditionCheck{
		checks: [2]check{
			check{down, right},
			check{up, left},
		},
	}
	winViaReverseDiagonalChecks := winConditionCheck{
		checks: [2]check{
			check{down, left},
			check{up, right},
		},
	}

	return []winConditionCheck{
		winViaRowChecks,
		winViaColChecks,
		winViaDiagonalChecks,
		winViaReverseDiagonalChecks,
	}
}
