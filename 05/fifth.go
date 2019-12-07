package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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

func computer(n []int) int {

	i := 0

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

			var in int
			fmt.Scanln(&in)
			c := getModedIndex(inst.m1, i+1, n)

			n[c] = in

			i += 2
		case 4:
			c := getModedValue(inst.m1, i+1, n)
			fmt.Println(c)
			i += 2
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
			return 0
		}
	}

	return 0
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input filename provided")
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Can't open the file")
	}

	fileContent := string(file)
	stringsSlice := strings.Split(fileContent, ",")
	n := []int{}

	for _, i := range stringsSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		n = append(n, j)
	}

	computer(n)

}
