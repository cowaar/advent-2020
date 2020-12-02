package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type password struct {
	min int
	max  int
	letter string
	pw string
}

func main() {
	file, _ := ioutil.ReadFile("day2/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")

	//testcase := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}


	result := validatePasswords2(fileArray)
	fmt.Println(result)

}

func validatePasswords1(pws []string) int {
	result := 0
	for _,x := range pws{
		splitSpaces := strings.Split(x," ")
		minMax := strings.Split(splitSpaces[0], "-")
		min,_ := strconv.Atoi(minMax[0])
		max,_ := strconv.Atoi(minMax[1])
		letter := splitSpaces[1][:1]
		pw := splitSpaces[2]
		letterRegex := regexp.MustCompile(letter)
		matches := len(letterRegex.FindAllStringIndex(pw, -1))
		if matches >= min && matches <= max{
			result++
		}
	}
	return result
}

func validatePasswords2(pws []string) int {
	result := 0
	for _,x := range pws{
		splitSpaces := strings.Split(x," ")
		n1n2 := strings.Split(splitSpaces[0], "-")
		n1,_ := strconv.Atoi(n1n2[0])
		n2,_ := strconv.Atoi(n1n2[1])
		letter := splitSpaces[1][:1]
		pw := splitSpaces[2]
		pwRune := []rune(pw)
		fmt.Println(n1,n2, pw,pwRune)
		n1present := (string(pwRune[n1-1]) == letter)
		n2present := (string(pwRune[n2-1]) == letter)
		println(pw[n1:n1])
		if n1present != n2present{
			result++
		}
	}
	return result
}