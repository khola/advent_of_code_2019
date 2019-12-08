package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func findWayToRoot(arr map[string]string, key string) int {
	value := 1
	val, ok := arr[key]

	if ok {
		value += findWayToRoot(arr, val)
	} else {
		value = 0
	}

	return value
}

func findWayToEl(arr map[string]string, key string, el string) int {
	value := 1

	if arr[key] != el {
		value += findWayToEl(arr, arr[key], el)
	} else {
		value = 0
	}

	return value
}

func findStepsToRoot(arr map[string]string, key string) []string {
	value := []string{}
	k := key
	for {

		_, ok := arr[k]

		if ok {
			value = append(value, arr[arr[k]])
			k = arr[arr[k]]

		} else {
			break
		}
	}

	return value
}

func findCommonRoot(arr map[string]string, key1 string, key2 string) string {
	r1, r2 := findStepsToRoot(arr, key1), findStepsToRoot(arr, key2)

	for _, k1 := range r1 {
		for _, k2 := range r2 {
			if k1 == k2 {
				return k1
			}
		}
	}

	return "COM"
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input filename provided")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal("Can't open the file")
	}
	i := ""
	arr := make(map[string]string)

	for {
		_, err := fmt.Fscanln(f, &i)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		stringsSlice := strings.Split(i, ")")
		k, v := stringsSlice[0], stringsSlice[1]

		arr[v] = k

	}

	commonRoot := findCommonRoot(arr, "SAN", "YOU")
	d1 := findWayToEl(arr, "SAN", commonRoot)
	d2 := findWayToEl(arr, "YOU", commonRoot)
	result := d1 + d2
	fmt.Println(d1, d2, result)

}
