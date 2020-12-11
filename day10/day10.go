package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("day10/input.txt")
	check(err)
r:= bufio.NewReader(f)
nums,_ := ReadInts(r)
	nums=append(nums, 0)
	sort.Ints(nums)
	lastItem := nums[len(nums)-1]
	nums=append(nums,lastItem+3)
fmt.Println(part2(nums))
}

// Part 2: was a bit stuck here but used this guy's approach https://www.youtube.com/watch?v=cE88K2kFZn0&feature=youtu.be&ab_channel=JonathanPaulson
// dynamic programming is a very fancy phrase for a recursion optimisation but overall its cool
func part2(ints []int)int{
	i:= 0
	m := make(map[int]int)
	return recursiveSolution(ints,i,&m)
}
func recursiveSolution(ints []int, i int, m *map[int]int) int {
	if i == len(ints)-1{
		return 1
	}

	if val,ok := (*m)[i];ok{
		return val
	}
	ans := 0
	for _,x := range makeRange(i+1,len(ints)-1){
		if ints[x]-ints[i]<=3{
			ans += recursiveSolution(ints,x,m)
		}
	}

	(*m)[i]=ans
	return ans
}
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}


// Part 1
func part1(ints []int) int {

	ones,twos,threes:= 0,0,0
	for i,x := range ints{
		if i+1 == len(ints){
			break
		}
		nextX:= ints[i+1]
		switch nextX - x {
		case 1:
			ones++
		case 2:
			twos++
		case 3:
			threes++
		}
	}
	return ones*threes
}



func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
