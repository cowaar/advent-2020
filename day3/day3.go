package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day3/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	//rowLength := len(fileArray[0])
	var twoDArray = make([][]string, len(fileArray))

	for i, row := range fileArray {
		splitRow := strings.Split(row, "")
		twoDArray[i] = splitRow
		fmt.Println(splitRow)
	}

	a1 := countTrees2(twoDArray, 1, 1)
	a2 := countTrees2(twoDArray, 3, 1)
	a3 := countTrees2(twoDArray, 5, 1)
	a4 := countTrees2(twoDArray, 7, 1)
	a5 := countTrees2(twoDArray, 1, 2)

	fmt.Println(a1, a2, a3, a4, a5)
	fmt.Println(a1 * a2 * a3 * a4 * a5)
	fmt.Println(countTrees1(twoDArray))
}

func countTrees1(input [][]string) int {

	result := 0
	x := 0
	for _, row := range input {
		if row[x] == "#" {
			result++
		}
		x = getInfXCoord(x+3, len(row))
	}

	return result
}

func countTrees2(input [][]string, xInc int, yInc int) int {

	result := 0
	x := 0
	for y, row := range input {
		if y%yInc == 0 {
			if row[x] == "#" {
				result++
			}

			x = getInfXCoord(x+xInc, len(row))
		}
	}
	return result
}

func getInfXCoord(initCoord int, rowLength int) int {
	if initCoord < rowLength {
		return initCoord
	} else {
		return initCoord - rowLength
	}
}
