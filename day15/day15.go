package main

import "fmt"

func main() {
	mem := map[int]int{0:1,5:2,4:3,1:4,10:5,14:6}
	prevNum := 7
	for i :=8; i <=30000000;i++{

		 if _,ok := mem[prevNum];ok{
			newNum := (i-1) - mem[prevNum]
			mem[prevNum] = i-1
			prevNum = newNum
		} else{
			mem[prevNum] = i - 1
			prevNum = 0}
	}
	fmt.Println(prevNum)
}

