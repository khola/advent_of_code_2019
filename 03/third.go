package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type point struct {
	x int
	y int
}

func (p point) EqualToAndNotZero(b point) bool {
	return p.x == b.x && p.y == b.y && p.x != 0 && p.y != 0
}

func linePoints(start point, operation string) []point {
	x := start.x
	y := start.y
	line := []point{}

	action := operation[0:1]
	s := string(operation[1:len(operation)])
	value, _ := strconv.Atoi(s)

	switch action {
	case "D":
		for i := 1; i <= value; i++ {
			line = append(line, point{x, y - i})
		}
	case "U":
		for i := 1; i <= value; i++ {
			line = append(line, point{x, y + i})
		}
	case "L":
		for i := 1; i <= value; i++ {
			line = append(line, point{x - i, y})
		}
	case "R":
		for i := 1; i <= value; i++ {
			line = append(line, point{x + i, y})
		}
	default:
		panic("Unknown command")
	}
	return line
}

func wire(steps []string) []point {
	first := point{0, 0}
	stepPoints := []point{}
	for _, step := range steps {
		points := linePoints(first, step)
		stepPoints = append(stepPoints, points...)
		first = points[len(points)-1]
	}
	return stepPoints
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input filename provided")
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	wires := [][]point{}
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		steps := strings.Split(string(line), ",")
		wirePoints := wire(steps)

		wires = append(wires, wirePoints)git
	}
	var distance, steps int
	for i1, pW1 := range wires[0] {
		for i2, pW2 := range wires[1] {
			if pW1.EqualToAndNotZero(pW2) {
				d := abs(pW1.x) + abs(pW2.y)
				s := i1 + i2 + 2
				if d < distance || distance == 0 {
					distance = d
				}
				if s < steps || steps == 0 {
					steps = s
				}
			}
		}
	}

	fmt.Println("Distance:")
	fmt.Println(distance)

	fmt.Println("Steps:")
	fmt.Println(steps)

}
