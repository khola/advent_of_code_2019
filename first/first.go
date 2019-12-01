package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fuelForModule(n int) int {
	return n/3 - 2
}

func fuelForFuels(n int) int {
	f := fuelForModule(n)
	if f <= 0 {
		return n
	}
	return n + fuelForFuels(f)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input filename provided")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal("Can't open the file")
	}

	totalModules := 0
	totalFuels := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		forModule := fuelForModule(i)
		totalModules += forModule
		totalFuels += fuelForFuels(forModule)
	}

	fmt.Println("Fuel for modules:")
	fmt.Println(totalModules)
	fmt.Println("Fuel for fuels:")
	fmt.Println(totalFuels)

}
