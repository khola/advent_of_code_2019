package main

import (
	"fmt"
	"strconv"
)

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

//Computer computes
func Computer(n []int, ins []int, i int) (int, []int, int) {

	result := 0
	inputs := 0
	for i < len(n) {
		inst := getOpCode(n[i])

		switch inst.code {
		case 1:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)
			n[c] = a + b
			i += 4

		case 2:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)
			n[c] = a * b
			i += 4
		case 3:
			c := getModedIndex(inst.m1, i+1, n)
			n[c] = ins[inputs]
			inputs++
			i += 2
		case 4:
			c := getModedValue(inst.m1, i+1, n)
			i += 2
			return c, n, i
		case 5:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)

			if a != 0 {
				i = b
			} else {
				i += 3
			}
		case 6:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)

			if a == 0 {
				i = b
			} else {
				i += 3
			}
		case 7:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)

			if a < b {
				n[c] = 1
			} else {
				n[c] = 0
			}

			i += 4
		case 8:
			a := getModedValue(inst.m1, i+1, n)
			b := getModedValue(inst.m2, i+2, n)
			c := getModedIndex(inst.m3, i+3, n)
			if a == b {
				n[c] = 1
			} else {
				n[c] = 0
			}

			i += 4
		case 99:
			return result, n, 99
		}
	}
	return result, n, 99
}
