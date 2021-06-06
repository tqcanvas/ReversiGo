package main

import (
	"math/rand"
	"time"
)

//Struct to hold a move and its playout outcome
type MoveResult struct {
	move   Cell
	winner int
}

//Function to run a single random playout given game state
func simulatePlayout(game Game, result chan<- MoveResult) {
	//Create a board state from passed game
	simulation := createGame(game.board, game.turn)

	//Get a random generator
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	//Store first move
	var firstMove Cell
	i := 0

	//Loop through random game
	for !simulation.isDone(){
		var ranMove Cell

		if len(simulation.validMoves) != 0 {
			//Play a random move
			ranMove = simulation.validMoves[r.Intn(len(simulation.validMoves))]
			simulation.playMove(ranMove)
		} else {
			simulation.nextTurn()
			simulation.findAdjacents()
			simulation.findValidMoves()

			ranMove = simulation.validMoves[r.Intn(len(simulation.validMoves))]
			simulation.playMove(ranMove)
		}

		if i == 0 {
			firstMove = ranMove
		}
		i++
	}

	//Return winner of the random playout
	simulation.findWinner()
	res := MoveResult{firstMove, simulation.outcome}
	result <- res
}

//Run a number of playouts to return a best move cell
func playouts(game Game, n int) Cell {
	switch{
	//Return only move if only one move possible
	case len(game.validMoves) == 1:
		return game.validMoves[0]
	//Run playouts
	default:
		//Channel to pull game playout results
		c := make(chan MoveResult, 100)
		defer close(c)

		//Run number of playouts
		for i := 0; i < n; i++ {
			go simulatePlayout(game, c)
		}

		//Count outcome of playouts for each move
		var wins [64]int
		var draws [64]int
		var losses [64]int
		var best [64]float32

		for i := 0; i < n; i++ {
			res := <- c

			switch {
			case res.winner == game.turn:
				wins[(res.move.x*8)+res.move.y]++
			case res.winner == 0:
				draws[(res.move.x*8)+res.move.y]++
			default:
				losses[(res.move.x*8)+res.move.y]++
			}
		}

		//Calculate best move
		max := float32(-1)
		maxIn := 0
		for i := 0; i < 64; i++ {
			if wins[i] + draws[i] + losses[i] != 0 {
				best[i] = float32(wins[i] + draws[i]) / float32(wins[i] + draws[i] + losses[i])

				if best[i] > max {
					max = best[i]
					maxIn = i
				}
			}
		}
		y := maxIn % 8
		x := (maxIn - y) / 8
		return Cell{x, y}
	}
}
