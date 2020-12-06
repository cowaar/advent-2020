package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day6/input.txt")
	fileString := string(file)
	fmt.Println(findQCount2(fileString))

}
func findQCount2(s string) int {
	findSegments := regexp.MustCompile("\n\n")
	segments := findSegments.Split(s, -1)
	result :=0
	for j, seg := range segments {
		findPeople := regexp.MustCompile("\n")
		people := findPeople.Split(seg, -1)
		m := make(map[string]bool)

		//initialise with 1st person
		firstPerson := people[0]
		for i := 0; i < len(firstPerson); i++{
			m[firstPerson[i:i+1]] = true
		}
		fmt.Println(j, m)
		for _, person := range people {
			for k,_ := range m {
				if !strings.Contains(person,k){
					delete(m,k)
				}
			}
		}
		result = result + len(m)
	}
	return result
}

func findQCount1(s string) int {
	removeSingleNewLine := regexp.MustCompile(`\b\n\b`)
	s = removeSingleNewLine.ReplaceAllString(s, "")
	findSegments := regexp.MustCompile("\n\n")
	segments := findSegments.Split(s, -1)
	result := 0
	for _, seg := range segments {
		chars := strings.Split(seg, "")

		m := make(map[string]bool)

		for _, char := range chars {
			m[char] = true
		}
		result = result + len(m)
	}
	return result
}
