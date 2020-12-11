package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day11/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	//rowLength := len(fileArray[0])
	var twoDArray = make([][]string, len(fileArray))

	for i, row := range fileArray {
		splitRow := strings.Split(row, "")
		twoDArray[i] = splitRow
	}
	fmt.Println(run(twoDArray,1))
	fmt.Println(run(twoDArray,2))

}
func run(grid [][]string, part int) int{
	inputGrid:=grid
	var outputGrid [][]string
	count:=0
	for {
		 outputGrid= getNewGrid(inputGrid, part)
		//fmt.Println(outputGrid)
		if reflect.DeepEqual(outputGrid,inputGrid){break}
		inputGrid = outputGrid
		count++

	}
	return countSeats(outputGrid)
}
func countSeats(grid [][]string)int{
	takenSeats:= 0
	for _, row := range grid{
		for _, x := range row {
			if x== "#" {
				takenSeats++
			}	}
	}
	return takenSeats
}

func getNewGrid(grid [][]string, part int) [][]string {
	newGrid := make([][]string, len(grid))
	for i, row := range grid {
		var newRow []string
		for j, _ := range row {
			switch part {
			case 1:
				newRow = append(newRow, updateSeat1(grid, j, i))
			case 2:
				newRow = append(newRow, updateSeat2(grid, j, i))
			}
		}
		newGrid[i] = newRow
	}
	return newGrid
}
//--------------------part 2---------------
func updateSeat2(grid [][]string, x,y int )string{
	seat := grid[y][x]
	if seat == "." {
		return "."
	}
	nhbrs := 0
	xStep := []int{0, 0, 1, 1, 1, -1, -1, -1}
	yStep := []int{1, -1, 0, 1, -1, 0, 1, -1}


	for i, _ := range xStep {
		if foundFilledSeats(grid,xStep[i],yStep[i],x,y){
			nhbrs++
		}
	}
	if seat == "L" && nhbrs == 0 {
		return "#"
	} else if seat == "#" && nhbrs >= 5 {
		return "L"
	} else {
		return seat
	}
}

func foundFilledSeats(grid [][]string, xStep,yStep int, x,y int) bool{
	//check a direction
	//for each direction - take steps in that direction until you reach the edge, or you reach a filled seat
	//
	for {
		newX,newY:= x+xStep,y+yStep
		if newX >= len(grid[0]) || newX < 0 || newY >= len(grid) || newY < 0{
			return false

		}else if grid[newY][newX] == "L"{
			return false
		}else if grid[newY][newX] == "#"{

			return true
		}
		x,y = newX,newY
	}
}
// -------------------part 1---------------------

func updateSeat1(grid [][]string, x, y int) string {
	seat := grid[y][x]
	if seat == "." {
		return "."
	}
	nhbrs := 0
	xStep := []int{0, 0, 1, 1, 1, -1, -1, -1}
	yStep := []int{1, -1, 0, 1, -1, 0, 1, -1}

	for i, _ := range xStep {
		newX := x + xStep[i]
		newY := y + yStep[i]
		if newX < len(grid[0]) && newX >= 0 && newY < len(grid) && newY >= 0 {
			if grid[y+yStep[i]][x+xStep[i]] == "#" {
				nhbrs++
			}
		}
	}
	if seat == "L" && nhbrs == 0 {
		return "#"
	} else if seat == "#" && nhbrs >= 4 {
		return "L"
	} else {
		return seat
	}
}
