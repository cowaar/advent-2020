package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, e := ioutil.ReadFile("day18/input.txt")
	if e != nil {
		panic(e)
	}
	fileString := string(file)
	fileString = strings.Replace(fileString, " ", "", -1)
	fileArray := strings.Split(fileString, "\n")
	fmt.Println(run2(fileArray))
}

func run2(ss []string) int {
	total := 0
	for _,item := range ss {
		eq := strings.Split(item, "")
		ans := solveEq2(eq)
		fmt.Println(ans)
		total += ans
	}
	return total
}

func solveEq2(eq []string) int {
	numReg := regexp.MustCompile(`[0-9]`)
	var ans int
	for i := 0; i < len(eq); i++{
		val := eq[i]
		if val == "(" {
			close := findCloseBracket(eq, i)
			res := solveEq2(eq[i+1 : close])
			newEq := make([]string,0)
			newEq = append(newEq, eq[:i]...)
			newEq = append(newEq,strconv.Itoa(res))
			newEq = append(newEq, eq[close +1:]...)
			eq = newEq
		}
	}
	for i := 0; i < len(eq); {
		val := eq[i]
		if val == "+" {
			a, _ := strconv.Atoi(eq[i-1])
			b, _ := strconv.Atoi(eq[i+1])
			res := a + b
			newEq := make([]string,0)
			newEq = append(newEq, eq[:i-1]...)
			newEq = append(newEq,strconv.Itoa(res))
			newEq = append(newEq, eq[i+2:]...)
			eq = newEq
			i--

		}else{i++}
	}
	for i := 0; i < len(eq); i++{
		val := eq[i]
		if numReg.MatchString(val) {
			num, _ := strconv.Atoi(val)
			if i == 0 {
				ans = num
			} else {
				ans = ans * num
			}
		}
	}
	return ans
}

func run1(ss []string) int {
	total := 0
	for _, val := range ss {
		eq := strings.Split(val, "")
		ans := solveEq1(eq)
		fmt.Println(ans)
		total += ans
	}
	return total
}

func solveEq1(eq []string) int {
	numReg := regexp.MustCompile(`[0-9(]`)
	cmdReg := regexp.MustCompile(`[+*]`)
	var ans int
	var cmd string
	for i := 0; i < len(eq); {
		val := eq[i]

		if numReg.MatchString(val) {
			var num int
			var incr int
			if val == "(" {
				close := findCloseBracket(eq, i)
				num = solveEq1(eq[i+1 : close])
				incr = (close + 1) - i
			} else {
				num, _ = strconv.Atoi(val)
				incr = 1
			}
			if i == 0 {
				ans = num
			} else {
				ans = useCmd(cmd, ans, num)
			}
			i = i + incr
		} else if cmdReg.MatchString(val) {
			cmd = val
			i++
		}
	}
	return ans
}

func findCloseBracket(eq []string, startPos int) int {
	count := 1
	for j := startPos + 1; j < len(eq); j++ {
		if eq[j] == "(" {
			count++
		} else if eq[j] == ")" {
			count--
		}

		if count == 0 {
			return j
		}
	}
	return 0
}

func useCmd(cmd string, a, b int) int {
	switch cmd {
	case "+":
		return a + b
	case "*":
		return a * b
	}
	return 0
}
