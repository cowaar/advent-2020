package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type status2 = struct {
	xPos int
	yPos int
	xWp  int
	yWp  int
}
type instruct = struct {
	cmd   string
	value int
}

func main() {
	file, _ := ioutil.ReadFile("day12/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	var instructs []instruct
	for _, row := range fileArray {
		val, _ := strconv.Atoi(row[1:])
		instructs = append(instructs, instruct{cmd: row[0:1], value: val})
	}
	fmt.Println(run(instructs, 2))

}
func run(instructs []instruct, part int) int {
	if part == 1 {
		stat := status1{
			x:   0,
			y:   0,
			dir: "E",
		}

		for _, ins := range instructs {
			stat = newStatus1(ins, stat)
		}
		return Abs(stat.x) + Abs(stat.y)
	}

	if part == 2 {
		stat := status2{
			xPos: 0,
			yPos: 0,
			xWp:  10,
			yWp:  1,
		}

		for _, ins := range instructs {
			stat = newStatus2(ins, stat)
		}
		return Abs(stat.xPos) + Abs(stat.yPos)
	}
	return 0
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func newStatus2(ins instruct, stat status2) status2 {
	var newStat = stat
	switch ins.cmd {
	case "N", "S", "E", "W":
		newStat.xWp, newStat.yWp = moveCoord(ins.cmd, ins.value, stat.xWp, stat.yWp)
	case "F":
		newStat.xPos, newStat.yPos = moveToWaypoint2(newStat.xPos, newStat.yPos, newStat.xWp, newStat.yWp, ins.value)
	case "R", "L":
		newStat.xWp, newStat.yWp = rotateWaypoint2(ins.cmd, stat.xWp, stat.yWp, ins.value)

	}
	fmt.Println(newStat)
	return newStat
}

func moveToWaypoint2(x, y, xWp, yWp, multi int) (int, int) {
	for i := 0; i < multi; i++ {
		x, y = x+xWp, y+yWp
	}
	return x, y
}
func rotateWaypoint2(rotation string, xWp, yWp, val int) (int,int){
	rotateInt := val / 90

	if rotation == "R" {
		for i := 0; i < rotateInt; i++ {
			newXWp := yWp
			newYWp := -xWp
			xWp,yWp = newXWp,newYWp
		}
	}
	if rotation == "L" {
		for i := 0; i < rotateInt; i++ {
			newXWp := -yWp
			newYWp := xWp
			xWp,yWp = newXWp,newYWp
		}
	}
	return xWp, yWp
}

func moveCoord(dir string, val, x, y int) (int, int) {
	switch dir {
	case "N":
		return x, y + val
	case "E":
		return x + val, y
	case "S":
		return x, y - val
	case "W":
		return x - val, y
	}
	return 0, 0
}
func newStatus1(ins instruct, stat status1) status1 {
	var newStat = stat
	switch ins.cmd {
	case "N", "S", "E", "W":
		newStat.x, newStat.y = moveCoord(ins.cmd, ins.value, stat.x, stat.y)

	case "F":
		newStat.x, newStat.y = moveCoord(stat.dir, ins.value, stat.x, stat.y)
	case "R", "L":
		newStat.dir = changeDir1(stat.dir, ins.cmd, ins.value)

	}
	return newStat
}

type status1 = struct {
	x   int
	y   int
	dir string
}

func changeDir1(currDir, rotateDir string, val int) string {
	mapDirToInt := map[string]int{"N": 0, "E": 1, "S": 2, "W": 3}
	mapIntToDir := map[int]string{0: "N", 1: "E", 2: "S", 3: "W"}
	intVal := val / 90
	currIntDir := mapDirToInt[currDir]
	var newIntDir int
	if rotateDir == "R" {
		newIntDir = (currIntDir + intVal) % 4

	}
	if rotateDir == "L" {
		newIntDir = (currIntDir - intVal + 40) % 4
	}
	return mapIntToDir[newIntDir]

}
