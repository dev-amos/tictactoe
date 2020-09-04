// package main defines the starting point to start the tic tac toe game
package main

import (
	"bufio"
	"log"
	"os"

	"github.com/dev-amos/tictactoe/board"
	"github.com/dev-amos/tictactoe/player"
	"github.com/dev-amos/tictactoe/player/real"
	"github.com/dev-amos/tictactoe/view"
	"github.com/dev-amos/tictactoe/view/terminal"
)

//TODO: handle packaging the code for running correctly
func main() {

	view := terminal.Terminal{
		InputReader: bufio.NewReader(os.Stdin),
	}

	players, err := createPlayers(view)
	if err != nil {
		log.Fatalf("create players failed, err=%v", err)
	}

	dimension, err := view.GetDimensions()
	if err != nil {
		log.Fatalf("get dimensions from user input failed, err=%v", err)
	}

	newBoardParams := board.NewBoardParams{
		WinCount:   3,
		Dimensions: dimension,
	}

	b, err := board.NewBoard(newBoardParams)
	if err != nil {
		log.Fatalf("create board failed, err=%v", err)
	}

	startGame(players, b, view)
}

// createPlayers takes 2 player names as input from command line and creates 2 player models
func createPlayers(v view.View) ([]player.Player, error) {
	firstPlayerName, err := v.GetUserName(1)
	if err != nil {
		return nil, err
	}

	newPlayerParams := real.NewPlayerParams{
		Name:   firstPlayerName,
		Symbol: board.X,
	}
	player1 := real.NewPlayer(newPlayerParams)

	secondPlayerName, err := v.GetUserName(2)
	if err != nil {
		return nil, err
	}

	newPlayerParams = real.NewPlayerParams{
		Name:   secondPlayerName,
		Symbol: board.O,
	}
	player2 := real.NewPlayer(newPlayerParams)

	return []player.Player{player1, player2}, nil
}

// startGame will get the players to choose their move on the tic tac toe board and constantly checks for win condition at every move
func startGame(players []player.Player, b *board.Board, v view.View) {

	dimension := len(b.Boxes)
	availableMoves := dimension * dimension
	playerIdx := 0

	// game ends when all possible moves have been made leading to a draw or when a player has won
	for availableMoves > 0 {
		player := players[playerIdx]

		v.PrintBoard(*b)

		getUserToSelectBoxParams := view.GetUserToSelectBoxParams{
			PlayerName:   player.GetName(),
			PlayerSymbol: player.GetSymbol(),
		}

		// get player selection on the box position to place their symbol
		idxChoice, err := v.GetUserToSelectBox(getUserToSelectBoxParams)
		if err != nil {
			log.Fatalf("get user selection failed, err=%v", err)
		}

		// get selected box's row and col index
		selectedBoardRowIdx := (idxChoice - 1) / dimension
		selectedBoardColIdx := (idxChoice - 1) % dimension

		insertBoxWithContentParams := board.InsertBoxWithContentParams{
			RowIdx:  selectedBoardRowIdx,
			ColIdx:  selectedBoardColIdx,
			Content: player.GetSymbol(),
		}

		// populate board with the choice
		b.SelectBox(insertBoxWithContentParams)

		checkForWinnerParams := board.CheckForWinnerParams{
			PlayerSymbol: player.GetSymbol(),
			RowIdx:       selectedBoardRowIdx,
			ColIdx:       selectedBoardColIdx,
		}

		// check if player's move has made him/her the winner
		if b.CheckForWinner(checkForWinnerParams) {
			v.PrintBoard(*b)
			v.DeclareWinner(player.GetName())
			return
		}

		// switch to next player
		playerIdx = (playerIdx + 1) % len(players)
		availableMoves--
	}

	v.DeclareDraw()
}
