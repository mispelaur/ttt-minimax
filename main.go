package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var computer = 1 // rendered as "O", this is our maximizing player
var user = -1    // rendered as "X", this is our minimizing player

func isWinner(board []int, player int) bool {
	winningCombinations := [][]int{
		[]int{2, 4, 6},
		[]int{0, 4, 8},
		[]int{2, 5, 8},
		[]int{1, 4, 7},
		[]int{0, 3, 6},
		[]int{6, 7, 8},
		[]int{3, 4, 5},
		[]int{0, 1, 2},
	}
	for _, wc := range winningCombinations {
		if (board[wc[0]]+board[wc[1]]+board[wc[2]])*player == 3 {
			return true
		}
	}
	return false
}

func getMinimaxValue(board []int, player int) (score int) {
	emptySpaces := findEmpties(board)
	if isWinner(board, user) {
		return -10
	}
	if isWinner(board, computer) {
		return 10
	}
	if len(emptySpaces) == 0 {
		return 0
	}

	// score for initial comparison is a really big if it's the user's turn
	// ... but really small if it's the computer's turn
	score = 99999999999 * player * -1
	for _, empty := range emptySpaces {
		newBoard := append([]int{}, board...)
		newBoard[empty] = player
		newBoardScore := getMinimaxValue(newBoard, player*-1)
		if player == user { // minimizing
			if newBoardScore < score {
				score = newBoardScore
			}
		}
		if player == computer { // maximizing
			if newBoardScore > score {
				score = newBoardScore
			}
		}
	}

	return score
}

func findEmpties(board []int) (empties []int) {
	for i, val := range board {
		if val == 0 {
			empties = append(empties, i)
		}
	}
	return empties
}

func chooseComputerMove(originalBoard []int) (bestMove int) {
	emptySpaces := findEmpties(originalBoard)
	bestScore := -99999999999 // computer player maximizes, so this is negative
	bestMove = -1
	for _, potentialMove := range emptySpaces {
		newBoard := append([]int{}, originalBoard...)
		newBoard[potentialMove] = computer
		newBoardScore := getMinimaxValue(newBoard, user)
		if newBoardScore > bestScore { // maximizing
			bestScore = newBoardScore
			bestMove = potentialMove
		}
	}
	return bestMove
}

func printBoard(board []int, overwrite bool) {
	transformed := []string{}
	for _, cell := range board {
		val := ""
		if cell == 0 {
			val = "-"
		}
		if cell == 1 {
			val = "O"
		}
		if cell == -1 {
			val = "X"
		}
		transformed = append(transformed, val)
	}
	if overwrite {
		fmt.Print("\033[6A")
	}
	fmt.Printf("\n   %v | %v | %v\n", transformed[6], transformed[7], transformed[8])
	fmt.Printf("  -----------\n")
	fmt.Printf("   %v | %v | %v\n", transformed[3], transformed[4], transformed[5])
	fmt.Printf("  -----------\n")
	fmt.Printf("   %v | %v | %v\n", transformed[0], transformed[1], transformed[2])
	return
}

func promptForMove(overwrite bool) int {
	if overwrite {
		fmt.Print("\033[7A")
		fmt.Print("\033[0K")
		fmt.Print("Chose your next move: ")
	} else {
		fmt.Print("Your move: ")
	}
	reader := bufio.NewReader(os.Stdin)
	move, err := reader.ReadString('\n')
	if err != nil {
		return -1
	}
	if overwrite {
		fmt.Print("\033[6B")
	}
	move = strings.TrimSuffix(move, "\n")
	i, err := strconv.Atoi(move)
	if err != nil {
		return -1
	}
	return i
}

func displayIntro() {
	intro := `Play "X" against me playing "O". The first
move is yours. Choose your cell by number:

   6 | 7 | 8
  -----------
   3 | 4 | 5
  -----------
   0 | 1 | 2 

`
	fmt.Print(intro)
}

func main() {
	// Play the game
	board := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	displayIntro()

	userMove := promptForMove(false)
	board[userMove] = user
	printBoard(board, false)
	computerMove := chooseComputerMove(board)
	board[computerMove] = computer
	// artificial delay to make the game output easier to follow
	time.Sleep(time.Second * 1)
	printBoard(board, true)

	for {
		userMove := promptForMove(true)
		board[userMove] = user
		printBoard(board, true)
		if isWinner(board, user) {
			fmt.Println("\nYOU WIN!")
			os.Exit(1)
		}
		if len(findEmpties(board)) == 0 {
			fmt.Println("\nWE TIE!")
			os.Exit(1)
		}
		printBoard(board, true)
		computerMove = chooseComputerMove(board)
		board[computerMove] = computer
		time.Sleep(time.Second * 1)
		printBoard(board, true)
		if isWinner(board, computer) {
			fmt.Println("\nI WIN!")
			os.Exit(1)
		}
		if len(findEmpties(board)) == 0 {
			fmt.Println("\nWE TIE!")
			os.Exit(1)
		}
	}
}
