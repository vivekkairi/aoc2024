package main

import (
	"bufio"
	"errors"
	"fmt"
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
		val1, val2, err := parseValue(scanner.Text())
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, val1)
		rightList = append(rightList, val2)
		if v, ok := similarityMap[val2]; ok {
			similarityMap[val2] = v + 1
		} else {
			similarityMap[val2] = 1
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
	fmt.Printf("Part 1: %v\nPart 2: %v", sum, similarityScore)
}

func parseValue(s string) (int, int, error) {
	vals := strings.Split(s, "  ")
	if len(vals) != 2 {
		return 0, 0, errors.New("invalid input")
	}
	left, err := strconv.ParseInt(vals[0], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	right, err := strconv.ParseInt(strings.Trim(vals[1], " "), 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return int(left), int(right), nil
}
