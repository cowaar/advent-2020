package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {

	file, _ := ioutil.ReadFile("day9/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	//fmt.Println(xmasDecoder1(fileArray, 25))
	start := time.Now()
	a1,a2 := addToTarget(fileArray, 26134589)
	elapsed := time.Since(start)
	log.Printf("func took %s", elapsed)
	fmt.Println(a1,a2,a1+a2)

}
func addToTarget(nums []string, tar int) (int,int) {

	for i, _ := range nums {
		result := 0
		for j, x := range nums[i:] {
			n, _ := strconv.Atoi(x)
			result = result + n
			if result == tar {
				return minMax(nums[i:j+i])
			}
			if result > tar {
				break
			}
		}
	}
	return 0,0
}
func minMax(s []string) (int,int) {
	min:=0
	max:=0
	for i, x := range s {
		n, _ := strconv.Atoi(x)
		if i == 0 || n < min {
			min = n
		}
		if i == 0 || n > max {
			max = n
		}
	}
	return min,max
}

func xmasDecoder1(nums []string, window int) string {
	for i, _ := range nums {
		minPos, maxPos := i, i+window
		target := nums[maxPos : maxPos+1][0]
		if !doubleSum(nums[minPos:maxPos], target) {
			return target
		}
	}
	return ""
}
func doubleSum(nums []string, target string) bool {
	m := make(map[int]int)
	for _, str := range nums {
		n, _ := strconv.Atoi(str)
		target, _ := strconv.Atoi(target)
		complement := target - n
		if _, ok := m[complement]; ok {
			return true
		} else {
			m[n] = complement
		}
	}
	return false
}
