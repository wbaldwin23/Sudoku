package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("William's Sudoku Solver!")
	var filename string
	fmt.Println("Enter file name: ")
	fmt.Scanln(&filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	fmt.Println("Reading file", filename)
	file.Close()

	var counter int
	for _, eachline := range txtlines {
		fmt.Println(eachline)
		counter++
	}

}
