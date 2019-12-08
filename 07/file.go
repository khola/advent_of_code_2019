package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

//GetInput reads provided file and return list of ints
func GetInput() []int {
	if len(os.Args) < 2 {
		log.Fatal("No input filename provided")
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Can't open the file")
	}

	fileContent := string(file)
	stringsSlice := strings.Split(fileContent, ",")
	t := []int{}

	for _, i := range stringsSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		t = append(t, j)
	}

	return t
}
