package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day17/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	run2(fileArray)
}

func run2(s []string) {
	file2d := [][]string{}
	for _, val := range s {
		file2d = append(file2d, strings.Split(val, ""))
	}
	file3d := [][][]string{file2d}
	file4d := [][][][]string{file3d}
	file4d = increaseSize2(file4d)
	input := make([][][][]string, len(file4d))
	copy(input, file4d)
	print4d(input)

	for count := 0; count < 6; count++ {

		input = increaseSize2(input)
		freshGrid := emptyFill2(len(input), len(input[0]), len(input[0][0]), len(input[0][0][0]))
		for h := 1; h < len(input)-1; h++ {
			for i := 1; i < len(input[0])-1; i++ {
				for j := 1; j < len(input[0][0])-1; j++ {
					for k := 1; k < len(input[0][0][0])-1; k++ {
						freshGrid[h][i][j][k] = updatePos2(input, k, j, i, h)
					}
				}
			}
		}
		print4d(freshGrid)
		copy(input, freshGrid)
		total := 0
		for _, w := range input {
			for _, z := range w {
				for _, y := range z {
					for _, x := range y {
						if x == "#" {
							total++
						}
					}
				}
			}
		}
		fmt.Println(total)
	}
}

func print4d(input [][][][]string) {
	for _, uni := range input {
		for _, row := range uni {
			for _, item := range row {
				fmt.Println(item)
			}
			fmt.Println("--")
		}
		fmt.Println("-----")
	}

}

func emptyFill2(hRange, iRange, jRange, kRange int) [][][][]string {
	grid := make([][][][]string, hRange)
	for h := 0; h < hRange; h++ {

		for i := 0; i < iRange; i++ {
			grid[h] = append(grid[h], [][]string{})
			for j := 0; j < jRange; j++ {
				grid[h][i] = append(grid[h][i], []string{})
				for k := 0; k < kRange; k++ {
					grid[h][i][j] = append(grid[h][i][j], ".")
				}
			}
		}
	}
	return grid
}

func increaseSize2(grid [][][][]string) [][][][]string {
	wRange := len(grid) + 2
	zRange := len(grid[0]) + 2
	yRange := len(grid[0][0]) + 2
	xRange := len(grid[0][0][0]) + 2
	newGrid := make([][][][]string, wRange)
	for h := range newGrid {
		newGrid[h] = make([][][]string, zRange)
		for i := range newGrid[h] {
			newGrid[h][i] = make([][]string, yRange)
			for j := range newGrid[h][i] {
				newGrid[h][i][j] = make([]string, xRange)
				for k := range newGrid[h][i][j] {
					newGrid[h][i][j][k] = "."
				}
			}
		}
	}
	for h := 0; h < len(grid); h++ {
		for i := 0; i < len(grid[0]); i++ {
			for j := 0; j < len(grid[0][0]); j++ {
				for k := 0; k < len(grid[0][0][0]); k++ {
					newGrid[h+1][i+1][j+1][k+1] = grid[h][i][j][k]
				}
			}
		}
	}
	return newGrid
}

func updatePos2(grid [][][][]string, x, y, z, w int) string {
	pos := grid[w][z][y][x]
	nhbrs := 0
	step := []int{-1, 0, 1}
	for _, h := range step {
		for _, i := range step {
			for _, j := range step {
				for _, k := range step {
					if i == j && j == k && k == h && h == 0 {
					} else {
						if grid[w+h][z+i][y+j][x+k] == "#" {
							nhbrs++
						}
					}
				}
			}
		}
	}

	if pos == "#" && nhbrs >= 2 && nhbrs <= 3 {
		return "#"
	} else if pos == "." && nhbrs == 3 {
		return "#"
	} else {
		return "."
	}
}

func run1(s []string) {
	file2d := [][]string{}
	for _, val := range s {
		file2d = append(file2d, strings.Split(val, ""))
	}
	file3d := [][][]string{file2d}
	file3d = increaseSize1(file3d)
	input := make([][][]string, len(file3d))
	copy(input, file3d)

	for count := 0; count < 6; count++ {

		input = increaseSize1(input)
		freshGrid := emptyFill(len(input), len(input[0]), len(input[0][0]))
		for i := 1; i < len(input)-1; i++ {
			for j := 1; j < len(input[0])-1; j++ {
				for k := 1; k < len(input[0][0])-1; k++ {
					freshGrid[i][j][k] = updatePos1(input, k, j, i)
				}
			}
		}
		print3d(freshGrid)
		copy(input, freshGrid)
		total := 0
		for _, z := range input {
			for _, y := range z {
				for _, x := range y {
					if x == "#" {
						total++
					}
				}
			}
		}
		fmt.Println(total)
	}

}
func print3d(input [][][]string) {
	for _, row := range input {
		for _, item := range row {
			fmt.Println(item)
		}
		fmt.Println("-----")
	}
}

func emptyFill(iRange, jRange, kRange int) [][][]string {
	grid := make([][][]string, iRange)
	for i := 0; i < iRange; i++ {
		for j := 0; j < jRange; j++ {
			grid[i] = append(grid[i], []string{})
			for k := 0; k < kRange; k++ {
				grid[i][j] = append(grid[i][j], ".")
			}
		}
	}
	return grid
}

func increaseSize1(grid [][][]string) [][][]string {
	zRange := len(grid) + 2
	yRange := len(grid[0]) + 2
	xRange := len(grid[0][0]) + 2
	newGrid := make([][][]string, zRange)
	for i := range newGrid {
		newGrid[i] = make([][]string, yRange)
		for j := range newGrid[i] {
			newGrid[i][j] = make([]string, xRange)
			for k := range newGrid[i][j] {
				newGrid[i][j][k] = "."
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < len(grid[0][0]); k++ {
				newGrid[i+1][j+1][k+1] = grid[i][j][k]
			}
		}
	}
	return newGrid
}

func updatePos1(grid [][][]string, x, y, z int) string {
	pos := grid[z][y][x]
	nhbrs := 0
	step := []int{-1, 0, 1}

	for _, i := range step {
		for _, j := range step {
			for _, k := range step {
				if i == j && j == k && k == 0 {
				} else {
					if grid[z+i][y+j][x+k] == "#" {
						nhbrs++
					}
				}
			}
		}
	}

	if pos == "#" && nhbrs >= 2 && nhbrs <= 3 {
		return "#"
	} else if pos == "." && nhbrs == 3 {
		return "#"
	} else {
		return "."
	}
}
