package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day13/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	//leaveTime,_ := strconv.Atoi(fileArray[0])
	busIds:= fileArray[1]
	//fmt.Println(part1(leaveTime,busIds))
	fmt.Println(part2(busIds))
}

// part 2
// each number has arithmetic sequence, with common difference [id] with first term being -[position in list]
// this allows us to find one value where they all line up
//
// formula for arithmetic seq: a_n = a_1 + (n-1)*d
// where a_n is nth term, a_1 is first term, n is position in seq and d is common difference
//
//
// we dont care about n-1 so can just call this x_1
// therefore we get equations :
// y = 7*x_1
// y = 13*x_2 -1
// y = 59*x_3 -4
// etc
//
// need to find integer solution for y for all equations

func part2(busIds string) int {
	splitIds := strings.Split(busIds,",")
	ids:= []int{}
	offset := make(map[int]int)
	for i,id := range splitIds{
		if id != "x"{
			intId,_ := strconv.Atoi(id)
			ids = append(ids, intId)
			offset[intId] = i
		}
	}
	fmt.Println(offset)

	prodOfVisitedIds := ids[0]
	startPoint :=0
	for i:=1; i<len(ids);i++{
		id := ids[i]
		startPoint = findSolutions(startPoint,prodOfVisitedIds,offset[id],id)
		prodOfVisitedIds = prodOfVisitedIds * id
	}
	return startPoint
}

//

func findSolutions(startPoint, addMultiplesOf,offset, id int) int {
	for i:=1; i>0; i++{
		result := (startPoint + (addMultiplesOf*i) +offset)
		modulus :=result % id
		if modulus ==0{
			return (addMultiplesOf*i) +startPoint
		}
	}
	return 0
}



// ---------- part 1 ------------------
func part1(leaveTime int, busIds string) int {
	splitIds := strings.Split(busIds,",")
	m := make(map[int]int)
	for _,id := range splitIds{
		if id != "x"{
			intId,_ := strconv.Atoi(id)
			m[intId-(leaveTime%intId)]=intId
		}
	}
	shortestWait:= getMin(m)
	fmt.Println(shortestWait)
	return shortestWait * m[shortestWait]

}
func getMin(m map[int]int) int {
	min:=1000000000000
	for k, _ := range m {
		if k < min {
			min = k
		}
	}
	return min
}