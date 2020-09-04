// Package terminal defines the command line interactions for the game tic tac toe
package terminal

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/dev-amos/tictactoe/board"
	"github.com/dev-amos/tictactoe/view"
)

type Terminal struct {
	InputReader *bufio.Reader
}

// PrintBoard prints out the tic tac toe's board on command line
func (t Terminal) PrintBoard(b board.Board) {

	var sb strings.Builder
	var boxPosition int = 1

	maxNumberOfBoxes := len(b.Boxes) * len(b.Boxes)
	paddingSizeForEachDigitOnBox := digitsCount(maxNumberOfBoxes)
	dashesPerBox := 4 + (paddingSizeForEachDigitOnBox - 1)

	for row := range b.Boxes {
		var horizontalDashLines strings.Builder

		for col := range b.Boxes[row] {
			sb.WriteString(" ")

			boxContentStr := convertBoxContent(b.Boxes[row][col])

			if boxContentStr != "" {
				sb.WriteString(fmt.Sprintf("%*s", paddingSizeForEachDigitOnBox, boxContentStr))
			} else {
				sb.WriteString(fmt.Sprintf("%*s", paddingSizeForEachDigitOnBox, strconv.Itoa(boxPosition)))
			}

			if col < len(b.Boxes[row])-1 {
				sb.WriteString(" |")
			}

			horizontalDashLines.WriteString(strings.Repeat("-", dashesPerBox))
			boxPosition++
		}

		sb.WriteString("\n")

		if row != len(b.Boxes)-1 {
			fmt.Fprintln(&sb, horizontalDashLines.String())
		}

	}

	fmt.Println(sb.String())
}

// GetUserName gets the nth player name from command line and returns it
func (t Terminal) GetUserName(playerCount int) (string, error) {
	fmt.Printf("Enter name for Player %d\n", playerCount)
	name, err := t.InputReader.ReadString('\n')
	name = strings.TrimSpace(name)

	if err != nil {
		return "", err
	}

	return name, nil
}

// GetDimensions gets the dimension size for the tic tac toe board form command line and returns it
func (t Terminal) GetDimensions() (int, error) {
	fmt.Println("Enter game dimensions for tictactoe:")

	input, err := t.InputReader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	dimension, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return dimension, nil
}

// GetUserToSelectBox gets user to choose a numbered position on the tic tac toe board from the command line to select their move
func (t Terminal) GetUserToSelectBox(p view.GetUserToSelectBoxParams) (int, error) {
	fmt.Printf("%s, choose a box to place an '%s' into:\n", p.PlayerName, convertBoxContent(p.PlayerSymbol))

	input, err := t.InputReader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	idxChoice, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return idxChoice, nil
}

// DeclareWinner prints out victory message on the command line for the player that has won
func (t Terminal) DeclareWinner(playerName string) {
	fmt.Printf("Congratulations %s! You have won.\n", playerName)
}

// DeclareDraw prints out a message on the command line indicating that the game has ended with a draw
func (t Terminal) DeclareDraw() {
	fmt.Println("This game has ended in a draw!")
}

// convertBoxContent converts the box content constant to its string representation on command line
func convertBoxContent(b board.BoxContent) string {
	switch b {
	case board.X:
		return "x"
	case board.O:
		return "o"
	default:
		return ""
	}
}

// digitsCount returns the number of digits of a number
func digitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}

	return count
}
