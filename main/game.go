package main

import (
	"fmt"
	"strconv"
)

//Struct to hold an instance of a Reversi game
type Game struct {
	board      [8][8]int //2D Array to store board state
	WhiteCount int       //White Score
	BlackCount int       //Black Score
	outcome    int       //Winner of the game, 0 if draw
	turn       int       //-1 for Black, 1 for White
	placed     []Cell    //Cells placed
	adjacents  []Cell    //Empty adjacent cells
	validMoves []Cell    //Possible adjacent moves for turn
}

//Game method to print the board
func (game Game) Print() {
	for i := 0; i < 8; i++ {
		if i == 0 {
			fmt.Printf("\u001b[40m\u001b[37m    A  B  C  D  E  F  G  H ")
			fmt.Printf("\u001b[0m")
			fmt.Println()
		}
		fmt.Printf("\u001b[40m\u001b[37m %s ", strconv.Itoa(i+1))
		fmt.Printf("\u001b[0m")

		for j := 0; j < 8; j++ {

			if game.board[i][j] == -1 {
				fmt.Printf("\u001b[42m\u001b[30m %s ", "●")
				fmt.Printf("\u001b[0m")
			} else if game.board[i][j] == 1 {
				fmt.Printf("\u001b[42m\u001b[37m %s ", "●")
				fmt.Printf("\u001b[0m")
			} else {
				fmt.Printf("\u001b[42m\u001b[90m . ")
				fmt.Printf("\u001b[0m")
			}
		}
		fmt.Println()
	}
}

//Function to create a new game
func newGame() Game {
	//Create a starting board state
	board := [8][8]int{}

	board[3][3] = 1
	board[3][4] = -1
	board[4][3] = -1
	board[4][4] = 1

	//Initialize game stats
	w := 2
	b := 2
	outcome := 0
	turn := -1

	//Create a game struct
	game := Game{board, w, b, outcome, turn, []Cell{}, []Cell{}, []Cell{}}

	//Fill placed slice
	game.placed = append(game.placed, Cell{3, 3})
	game.placed = append(game.placed, Cell{3, 4})
	game.placed = append(game.placed, Cell{4, 3})
	game.placed = append(game.placed, Cell{4, 4})

	//Fill other slices
	game.findAdjacents()
	game.findValidMoves()

	return game
}

//Function to create a game given board state
func createGame(state [8][8]int, turn int) Game {
	//Copy board state
	board := state

	//Generate the game
	game := Game{board, 0, 0, 0, turn, []Cell{}, []Cell{}, []Cell{}}

	//Fill placed slice
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] != 0 {
				game.placed = append(game.placed, Cell{i, j})
			}
		}
	}

	//Fill other slices
	game.findAdjacents()
	game.findValidMoves()
	game.calcScore()

	return game
}
