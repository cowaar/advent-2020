package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instr struct {
	code string
	val  int
}

func main() {
	file, _ := ioutil.ReadFile("day8/input.txt")
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n")
	instructions := parseInstructions(fileArray)
	fmt.Println(fixProgram(instructions))
}

func fixProgram(instructions []instr) int {
	for i, x := range instructions {
		newInstructions := instructions
		switch x.code {
		case "jmp":
			newInstructions[i] = instr{code: "nop", val: x.val}
			if ans := checkForAnswer(newInstructions); ans != 0 {
				return ans
			}
			newInstructions[i] = instr{code: "jmp", val: x.val}
			break

		case "nop":
			newInstructions[i] = instr{code: "jmp", val: x.val}
			if ans := checkForAnswer(newInstructions); ans != 0 {
				return ans
			}
			newInstructions[i] = instr{code: "nop", val: x.val}
			break
		}
	}
	return 0
}

func checkForAnswer(instructions []instr) int {
	acc := 0
	pos := 0
	visitedPos := make(map[int]bool)
	for {
		if pos == len(instructions) {
			fmt.Println("End Reached! Answer: ", acc)
			return acc
		}
		if _, ok := visitedPos[pos]; !ok {
			visitedPos[pos] = true
			pos = runOneBootCode(instructions[pos], &acc, pos)
			fmt.Println(pos, acc)
			_, ok = visitedPos[pos]
		} else {
			fmt.Println("infinite loop found!")
			break
		}
	}
	return 0

}

func part1(instructions []instr) int {
	acc := 0
	pos := 0
	visitedPos := make(map[int]bool)
	_, ok := visitedPos[pos]
	for !ok {
		visitedPos[pos] = true
		pos = runOneBootCode(instructions[pos], &acc, pos)
		_, ok = visitedPos[pos]
	}
	return acc
}

func runOneBootCode(ins instr, acc *int, pos int) int {
	switch ins.code {
	case "nop":
		return pos + 1
	case "acc":
		*acc = *acc + ins.val
		return pos + 1
	case "jmp":
		return pos + ins.val
	}
	return pos
}

func parseInstructions(s []string) []instr {
	var instructs []instr
	for _, line := range s {
		splitLine := strings.Split(line, " ")

		value, _ := strconv.Atoi(splitLine[1])
		instr := instr{code: splitLine[0], val: value}
		instructs = append(instructs, instr)
	}
	return instructs
}
