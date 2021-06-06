package main

import "fmt"

//Cell in game board struct
type Cell struct {
	x int
	y int
}

//Cell method to check if it fits in reversi bounds
func (c Cell) valid() bool {
	if c.x < 8 && c.x >= 0 {
		if c.y < 8 && c.y >= 0 {
			return true
		}
		return false
	}
	return false
}

//Helper function to check if cell exists in slice
func Find(slice []Cell, cell Cell) bool {
	for _, x := range slice {
		if x == cell {
			return true
		}
	}
	return false
}

//Variable declaration to find 8 surrounding cells
var Shift = []Cell{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
}

//Find adjacent cells returning a slice given a cell
func (game Game) Adjacent(cell Cell) []Cell {
	var result []Cell

	//Iterate over possible adjacent cells
	for _, move := range Shift {
		freeCell := Cell{cell.x + move.x, cell.y + move.y}

		//If adjacent cell is valid, add to slice
		if freeCell.valid() {
			if game.board[freeCell.x][freeCell.y] == 0 {
				result = append(result, freeCell)
			}
		}
	}
	return result
}

//Game method to find all valid adjacent cells on board
func (game *Game) findAdjacents() {
	placedCells := game.placed
	var adjacents []Cell

	//Iterate over placed cells
	for _, cell := range placedCells {
		foundAdj := game.Adjacent(cell)

		//Find adjacent cells and add to result slice if unique
		for _, x := range foundAdj {
			if !Find(adjacents, x) {
				adjacents = append(adjacents, x)
			}
		}
	}

	game.adjacents = adjacents
}

//Method to find all valid moves on a turn
func (game *Game) findValidMoves() {
	//Game state variables
	adjacents := game.adjacents
	turn := game.turn

	//Slice to fill and return
	var validMoves []Cell

	//Iterate over all free adjacents
	for _, adj := range adjacents {

		//Iterate over all surrounding cells for each adjacent
		for _, move := range Shift {
			maybeCell := Cell{adj.x + move.x, adj.y + move.y}

			//Check if cell is valid and played
			if maybeCell.valid() && game.board[maybeCell.x][maybeCell.y] != 0{
				//If adjacent is same color as turn, change directions
				if game.board[maybeCell.x][maybeCell.y] == turn {
					continue
				}

				//Find a cell of same color in direction
				itrCell := maybeCell
				for itrCell.valid() {
					itrCell.x += move.x
					itrCell.y += move.y

					if !itrCell.valid() {
						break
					}

					//Exit loop if empty cell is reached
					if game.board[itrCell.x][itrCell.y] == 0 {
						break
					}
					//Same color is found
					if game.board[itrCell.x][itrCell.y] == turn {
						if !Find(validMoves, adj) {
							validMoves = append(validMoves, adj)
						}
					}
				}
			}
		}
	}

	game.validMoves = validMoves
}

//Next turn method
func (game *Game) nextTurn() {
	if game.turn == -1 {
		game.turn = 1
	} else {
		game.turn = -1
	}
}

//Method to play a move
//Assumes it is valid
func (game *Game) playMove(cell Cell) {
	//Error check cell
	if !cell.valid(){
		fmt.Printf("Move is out of legal range.\n")
	}

	//Put cell in placed slice
	game.placed = append(game.placed, cell)

	//Change board cell
	turn := game.turn
	game.board[cell.x][cell.y] = turn

	//Iterate over adjacent cells
	for _, move := range Shift {
		nextCell := Cell{cell.x + move.x, cell.y + move.y}

		//Check next cell is valid and opposite color
		if nextCell.valid() && game.board[nextCell.x][nextCell.y] != 0 && game.board[nextCell.x][nextCell.y] != turn {
			flip := false

			//Find the same color
			itrCell := Cell{nextCell.x, nextCell.y}
			for itrCell.valid() {
				itrCell.x += move.x
				itrCell.y += move.y

				if !itrCell.valid() {
					break
				}

				//Exit loop if empty cell is reached
				if game.board[itrCell.x][itrCell.y] == 0 {
					break
				}
				//Same color found
				if game.board[itrCell.x][itrCell.y] == turn {
					flip = true
					break
				}
			}

			//Flip cells
			if flip {
				itrCell := Cell{nextCell.x, nextCell.y}
				for game.board[itrCell.x][itrCell.y] != turn {
					game.board[itrCell.x][itrCell.y] = turn
					itrCell.x += move.x
					itrCell.y += move.y
				}
			}
		}
	}

	//Advance the game state
	game.nextTurn()
	game.calcScore()
	game.findAdjacents()
	game.findValidMoves()
}

//Game method to update score
func (game *Game) calcScore() {
	b := 0
	w := 0

	//Iterate over game board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if game.board[i][j] == -1 {
				b++
			} else if game.board[i][j] == 1 {
				w++
			}
		}
	}

	game.BlackCount = b
	game.WhiteCount = w
}

//Game method to decide winner
func (game *Game) findWinner() {
	if game.BlackCount > game.WhiteCount {
		game.outcome = -1
	} else if game.BlackCount < game.WhiteCount {
		game.outcome = 1
	} else {
		game.outcome = 0
	}
}

//Game method to determine if game is finished
func (game *Game) isDone() bool {
	if len(game.validMoves) == 0 {
		game.nextTurn()
		game.findAdjacents()
		game.findValidMoves()

		if len(game.validMoves) == 0 {
			return true
		} else {
			game.nextTurn()
			game.findAdjacents()
			game.findValidMoves()
			return false
		}
	} else {
		return false
	}
}