package main

import (
	"fmt"
	"strconv"
	"time"

	//"strconv"
	//"time"
)

//Helper function to convert cell input into move
func cellConv(input string) Cell {
	if len(input) != 2 {
		return Cell{100,100}
	}

	abcs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var col int
	col = 99

	for i := 0; i < len(abcs); i++ {
		if abcs[i] == input[0] {
			col = i
		}
	}

	digits := "0123456789"

	var row int
	row = 99

	for i := 0; i < len(digits); i++ {
		if digits[i] == input[1] {
			row = i
		}
	}
	row--

	return Cell{row, col}
}

//Main function to set up a game of ReversiGo
func main() {
	//Number of playouts in Pure MCTS
	n := 10000
	//Depth traversed by Alpha Beta Pruning of Minimax
	deep := 6

	//Parse player choice and setup Reversi
	var player1 string
	var player2 string

	fmt.Println("Welcome to ReversiGo!")
	fmt.Println("Please pick a player for Black")
	fmt.Println("There can only be one human player")
	fmt.Println("Input 1 for Human Input, 2 for PureMCTS, and 3 for Alpha Beta Pruning Heuristics")
	fmt.Scanln(&player1)
	p1, _ := strconv.Atoi(player1)

	for p1 < 1 || p1 > 3 {
		fmt.Println("Non valid player number input in ReversiGo")
		fmt.Println("Please pick a player for Black")
		fmt.Scanln(&player1)
		p1, _ = strconv.Atoi(player1)
	}

	fmt.Println("Please pick a player for White")
	fmt.Scanln(&player2)
	p2, _ := strconv.Atoi(player2)

	for p2 < 1 || p2 > 3 {
		fmt.Println("Non valid player number input in ReversiGo.")
		fmt.Println("Please pick a player for White")
		fmt.Scanln(&player2)
		p2, _ = strconv.Atoi(player2)
	}

	game := newGame()
	start := time.Now()
	elapsed := time.Since(start)

	//Human input vs AI game case
	if p1 == 1 || p2 == 1 {
		fmt.Println("Have fun versing ReversiGo! Please input moves in form 'A1'")
		game.Print()
		fmt.Println()

		var move string
		aiturn := -1

		if p1 == 1 {
			fmt.Println("Input your move: ")
			fmt.Scanln(&move)

			humanMove := cellConv(move)

			for !humanMove.valid() || !Find(game.validMoves, humanMove) {
				fmt.Println("Your move is not valid. Please input a new move: ")
				fmt.Scanln(&move)

				humanMove = cellConv(move)
			}
			game.playMove(humanMove)
			game.Print()
			fmt.Println()
			aiturn = 1
		}

		for !game.isDone(){
			if len(game.validMoves) == 0 {
				game.turn = enemy(game.turn)
				game.findAdjacents()
				game.findValidMoves()
			}

			switch {
			case game.turn == aiturn:
				switch {
				case p1 == 2 || p2 == 2:
					start = time.Now()
					best := playouts(game, n)
					fmt.Printf("AI chose move: %v\n", best)
					elapsed = time.Since(start)
					fmt.Printf("MCTS %v random playouts took %s\n", n, elapsed)
					game.playMove(best)
					game.Print()
					fmt.Println()
				case p1 == 3 || p2 == 3:
					start = time.Now()
					cell := abMove(game.turn, game.board, deep)
					fmt.Printf("AI chose move: %v\n", cell)
					elapsed = time.Since(start)
					fmt.Printf("AB Search took %s\n", elapsed)
					game.playMove(cell)
					game.Print()
					fmt.Println()
				}

			case game.turn != aiturn:
				fmt.Println("Input your move: ")
				fmt.Scanln(&move)
				humanMove := cellConv(move)

				for !humanMove.valid() || !Find(game.validMoves, humanMove) {
					fmt.Println("Your move is not valid. Please input a new move: ")
					fmt.Scanln(&move)

					humanMove = cellConv(move)
				}
				game.playMove(humanMove)
				game.Print()
				fmt.Println()
			}
		}
	} else { //2 AI agents
		fmt.Println("ReversiGo Simulation between 2 AI Agents")
		game.Print()
		fmt.Println()

		for !game.isDone() {
			if len(game.validMoves) == 0 {
				game.turn = enemy(game.turn)
				game.findAdjacents()
				game.findValidMoves()
			}
			switch {
			case game.turn == -1:
				blackAgent := p1
				switch {
				case blackAgent == 2:
					start = time.Now()
					best := playouts(game, n)
					fmt.Printf("AI chose move: %v\n", best)
					elapsed = time.Since(start)
					fmt.Printf("MCTS %v random playouts took %s\n", n, elapsed)
					game.playMove(best)
					game.Print()
					fmt.Println()
				case blackAgent == 3:
					start = time.Now()
					cell := abMove(game.turn, game.board, deep)
					fmt.Printf("AI chose move: %v\n", cell)
					elapsed = time.Since(start)
					fmt.Printf("AB Search took %s\n", elapsed)
					game.playMove(cell)
					game.Print()
					fmt.Println()
				}
			case game.turn == 1:
				whiteAgent := p2
				switch {
				case whiteAgent == 2:
					start = time.Now()
					best := playouts(game, n)
					fmt.Printf("AI chose move: %v\n", best)
					elapsed = time.Since(start)
					fmt.Printf("MCTS %v random playouts took %s\n", n, elapsed)
					game.playMove(best)
					game.Print()
					fmt.Println()
				case whiteAgent == 3:
					start = time.Now()
					cell := abMove(game.turn, game.board, deep)
					fmt.Printf("AI chose move: %v\n", cell)
					elapsed = time.Since(start)
					fmt.Printf("AB Search took %s\n", elapsed)
					game.playMove(cell)
					game.Print()
					fmt.Println()
				}
			}
		}
	}
	//Print Outcomes
	game.calcScore()
	fmt.Printf("Black scored %v tiles\n", game.BlackCount)
	fmt.Printf("White scored %v tiles\n", game.WhiteCount)
	if game.BlackCount > game.WhiteCount {
		fmt.Printf("The winner is Black!\n")
	} else if game.BlackCount < game.WhiteCount {
		fmt.Printf("The winner is White!\n")
	} else{
		fmt.Printf("It was a draw!\n")
	}

	//testAI(5, 1000, 25)
	//testAI(6, 1000, 25)
	//testAI(7, 1000, 25)
	//testAI(5, 5000, 25)
	//testAI(6, 5000, 25)
	//testAI(7, 5000, 25)
	//testAI(5, 10000, 25)
	//testAI(6, 10000, 25)
	//testAI(7, 10000, 25)
}

