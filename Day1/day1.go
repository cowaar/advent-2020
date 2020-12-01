package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//testData := []string{"1721", "979", "366", "299", "675", "1456"}
	//fmt.Println(tripleSum2020(testData))

	file, _ := ioutil.ReadFile("input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	fmt.Println(tripleSum2020(fileArray))

}

func doubleSum2020(nums []string) (int, int) {

	m := make(map[int]int)

	for _, str := range nums {

		n, _ := strconv.Atoi(str)

		complement := 2020 - n
		if _, ok := m[complement]; ok {
			return n, complement
		} else {
			m[n] = complement
		}
	}
	return 0, 0
}

func tripleSum2020(nums []string) (int, int,int) {

	m := make(map[int]int)
	for i, str := range nums[:len(nums)-1] {

		n1, _ := strconv.Atoi(str)
		otherTwoNumbers := 2020 - n1
		for _, str := range nums[i+1:]{
			n2, _ := strconv.Atoi(str)
			n3 := otherTwoNumbers - n2

			if _,ok := m[n3];ok {
				return n1,n2,n3
			} else{
				m[n2]=0
			}
		}
	}
	return 0, 0, 0
}
