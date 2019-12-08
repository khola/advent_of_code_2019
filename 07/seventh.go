package main

import (
	"fmt"
)

type instruction struct {
	code int
	m1   int
	m2   int
	m3   int
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
