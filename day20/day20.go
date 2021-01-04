package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// edge directions defined as:
//             0
//
//          ------->
//
//          +----+
//   3  |   |    |   |   1
//      |   |    |   |
//      v   +----+   v
//
//          ------->
//
//             2
//
//             4
//
//          <-------
//
//      ^   +----+   ^
//   7  |   |    |   |   5
//      |   |    |   |
//          +----+
//
//          <-------
//
//             6

func main() {
	file, err := ioutil.ReadFile("day20/test.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(file)
	fileArray := strings.Split(fileString, "\n\n")
	run2(fileArray)
}

typ

type edgeCoords struct {
	tile int
	dir  int
}

func run2(ss []string) {
	allEdges := getMapOfBothDirEdges(ss)
	allTiles := getMapOfTiles(ss)
	nhbrEdges := make(map[int]map[int]edgeCoords)
	for k, v := range allEdges {
		nhbrEdges[k] = make(map[int]edgeCoords)
		for i := 0; i < 4; i++ {
			for nK, nV := range allEdges {
				if nK == k {
					continue
				}
				for j, nEdge := range nV {
					if reflect.DeepEqual(v[i], nEdge) {
						nhbrEdges[k][i] = edgeCoords{
							tile: nK,
							dir:  j,
						}
					}
				}
			}
		}
	}
//	make one row
	for k,v := range nhbrEdges{
		if len(v) == 2{
			nextTileCoords := v[0]




			break
		}
	}

}

func getMapOfTiles(ss []string) map[int][][]string {
	mapOfTiles := make(map[int][][]string)
	for _, item := range ss {
		lines := strings.Split(item, "\n")
		label, _ := strconv.Atoi(regexp.MustCompile(`Tile |:`).ReplaceAllString(lines[0], ""))
		tile := lines[1:]
		rows := make([][]string, len(tile))
		for i, line := range tile {
			chars := strings.Split(line, "")
			rows[i] = chars
		}
		mapOfTiles[label] = rows
	}
	return mapOfTiles

}
func run1(ss []string) {
	mList := getMapOfBothDirEdges(ss)
	fmt.Println(findCornersProduct(mList))

}

func getMapOfBothDirEdges(ss []string) map[int][][]string {

	mList := make(map[int][][]string)
	for _, item := range ss {
		lines := strings.Split(item, "\n")
		label, _ := strconv.Atoi(regexp.MustCompile(`Tile |:`).ReplaceAllString(lines[0], ""))
		tile := lines[1:]
		edges := make([][]string, len(tile))
		for i, line := range tile {
			chars := strings.Split(line, "")
			if i == 0 {
				edges[0] = chars
			}
			if i == len(tile)-1 {
				edges[2] = chars
			}
			edges[1] = append(edges[1], chars[len(tile)-1])
			edges[3] = append(edges[3], chars[0])
		}
		edges[4], edges[5], edges[6], edges[7] = reverse(edges[0]), reverse(edges[1]), reverse(edges[2]), reverse(edges[3])

		mList[label] = edges
	}
	return mList
}

func findCornersProduct(mList map[int][][]string) int {
	nbhrs := make(map[int]int)

	for k, v := range mList {
		count := 0
		for i := 0; i < 4; i++ {
			for nK, nV := range mList {
				if nK == k {
					continue
				}
				for _, nEdge := range nV {
					if reflect.DeepEqual(v[i], nEdge) {
						count++

					}
				}
			}
		}
		nbhrs[k] = count
	}
	prod := 1
	for k, v := range nbhrs {
		if v == 2 {
			prod = prod * k
		}
	}
	return prod
}

func reverse(s []string) []string {
	result := make([]string, len(s))
	copy(result, s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
