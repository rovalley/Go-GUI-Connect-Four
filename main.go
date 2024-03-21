/*
NAME
main

DESCRIPTION
This module provides a class to play a game of connect four

Created on March 19, 2024

@author: Ryan O'Valley
*/

package main

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
	"image/color"
	"math/rand"
)

// global variables
var redScore int
var blueScore int
var status string
var playerTurn string
var gameOver bool
var board [6][7]string
var emptyToken = "-"

// function to create a new board
func createNewBoard() {
	// loop through rows and columns and assign to empty token
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			board[r][c] = emptyToken
		}
	}
}

// function for AI turn
func AITurn() {
	// while true
	for true {
		// set column value to random int from 0 to 7
		c := rand.Intn(7)
		// if board at the row and column is equal to empty token
		if board[0][c] == emptyToken {
			// drop token at that column
			dropToken(c)
			break
		}
	}
}

// function to drop token
func dropToken(c int) {
	// if it is not game over
	if gameOver == false {
		// set status to a empty string
		status = ""
		// if the row and column is not empty token
		if board[0][c] != emptyToken {
			// update status
			status = "This column is full"
		} else {
			// loop through rows backwards
			for r := 5; r >= 0; r-- {
				// if row and column is empty token
				if board[r][c] == emptyToken {
					// set player turn
					board[r][c] = playerTurn
					// if there is no winner
					if !checkWinner(playerTurn) {
						// if board is full
						if isBoardFull() {
							// set game over to true
							gameOver = true
							// update status
							status = "Tie Game"
						} else {
							// if player turn is equal to B
							if playerTurn == "B" {
								// set player turn equal to R
								playerTurn = "R"
								// update status
								status = "Red's Turn"
							} else {
								// set player turn equal to B
								playerTurn = "B"
								// update status
								status = "Blue's Turn"
							}
						}
					}
					break
				}
			}
		}
	} else {
		// update status
		status = "Click play again"
	}
}

// function to check horizontal connect four
func checkHorFour(color string) bool {
	// loop through rows and columns
	for r := 0; r <= 5; r++ {
		for c := 0; c <= 3; c++ {
			// check horizontal and see if there is four in a row
			if board[r][c] == color && board[r][c+1] == color && board[r][c+2] == color && board[r][c+3] == color {
				return true
			}
		}
	}
	return false
}

// function to check vertical connect four
func checkVertFour(color string) bool {
	// loop through rows and columns
	for r := 0; r <= 2; r++ {
		for c := 0; c <= 6; c++ {
			// check vertical and see if there is four in a row
			if board[r][c] == color && board[r+1][c] == color && board[r+2][c] == color && board[r+3][c] == color {
				return true
			}
		}
	}
	return false
}

// function to check diagonal connect four
func checkDiagFour(color string) bool {
	// loop through rows and columns
	for r := 0; r <= 2; r++ {
		for c := 0; c <= 3; c++ {
			// check left to right diagonal and see if there is four in a row
			if board[r][c] == color && board[r+1][c+1] == color && board[r+2][c+2] == color && board[r+3][c+3] == color {
				return true
			}
		}
	}
	// loop through rows and columns
	for r := 0; r <= 2; r++ {
		for c := 3; c <= 6; c++ {
			// check right to left diagonal and see if there is four in a row
			if board[r][c] == color && board[r+1][c-1] == color && board[r+2][c-2] == color && board[r+3][c-3] == color {
				return true
			}
		}
	}
	return false
}

// function to check for a winner
func checkWinner(color string) bool {
	// if horizontal or vertical or diagonal connect four is met
	if checkHorFour(color) || checkVertFour(color) || checkDiagFour(color) {
		// if color equals B
		if color == "B" {
			// increment blue score
			blueScore += 1
			// update status
			status = "Blue has won!"
		} else {
			// increment red score
			redScore += 1
			// update status
			status = "Red has won!"
		}
		// set game over to true
		gameOver = true
		return true
	}
	return false
}

// function to play again
func playAgain() {
	// call create new board function
	createNewBoard()
	// set game over to false
	gameOver = false
	// if player turn equals B
	if playerTurn == "B" {
		// update status
		status = "Blue's Turn"
	} else {
		// call AI turn function
		AITurn()
	}
}

// function to check if the game board is full
func isBoardFull() bool {
	// loop through columns
	for c := 0; c <= 6; c++ {
		// if top row at the columns is equal to empty token return false
		if board[0][c] == emptyToken {
			return false
		}
	}
	return true
}

// function to create the window functionality and layout
func createLayout(window *nucular.Window) {
	window.Row(50).Dynamic(1)
	if window.ButtonText("New Game") {
		playAgain()
	}
	window.Row(50).Dynamic(1)
	window.Label(status, "C")
	window.Row(50).Dynamic(3)
	window.Label(fmt.Sprintf("Blue: %d", blueScore), "C")
	if window.ButtonText("AI's Turn") {
		if playerTurn == "R" {
			AITurn()
		}
	}
	window.Label(fmt.Sprintf("Red: %d", redScore), "C")
	window.Row(50).Dynamic(7)
	for i := 0; i < 7; i++ {
		if window.ButtonText("Drop Token") {
			if playerTurn == "B" {
				dropToken(i)
			}
		}
	}
	for r := 0; r < 6; r++ {
		window.Row(30).Dynamic(7)
		for c := 0; c < 7; c++ {
			if board[r][c] == "B" {
				window.LabelColored(board[r][c], "C", color.RGBA{0, 0, 200, 255})
			} else if board[r][c] == "R" {
				window.LabelColored(board[r][c], "C", color.RGBA{200, 0, 0, 255})
			} else {
				window.Label(board[r][c], "C")
			}
		}
	}

}

func main() {
	// assign variables
	status = "Blues's turn"
	playerTurn = "B"
	gameOver = false
	createNewBoard()
	window := nucular.NewMasterWindow(0, "Connect Four", createLayout)
	window.SetStyle(style.FromTheme(style.WhiteTheme, 1.0))
	window.Main()
}
