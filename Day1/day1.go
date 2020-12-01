package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file,_ := ioutil.ReadFile("input.txt")

	fileString := string(file)
	fileArray := strings.Split(fileString,"\n")


}

func doubleSum2020(nums []string) (int,int) {

	m := make(map[int]int)

	for _,str := range nums{

		n,_ := strconv.Atoi(str)

		complement := 2020 - n
		if _,ok := m[complement]; ok {
			return n,complement
		} else {
			m[n]=complement
		}
	}
	return 0, 0
}

