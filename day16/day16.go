package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day16/input.txt")
	fileString := string(file)
	run2(fileString)
}

func run2(s string) {
	segs := strings.Split(s, "\n\n")
	fstr := strings.Split(segs[0], "\n")
	fieldRegex := regexp.MustCompile(`\: | or `)
	fields := [][]string{}
	for _, k := range fstr {
		fields = append(fields, fieldRegex.Split(k, -1))
	}

	allRange := getFullRange(fields)
	tickets := segs[2]
	ticketNums := [][]int{}
	// get valid tickets
	for _, ticket := range strings.Split(tickets, "\n")[1:] {
		row := []int{}
		valid := true
		for _, y := range strings.Split(ticket, ",") {
			num, _ := strconv.Atoi(y)
			if _, ok := allRange[num]; !ok {
				valid = false
			}
			row = append(row, num)
		}
		if valid {
			ticketNums = append(ticketNums, row)
		}
	}
	fieldRange := getFieldRangeMap(fields)
	// find valid positions
	possiblePositions := make(map[string][]int)
	for k, _ := range fieldRange {
		for i := 0; i < len(ticketNums[0]); i++ {
			valid := true
			for j := range ticketNums {
				if _, ok := fieldRange[k][ticketNums[j][i]]; !ok {
					valid = false
				}
			}
			if valid {
				possiblePositions[k] = append(possiblePositions[k], i)
			}
		}
	}
	//solve the order
	order := solveOrder(possiblePositions)
	myTicket := segs[1]
	myNumbers :=  regexp.MustCompile(`\n|,`).Split(myTicket, -1)[1:]
	finalProd := 1
	for k,v := range order{
		if startsWith(k,"departure"){
			multiplier,_ := strconv.Atoi(myNumbers[v])
			finalProd = finalProd * multiplier
		}
	}
	fmt.Println(finalProd)



}

func solveOrder(orders map[string][]int) map[string]int {
	output := make(map[string]int)
	for len(orders) > 0 {
		for k, v := range orders {
			if len(v) == 1 {
				output[k] = v[0]
				fmt.Println(output)
				delete(orders, k)
				for k,vals := range orders{
					orders [k] = remove(vals,v[0])
				}
			}
		}
	}
	return output
}

func getFieldRangeMap(fields [][]string) map[string]map[int]bool {
	mem := make(map[string]map[int]bool)
	for _, field := range fields {
		fieldLabel := field[0]
		fieldRange := make(map[int]bool)
		for i := 1; i <= 2; i++ {
			string1 := strings.Split(field[i], "-")
			num1, _ := strconv.Atoi(string1[0])
			num2, _ := strconv.Atoi(string1[1])
			for _, num := range makeRange(num1, num2) {
				fieldRange[num] = true
			}
		}
		mem[fieldLabel] = fieldRange
	}
	return mem
}

func run1(s string) {
	segs := strings.Split(s, "\n\n")
	fstr := strings.Split(segs[0], "\n")
	fieldRegex := regexp.MustCompile(`\: | or `)
	fields := [][]string{}
	for _, k := range fstr {
		fields = append(fields, fieldRegex.Split(k, -1))
	}

	ranges := getFullRange(fields)
	tickets := segs[2]
	ticketNos := getAllFields(tickets)

	total := 0
	for _, x := range ticketNos {

		if _, ok := ranges[x]; !ok {
			total = x + total
		}

	}
	fmt.Println(total)
}

func getAllFields(tickets string) []int {
	ticketNos := []int{}
	for _, x := range regexp.MustCompile(`\n|,`).Split(tickets, -1)[1:] {
		int, _ := strconv.Atoi(x)
		ticketNos = append(ticketNos, int)
	}
	return ticketNos
}

func getFullRange(fields [][]string) map[int]bool {
	rng := make(map[int]bool)
	for _, field := range fields {
		for i := 1; i <= 2; i++ {
			string1 := strings.Split(field[i], "-")
			num1, _ := strconv.Atoi(string1[0])
			num2, _ := strconv.Atoi(string1[1])
			for _, num := range makeRange(num1, num2) {
				rng[num] = true
			}
		}
	}
	return rng
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func remove(s []int, target int) []int {
	for i,v := range s{
		if v == target{
			s[len(s)-1], s[i] = s[i], s[len(s)-1]
			return s[:len(s)-1]
		}
	}
	return nil
}

func startsWith(s, subs string) bool {
	if len(s) < len(subs) {
		return false
	} else if s[0:len(subs)] == subs {
		return true
	}
	return false
}