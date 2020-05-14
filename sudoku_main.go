package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Constants
const boardSize = 171

//Validating the input file.
//This function uses multiple tests to check if
//the provided file contents represent a
//solvable Sudoku board.
func validator(input string) bool {

	//Variables
	rowCount := 0
	testPassed := 0
	isValid := false

	//Test 1
	//Checks if the file contains the required number
	//of characters needed for a 9x9 board. This is
	//done by finding the size of the file contents.
	b, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	if len(str) != boardSize {
		fmt.Println("Error: The file contents are not valid. Please use a file with" +
			"81 integers. (Blanks should be represented with zeros)")
		testPassed++
	}

	//Test 2
	//Checks if the board is organized in a 9x9
	//configuration. This is done by counting
	//the number of lines that contain the file's
	//contents.
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowCount++
	}

	if rowCount == 9 {
		testPassed++
	} else {
		fmt.Println("Error: The file contents are not valid. Please organize file contents" +
			"into a 9x9 board with spaces. (Blanks should be represented with zeros)")
	}

	if testPassed == 2 {
		isValid = true
	}

	return isValid

}

func main() {

	//Variables
	var fileName string

	//Initial Output
	//Prompts user for file name.
	fmt.Printf("William's Sudoku Solver\n")
	fmt.Println("Enter name of Sudoku file: ")
	fmt.Scanln(&fileName)
	fmt.Println("Reading", fileName)

	//Reading File
	//Reads file with given name &

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	board := parseInput(string(data))
	printBoard(board)
	if backtrack(&board) {
		fmt.Println("The Sudoku was solved successfully:")
		printBoard(board)
	} else {
		fmt.Printf("The Sudoku can't be solved.")
	}

	//Validating File
	//The validator function is used to
	//check if the file contents contain
	//a solvable Sudoku board.
	fmt.Println("\nChecking if file contains a valid board...")

	if validator(fileName) == true {
		fmt.Println("The file contains a valid board.")
	}

}
func parseInput(fileName string) [9][9]int {
	board := [9][9]int{}
	scanner := bufio.NewScanner(strings.NewReader(fileName))

	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			i1, _ := strconv.Atoi(scanner.Text())
			board[row][col] = i1
		}
	}
	return board
}
func printBoard(board [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%d ", board[row][col])
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("")
		} else {
			fmt.Println()
		}

	}

}
func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}
func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}
func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}
func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
