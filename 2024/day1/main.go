package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	leftList := []int{}
	rightList := []int{}
	similarityMap := make(map[int]int)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "  ")
		if len(vals) == 2 {
			val, _ := strconv.ParseInt(vals[0], 10, 32)
			leftList = append(leftList, int(val))
			val, _ = strconv.ParseInt(strings.Trim(vals[1], " "), 10, 32)
			rightList = append(rightList, int(val))
			if v, ok := similarityMap[int(val)]; ok {
				similarityMap[int(val)] = v + 1
			} else {
				similarityMap[int(val)] = 1
			}
		}
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
	sum := 0
	similarityScore := 0
	for i := 0; i < len(leftList); i++ {
		if leftList[i] > rightList[i] {
			sum += leftList[i] - rightList[i]
		} else {
			sum += rightList[i] - leftList[i]
		}
		if val, ok := similarityMap[leftList[i]]; ok {
			similarityScore += leftList[i] * val
		}

	}
	println(sum, similarityScore)
}
