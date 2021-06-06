package main

//Helper function, return opponent given player
func enemy(player int) int{
	if player == 1 {
		return -1
	} else {
		return 1
	}
}

//Evaluate a board state given player
func evaluate(board [8][8]int, player int) float32 {
	var score float32

	//Created a grid of relative score values for each cell
	//Based on Reversi strategy
	//Corners are stable and valuable, but adjacent cells can lead to losing a corner
	//Edge pieces are somewhat stable and valuable otherwise
	//Cells 2 away from corner facilitate corner capturing
	//All other tiles are weighted at 1

	cellValue := [][]int{
		{50, -20, 10, 5, 5, 10, -20, 50},
		{-20, -20, 1, 1, 1, 1, -20, -20},
		{10, 1, 5, 1, 1, 5, 1, 10},
		{5, 1, 1, 1, 1, 1, 1, 5},
		{5, 1, 1, 1, 1, 1, 1, 5},
		{10, 1, 5, 1, 1, 5, 1, 10},
		{-20, -20, 1, 1, 1, 1, -20, -20},
		{50, -20, 10, 5, 5, 10, -20, 50},
	}

	//Coin Parity
	var pScore float32
	var pCoin float32
	var oCoin float32

	//Score player and opponent cell
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == player {
				pCoin += float32(cellValue[i][j])
			} else if board[i][j] == enemy(player) {
				oCoin += float32(cellValue[i][j])
			}
		}
	}

	//Parity score is difference in scores over total score
	pScore = (pCoin - oCoin) / (pCoin + oCoin)

	//Difference in Mobility
	//Create games assuming board state and turn to calculate size of valid moves
	var mScore float32

	pMob := createGame(board, player)
	oMob := createGame(board, enemy(player))

	playerMobility := len(pMob.validMoves)
	enemyMobility := len(oMob.validMoves)

	switch{
	case playerMobility > enemyMobility:
		mScore = float32(playerMobility)/float32(playerMobility + enemyMobility)
	case enemyMobility > playerMobility:
		mScore = -float32(enemyMobility)/float32(playerMobility + enemyMobility)
	default:
		mScore = 0
	}

	//fmt.Printf("Parity:%v Mobility:%v\n", pScore, mScore)
	//Total evaluated score
	score = (pScore * 500) + (mScore * 600)

	return score
}

//Struct wrapper for returning a maybe Cell
type mCell struct {
	cell Cell
	true bool
}

//Helper function for opponent score, traverses down the tree
func value(player int, board [8][8]int, alpha float32, beta float32, depth int) float32 {
	val, _ := AlphaBeta(enemy(player), board, -beta, -alpha, depth-1)

	return -val
}

//Helper function for final game board value
func final(player int, game Game) float32 {
	game.calcScore()
	var MAX float32
	MAX = 10000

	if player == 1 {
		if game.WhiteCount > game.BlackCount {
			return MAX
		} else if game.BlackCount > game.WhiteCount {
			return -MAX
		} else {
			return 0
		}
	} else {
		if game.WhiteCount > game.BlackCount {
			return -MAX
		} else if game.BlackCount > game.WhiteCount {
			return MAX
		} else {
			return 0
		}
	}
}

//Alpha Beta Search
func AlphaBeta(player int, board [8][8]int, alpha float32, beta float32, depth int) (float32, mCell){
	//End tree traversal
	if depth == 0 {
		return evaluate(board, player), mCell{Cell{0,0}, false}
	}

	//Create a list of all player moves
	temp := createGame(board, player)
	possibleMoves := temp.validMoves

	//No player moves left
	if len(possibleMoves) == 0 {
		//Finished game
		if temp.isDone() {
			return final(player, temp), mCell{Cell{0,0}, false}
		} else { //Evaluate opponent's move
			return value(player, board, alpha, beta, depth), mCell{Cell{0,0}, false}
		}
	}

	//Store best move in best var
	best := mCell{possibleMoves[0], true}

	//Search all children from valid moves
	for i := 0; i < len(possibleMoves); i++{
		//Alpha Beta pruning step
		if alpha >= beta {
			break
		}

		//Play a move to get children state
		state := createGame(board, player)
		state.playMove(possibleMoves[i])

		//Check opponent value
		val := value(player, state.board, alpha, beta, depth)

		//Best score we can achieve comes from best move
		if val > alpha {
			alpha = val
			best = mCell{possibleMoves[i], true}
		}
	}

	return alpha, best
}

//Alpha Beta, returns best move
func abMove(player int, board [8][8]int, depth int) Cell {
	_, someCell := AlphaBeta(player, board, float32(-10000), float32(10000), depth)

	return someCell.cell
}