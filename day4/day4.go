package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	conv "strconv"
)

func main() {
	file, _ := ioutil.ReadFile("day4/input.txt")
	fileString := string(file)
	fmt.Println(validateBatchOfPasswords(fileString))

}

func validateBatchOfPasswords(rawString string) int {
	findSegments := regexp.MustCompile("\n\n")
	segments := findSegments.Split(rawString, -1)

	findPairs := regexp.MustCompile(`\s`)

	result := 0
	for _, x := range segments {
		fmt.Println(x)
		pairs := findPairs.Split(x, -1)
		pairSlice := regexp.MustCompile(`:`)
		m := make(map[string]string)

		for _, pair := range pairs {
			keyvalue := pairSlice.Split(pair, -1)
			m[keyvalue[0]] = keyvalue[1]
		}
		if validatePwMap2(m) == true {
			result++
		}
	}
	return result
}
func validatePwMap1(m map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	ticks := 0
	for _, x := range requiredFields {
		if _, ok := m[x]; ok {
			ticks++
		}
	}
	if ticks == 7 {
		return true
	} else {
		return false
	}
}

func validatePwMap2(m map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	ticks := 0
	for _, x := range requiredFields {
		if val, ok := m[x]; ok {
			switch x {
			case "byr":
				if byr(val) {
					ticks++
				}
			case "iyr":
				if iyr(val) {
					ticks++
				}
			case "eyr":
				if eyr(val) {
					ticks++
				}
			case "hgt":
				if hgt(val) {
					ticks++
				}
			case "hcl":
				if hcl(val) {
					ticks++
				}
			case "ecl":
				if ecl(val) {
					ticks++
				}
			case "pid":
				if pid(val) {
					ticks++
				}
			}

		}
	}

	return ticks==7
}

func checkRange(s string, min int, max int) bool {
	i, _ := conv.Atoi(s)
	if i >= min && i <= max {
		return true
	}
	return false
}

func matchRegex(s string, regex string) bool {
	hexReg := regexp.MustCompile(regex)
	return hexReg.MatchString(s)
}

func byr(s string) bool {
	return checkRange(s, 1920, 2002)
}
func iyr(s string) bool {
	return checkRange(s, 2010, 2020)
}
func eyr(s string) bool {
	return checkRange(s, 2020, 2030)
}
func hgt(s string) bool {
	regex := regexp.MustCompile(`cm|in`)
	pair := regex.FindStringIndex(s)

	if len(pair) > 0 {
		unit := s[pair[0]:]
		value := s[:pair[0]]
		switch unit {
		case "cm":
			return checkRange(value, 150, 193)
		case "in":
			return checkRange(value, 59, 76)
		}
	}
	return false
}

func hcl(s string) bool { return matchRegex(s, `^\#[a-f0-9]{6}$`) }
func pid(s string) bool { return matchRegex(s, `^[0-9]{9}$`) }
func ecl(s string) bool {
	options := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	return contains(options, s)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if reflect.DeepEqual(e, a) {
			return true
		}
	}
	return false
}
