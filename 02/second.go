package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func computer(n []int) int {
	l := len(n) - 1
	for i := 0; i < len(n)-1; i += 4 {
		if i < l-3 && n[i+1] < l && n[i+2] < l && n[i+3] < l {
			switch n[i] {
			case 99:
				break
			case 1:
				n[n[i+3]] = n[n[i+1]] + n[n[i+2]]
			case 2:
				n[n[i+3]] = n[n[i+1]] * n[n[i+2]]
			default:
				log.Panic("Unknown instruction")
			}
		} else {
			log.Panic("Corrupted input")
		}
	}
	return n[0]
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

	tempN := make([]int, len(n))

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(tempN, n)
			tempN[1] = i
			tempN[2] = j
			result := computer(tempN)
			if result == 19690720 {
				fmt.Println(i, j, result)
				break
			}
		}
	}
}
