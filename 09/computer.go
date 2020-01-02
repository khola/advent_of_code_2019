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

func getModedIndex(mode int, index int, n []int, relative int) int {

	if mode == 1 {
		return index
	}
	if mode == 2 {
		return n[index] + relative
	}
	return n[index]
}

func resizeMemory(n []int, indexes []int) []int {
	l := len(n)
	max := 0
	for i, e := range indexes {
		if i == 0 || e > max {
			max = e
		}
	}
	diff := max - l
	if diff > -1 {
		for i := -1; i < diff; i++ {
			n = append(n, 0)
		}
	}
	return n
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
func Computer(n []int) int {

	i := 0
	relative := 0

	for i < len(n) {
		inst := getOpCode(n[i])
		switch inst.code {
		case 1:
			a := getModedIndex(inst.m1, i+1, n, relative)
			b := getModedIndex(inst.m2, i+2, n, relative)
			c := getModedIndex(inst.m3, i+3, n, relative)
			in := []int{a, b, c}
			n = resizeMemory(n, in)

			n[c] = n[a] + n[b]
			i += 4

		case 2:
			a := getModedIndex(inst.m1, i+1, n, relative)
			b := getModedIndex(inst.m2, i+2, n, relative)
			c := getModedIndex(inst.m3, i+3, n, relative)
			in := []int{a, b, c}
			n = resizeMemory(n, in)
			n[c] = n[a] * n[b]
			i += 4
		case 3:
			var in int
			fmt.Scanln(&in)
			c := getModedIndex(inst.m1, i+1, n, relative)
			inx := []int{c}
			n = resizeMemory(n, inx)
			n[c] = in
			i += 2
		case 4:
			c := getModedIndex(inst.m1, i+1, n, relative)
			inx := []int{c}
			n = resizeMemory(n, inx)
			fmt.Println(n[c])
			i += 2
		case 5:
			a := getModedIndex(inst.m1, i+1, n, relative)
			b := getModedIndex(inst.m2, i+2, n, relative)
			inx := []int{a, b}
			n = resizeMemory(n, inx)
			if n[a] != 0 {
				i = n[b]
			} else {
				i += 3
			}
		case 6:
			a := getModedIndex(inst.m1, i+1, n, relative)
			b := getModedIndex(inst.m2, i+2, n, relative)
			inx := []int{a, b}
			n = resizeMemory(n, inx)
			if n[a] == 0 {
				i = n[b]
			} else {
				i += 3
			}
		case 7:
			a := getModedIndex(inst.m1, i+1, n, relative)
			b := getModedIndex(inst.m2, i+2, n, relative)
			c := getModedIndex(inst.m3, i+3, n, relative)
			inx := []int{a, b, c}
			n = resizeMemory(n, inx)
			if n[a] < n[b] {
				n[c] = 1
			} else {
				n[c] = 0
			}

			i += 4
		case 8:
			a := getModedIndex(inst.m1, i+1, n, relative)
			b := getModedIndex(inst.m2, i+2, n, relative)
			c := getModedIndex(inst.m3, i+3, n, relative)
			inx := []int{a, b, c}
			n = resizeMemory(n, inx)
			if n[a] == n[b] {
				n[c] = 1
			} else {
				n[c] = 0
			}

			i += 4
		case 9:
			a := getModedIndex(inst.m1, i+1, n, relative)
			inx := []int{a}
			n = resizeMemory(n, inx)
			relative += n[a]
			i += 2
		case 99:
			return -1
		}
	}
	return -1
}
