package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func pow10(y int) int {
	r := 10
	for i := 2; i < y; i++ {
		r = r * 10
	}
	return r
}

func countDigits(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func match(input int) bool {
	digits := []int{}
	count := countDigits(input)
	if count != 6 {
		return false
	}

	for i := count; i > 1; i-- {
		digit := input / pow10(i)
		input = input - digit*pow10(i)
		digits = append(digits, digit)
	}
	digits = append(digits, input%10)

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false

		}
	}
	matches := 0
	i := 0
	max := len(digits) - 1

	for {
		ii := 1
		for {
			if i+ii <= max {
				if digits[i] == digits[i+ii] {
					ii++
				} else {
					break
				}
			} else {
				break
			}
		}

		if ii == 2 {
			matches++

		}
		i = i + ii
		if i >= max {
			break
		}
	}

	if matches > 0 {
		return true
	}
	return false
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please provide interval")
	}

	from, _ := strconv.Atoi(os.Args[1])
	to, _ := strconv.Atoi(os.Args[2])
	matches := 0
	for i := from; i <= to; i++ {
		if match(i) {
			matches++
		}
	}
	fmt.Println(matches)

}
