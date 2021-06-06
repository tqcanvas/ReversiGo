# ReversiGo
 Reversi game built in Go with MCTS implementation

**Instructions**
My Reversi game project is built in Go, utilizing goroutines to speed up the execution of random
playouts in Pure MCTS implementation.

To run the program, there are two separate sets of instructions based on the operating system
due to the ANSI escape codes used in the code to draw a colorful Reversi board onto the
Command Prompt or Terminal. Please ensure your computer has an updated version of Go
before running the program.

For Windows OS:
* Open Command Prompt
* Run &quot;go get golang.org/x/sys/windows&quot;
* Navigate to /ReversiMCTS/main folder
* Run &quot;go run . &quot;

For Unix based OS:
* Delete &quot;init\_windows.go&quot; from the ReversiMCTS/main folder
* Open Terminal
* Navigate to /ReversiMCTS/main folder
* Run &quot;go run \*.go&quot;

**Implementation Breakdown**
I have broken up my Go code into 5 .go files for easier access to the implementation details of
my Reversi program. These files are called: game.go, logic.go, MCTS.go, heuristic.go, and
main.go. I will write down in detail the functions and methods implemented in each file below.

_game.go:_
This file details the Game struct that is designed to hold the Reversi game board state and
information on the game state like the score of black and white pieces, the player&#39;s turn (-1 for
black and 1 for white), and the valid moves on the board for the specified player.

This file also implements a Print method for the Game struct that uses ANSI escape codes to
print a colored Reversi board with cells filled in for the given board state. Further, I have
implemented a newGame() function and a createGame() function in this file. The newGame
function sets up a starting board state and passes it into a Game struct to be returned. This
function will be used to initialize every game in the program. The createGame function takes an
8x8 integer array board state and a player turn parameter to create a Game struct with those
details. This is very useful for my AI agents to copy the original Reversi game state and run
random playouts or run tree traversals on children states from possible moves.

_logic.go:_
This file contains the bulk of the Reversi game logic implemented to have a running error-free
Reversi game. It contains a definition of the Cell struct that is a coordinate for the [8][8]int board
state. Further, functions are defined as helper functions to see if a Cell is valid or if a Cell is
present in an array slice of Cells.

There are Game methods defined to findAdjacents and findValidMoves by iterating through
placed cells on a board state and their cardinal directions to generate an array of valid Cells that
can be used as a move by a player. To compliment this, there is a playMove function that takes
in a valid Cell move and advances the board state in a game struct, updating player turn,
flipping cells as necessary, and recalculating the next player&#39;s valid moves.

Finally, there are helper functions defined to advance a Game struct&#39;s turn, update score fields
in the Game struct, determine a winner for a Game, and check if a Game is finished (neither
player has any valid moves). The isDone() function to check if a Game is finished is especially
useful in running AI agents and determining the end point for random playouts or DFS search.

_MCTS.go:_
This file contains a simulatePlayout function and a playouts function to represent a Pure MCTS
implementation.
The simulatePlayout function copies a given game state and plays random moves until the
game is finished. It then passes the first random move played and the ending outcome of the
game into a channel.
The playouts function runs concurrent goroutines of the simulatePlayout function for a specified
n times. Every result is passed into a channel and saved into arrays which is then parsed for the
best starting random move. The best move is returned from the function as a Cell struct.

_heuristics.go:_
This file contains the functions used in a MiniMax AI agent using Alpha Beta Pruning and
custom heuristics to evaluate given board states.

There is an evaluate function that takes a given board state and player turn and returns a
float32 score for the position. This evaluate function is used in the AlphaBeta search that uses
DFS to traverse down board states that generate children states from possible moves. The
results from AlphaBeta search are passed into abMove() which returns a Cell for move chosen
by agent.

_main.go:_
This file contains the main function which is a text based interface for interacting with the
ReversiGo program. It enables the user to select players for the Reversi game (ranging from
human input, to the pure MCTS agent, to the Alpha Beta search agent). There can only be one
human player in the current implementation. There is also a testAI implementation that can run
a specified number of games with a given playouts number for MCTS and a given max depth for
the Alpha Beta search.