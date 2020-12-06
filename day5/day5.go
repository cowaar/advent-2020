package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day5/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	seatArray := getSeatArray(fileArray)
	fmt.Println(seatArray)
	listUnfilledSeats(seatArray)
}
func listUnfilledSeats(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++{
			if arr[i][j] ==0{
				fmt.Println(i,j)
			}
		}
	}
}

func getSeatArray(arr []string) [][]int {
	result := make([][]int, 127)
	for i := range result {
		result[i] = make([]int, 8)
	}
	for _, s := range arr {
		row, col := getLoc(s)
		result[row][col] = 1
	}
	return result
}

func findHighestId(arr []string) int {
	result := 0
	for _, s := range arr {
		row, col := getLoc(s)
		id := row*8 + col
		if id > result {
			result = id
		}
	}
	return result
}
func getLoc(s string) (int, int) {
	rowStr := s[:7]
	colStr := s[7:]
	return getLoc1D(rowStr), getLoc1D(colStr)
}
func getLoc1D(s string) int {
	length := len(s)
	lower, upper := 0, int(math.Exp2(float64(length)))
	for i := 0; i < length; i++ {
		var up = -1
		char := s[i : i+1]
		if char == "F" || char == "L" {
			up = 0
		} else if char == "B" || char == "R" {
			up = 1
		}
		lower, upper = binSearch(lower, upper, up)
	}
	return lower
}
func binSearch(lower, upper int, up int) (int, int) {
	dif := upper - lower
	halfDif := dif / 2
	if up == 1 {
		return lower + halfDif, upper
	} else if up == 0 {
		return lower, lower + halfDif
	} else {
		fmt.Println("unable to run binSearch")
		return 0, 0
	}

}
