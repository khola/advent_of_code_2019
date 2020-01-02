package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	Computer(n)

}
