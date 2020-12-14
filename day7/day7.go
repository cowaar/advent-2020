package main

import (
	"fmt"
	"github.com/twmb/algoimpl/go/graph"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day7/test.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	// have to use a map as a set for quick lookup
	run(fileArray)
}
func run(ss []string){
	g:= graph.New(graph.Directed)
	bags := make(map[string]graph.Node, 0)
	for _, line := range ss {
		bagStrings := parseRule1(line)
		outerBag := bagStrings[0]
		innerBags := bagStrings[1:]
		bags[outerBag] = g.MakeNode()
		for i := 0; i < len(innerBags); i++ {
			if innerBags[0] == "no other bag" {
				break
			}
			bags[innerBags[i]] = g.MakeNode()
			g.MakeEdge( bags[innerBags[i]],bags[outerBag])

		}

	}
	// Make references back to the string values
	for key, node := range bags {
		*node.Value = key
	}
	goldSet := make(map[string]bool)
	for _,node := range bags {
		searchBag(node,goldSet,bags,g)
	}


}

func searchBag(node graph.Node, goldSet map[string]bool, bags map[string]graph.Node, g *graph.Graph){

	//fmt.Println(*node.Value)
	nbhrs:= g.Neighbors(node)
	fmt.Println(nbhrs)
}

func parseRule1(s string)[]string{
	findContain:=regexp.MustCompile("s contain")
	splitRule := findContain.Split(s,2)
	if splitRule[1] == " no other bags."{

	}
	findNums := regexp.MustCompile(`[0-9]\s|,|\.`)
	noNums := findNums.ReplaceAllString(splitRule[1],"")
	findBags := regexp.MustCompile(`(\w+\s\w+ bag)`)
	final := findBags.FindAllString(noNums,-1)
	return append([]string{splitRule[0]},final...)
}