//Function call to simulate many games between AI agents
func testAI(depth int, playCount int, tests int) {
	//Number of tests
	n := tests
	outcomes := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Simulation game: %v\n", i)
		game := newGame()

		mctsColor := 0
		abColor := 0

		if i%2 == 0 {
			mctsColor = -1
			abColor = 1
		} else {
			mctsColor = 1
			abColor = -1
		}

		for !game.isDone() {
			if len(game.validMoves) == 0 {
				game.turn = enemy(game.turn)
				game.findAdjacents()
				game.findValidMoves()
			}
			switch {
			//Can swap cases to test different starting colors for each AI agent
			case game.turn == mctsColor:
				best := playouts(game, playCount)
				game.playMove(best)
			case game.turn == abColor:
				cell := abMove(game.turn, game.board, depth)
				game.playMove(cell)
			}
		}
		game.calcScore()
		if game.BlackCount > game.WhiteCount {
			outcomes[i] = -1
			fmt.Printf("The winner is Black!\n")
		} else if game.BlackCount < game.WhiteCount {
			outcomes[i] = 1
			fmt.Printf("The winner is White!\n")
		} else {
			outcomes[i] = 0
			fmt.Printf("It was a draw!\n")
		}
	}

	//Count outcomes
	mctsWin := 0
	abWin := 0
	draws := 0

	for i := 0; i < len(outcomes); i++{
		if i % 2 == 0 {
			switch {
			case outcomes[i] == -1:
				mctsWin++
			case outcomes[i] == 1:
				abWin++
			default:
				draws++
			}
		} else {
			switch {
			case outcomes[i] == 1:
				mctsWin++
			case outcomes[i] == -1:
				abWin++
			default:
				draws++
			}
		}
	}

	fmt.Println()
	fmt.Printf("MCTS won %v/%v times\n", mctsWin, tests)
	fmt.Printf("AB won %v/%v times\n", abWin, tests)
	fmt.Printf("Draws occured %v/%v times\n", draws, tests)
	fmt.Println()
}
