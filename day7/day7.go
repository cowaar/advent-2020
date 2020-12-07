package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day7/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	// have to use a map as a set for quick lookup
	m := make(map[string]map[string]bool)
	for _, line := range fileArray {
		parseRule1(line, m)
	}

	result:= 0
	for k,_ := range m{
		if searchForGolden(m,k){
			result++
		}
	}

	fmt.Println(result)
}
func searchForGolden(m map[string]map[string]bool, s string ) bool{
	contents := m[s]
	if _,ok := contents["shiny gold bag"]; ok{
		return true
	}else{ for k,_ := range contents{
		return searchForGolden(m,k)
		}
	}

	return false
}

func parseRule1(s string, m map[string]map[string]bool ) map[string]map[string]bool{
	findContain:=regexp.MustCompile("s contain")
	splitRule := findContain.Split(s,2)
	currentBag := splitRule[0]

	if splitRule[1] == " no other bags."{
		m[currentBag] = map[string]bool{}
		return m
	}
	findNums := regexp.MustCompile(`[0-9]\s|,|\.`)
	noNums := findNums.ReplaceAllString(splitRule[1],"")
	findBags := regexp.MustCompile(`(\w+\s\w+ bag)`)
	final := findBags.FindAllString(noNums,-1)
	set := make(map[string]bool)
	for _,row := range final{
		set[row]=true
	}
	m[currentBag] = set
	return m
}
