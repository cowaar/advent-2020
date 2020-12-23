package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("day14/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	run2(fileArray)
}

func run2(ss []string) {
	mem := make(map[int]int)
	mask := ""
	for _, s := range ss {
		pair := strings.Split(s, " = ")
		if pair[0] == "mask" {
			mask = pair[1]
		} else if startsWith(pair[0], "mem") {
			regexNum := regexp.MustCompile(`[0-9]+`)
			pos, _ := strconv.Atoi(regexNum.FindString(pair[0]))
			val, _ := strconv.Atoi(pair[1])
			maskedPos := applyMask2(mask, int64(pos))
			fmt.Println(pos,val)
			for _,pos := range maskedPos{
				mem[pos] = val
		}
	}
	total := 0
	for _, v := range mem {
		total += v
	}
	fmt.Println(total)
}
}

func applyMask2(mask string, loc int64) []int {
	maskSlice := strings.Split(mask, "")
	bin := strconv.FormatInt(loc, 2)
	digits := strings.Split(bin, "")
	for len(digits) < 36 {
		digits = append([]string{"0"}, digits...)
	}
	for i, char := range maskSlice {
		if char == "1" {
			digits[i] = "1"
		}
	}
	memList := []string{}

	for i,char := range maskSlice{
		if char == "X"{
			memList = addFloatingBit(memList)
		} else {
			memList = addBit(memList, digits[i])
		}
	}
	memIntList := []int{}
	for _,k := range memList{
		valInt64,_ := strconv.ParseInt(k, 2, 0)
		memIntList = append(memIntList,int(valInt64))
	}
	return memIntList

}

func addBit (inp []string , digit string) []string {
	if len(inp)==0{
		inp = []string{""}
	}
	result := inp
	for i, _ := range inp {
		result[i] = result[i] +digit
	}
	return result
}

func addFloatingBit(inputs []string) []string {
	if len(inputs)==0{
		inputs = []string{""}
	}
	noOfMems := len(inputs)
	for i := 0; i < noOfMems; i++{
		inputs = append(inputs, inputs[i])
		inputs[i] = inputs[i] + "0"
	}
	for i:=noOfMems; i<len(inputs);i++{
		inputs[i] = inputs[i] + "1"
	}
	return inputs
}

func run1(ss []string) int {
	mem := make(map[int]int)
	mask := ""
	for _, s := range ss {
		pair := strings.Split(s, " = ")
		if pair[0] == "mask" {
			mask = pair[1]
		} else if startsWith(pair[0], "mem") {
			regexNum := regexp.MustCompile(`[0-9]+`)
			pos, _ := strconv.Atoi(regexNum.FindString(pair[0]))
			val, _ := strconv.Atoi(pair[1])
			maskedVal := applyMask1(mask, int64(val))
			mem[pos] = maskedVal
		}
	}
	total := 0
	for _, v := range mem {
		total += v
	}
	return total
}

func applyMask1(mask string, number int64) int {
	maskSlice := strings.Split(mask, "")
	bin := strconv.FormatInt(number, 2)
	digits := strings.Split(bin, "")
	for len(digits) < 36 {
		digits = append([]string{"0"}, digits...)
	}

	for i, char := range maskSlice {
		if char == "1" {
			digits[i] = "1"
		} else if char == "0" {
			digits[i] = "0"
		}
	}
	resultBin := strings.Join(digits, "")
	result, _ := strconv.ParseInt(resultBin, 2, 0)
	return int(result)
}

func startsWith(s, subs string) bool {
	if len(s) < len(subs) {
		return false
	} else if s[0:len(subs)] == subs {
		return true
	}
	return false
}
