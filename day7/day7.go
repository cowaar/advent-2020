package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day7/test.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	// have to use a map as a set for quick lookup
	m := make(map[string]map[string]bool)
	for _, line := range fileArray {
		parseRule1(line, m)
	}
	result:= 0
	for k,_ := range m{
		searchForGolden(m,k,&result)
	}

	fmt.Println(result)
	fmt.Println(len(m))
}


func searchForGolden(m map[string]map[string]bool, s string , res *int){
	contents := m[s]
	for k,_ := range contents {
		if k == "shiny gold bag"{
			fmt.Println(k)
			*res++

			break

		}else {
			searchForGolden(m,k,res)
		}
	}

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
