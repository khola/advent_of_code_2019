package main

import (
	"fmt"
	"log"
	"os"
)

func fuelConsumption(n int) int {
	return n/3 - 2
}

func fuelForFuels(n int) int {
	f := fuelConsumption(n)
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

	totalModules, totalFuels, i := 0, 0, 0

	for {
		_, err := fmt.Fscanln(f, &i)
		if err != nil {
			break
		}
		forModule := fuelConsumption(i)
		totalModules += forModule
		totalFuels += fuelForFuels(forModule)
	}

	fmt.Println("Fuel for modules:")
	fmt.Println(totalModules)
	fmt.Println("Fuel for fuels:")
	fmt.Println(totalFuels)

}
