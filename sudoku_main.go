package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("William's Sudoku Solver!")
	var filename string
	fmt.Println("Enter file name: ")
	fmt.Scanln(&filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Reading file", filename)

	fmt.Println("Contents of file:")
	fmt.Println(string(data))
}
