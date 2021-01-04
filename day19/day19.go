package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("day19/input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(file)

	run1(fileString)
}

func run1(ss string) {
	splt := strings.Split(ss, "\n\n")
	strRules := strings.Split(splt[0], "\n")
	strInputs := strings.Split(splt[1],"\n")

	rules := makeRulesMap(strRules)
	res := make(map[string][]string)
	for num, deets := range rules {
		if strings.Contains(deets[0], "\"") {
			res[num] = []string{deets[0][1:2]}
			delete(rules, num)
		}
	}
	res = getAllValid(rules, res)
	count := 0
	for _,input := range strInputs{
		fmt.Println(input)
		_ , cont := contains(res["0"],input)
		if cont{
		count++
		}
	}
	fmt.Println(count)
}

func getAllValid(rules map[string][]string, res map[string][]string) map[string][]string {
	for len(rules) > 0 {
		for num, deets := range rules {
			found := true
			for _, char := range deets {
				if char != "|" {
					if _, ok := res[char]; !ok {
						found = false
						break
					}
				}
			}
			if found == true {
				var valid []string
				loc, cont := contains(deets, "|")
				if cont{
					valid1 := getNewValidStrings(deets[:loc], res)
					valid2 := getNewValidStrings(deets[loc+1:], res)
					comb := append(valid1, valid2...)
					valid = append(valid, comb...)
				}  else {
					valid = getNewValidStrings(deets, res)
				}
				res[num] = valid
				delete(rules, num)
			}
		}
	}
	return res
}
func getNewValidStrings(deets []string, knownRules map[string][]string) []string {
	var output []string
	var strs1 []string
	var strs2 []string
	if len(deets) == 1{
		return knownRules[deets[0]]
	}
	for _, val := range knownRules[deets[0]] {
		strs1 = append(strs1, val)
	}
	for _, val := range knownRules[deets[1]] {
		strs2 = append(strs2, val)
	}
	for _, val1 := range strs1 {
		for _, val2 := range strs2 {
			output = append(output, val1+val2)
		}
	}
	return output
}


func makeRulesMap(ss []string) map[string][]string {
	var res = make(map[string][]string)
	reg := regexp.MustCompile(`: | `)
	for _, val := range ss {
		splt := reg.Split(val, -1)
		res[splt[0]] = splt[1:]
	}
	return res
}

func contains(slice []string, val string) (int, bool) {
	for i,item := range slice {
		if item == val {
			return i,true
		}
	}
	return -1, false
}