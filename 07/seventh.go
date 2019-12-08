package main

import (
	"fmt"
	"strconv"
)

type instruction struct {
	code int
	m1   int
	m2   int
	m3   int
}

func getModedValue(mode int, index int, n []int) int {
	if mode == 1 {
		return n[index]
	}
	return n[n[index]]

}

func getModedIndex(mode int, index int, n []int) int {
	if mode == 1 {
		return index
	}
	return n[index]
}

func getOpCode(n int) instruction {
	s := fmt.Sprintf("%05d", n)
	code, _ := strconv.Atoi(s[3:])
	m1, _ := strconv.Atoi(s[2:3])
	m2, _ := strconv.Atoi(s[1:2])
	m3, _ := strconv.Atoi(s[0:1])
	return instruction{
		code, m1, m2, m3,
	}
}

func amplifiers(t []int, inputs []int) int {
	statesMemo := [][]int{}
	indexesMemo := []int{}
	result := 0
	for _, input := range inputs {
		r, s, i := Computer(t, []int{input, result}, 0)
		result = r
		statesMemo = append(statesMemo, s)
		indexesMemo = append(indexesMemo, i)
	}
	for {
		ii := 0
		for i := range inputs {
			r, s, j := Computer(statesMemo[i], []int{result}, indexesMemo[i])
			ii = j
			statesMemo[i] = s
			indexesMemo[i] = j
			if ii != 99 {
				result = r
			} else {
				break
			}
		}
		if ii == -1 || ii == 99 {
			break
		}
	}
	return result
}

func main() {

	t := GetInput()

	inputs := Permutations([]int{5, 6, 7, 8, 9})
	maxResult := 0

	for _, input := range inputs {
		result := amplifiers(t, input)
		if result > maxResult {
			maxResult = result
		}
	}

	fmt.Println(maxResult)

}
