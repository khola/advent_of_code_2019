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

func newPoint(x int, y int) *point {
	p := point{x, y}
	return &p
}

func linePoints(start *point, operation string) []*point {
	x := start.x
	y := start.y
	line := []*point{}

	action := operation[0:1]
	s := string(operation[1:len(operation)])
	value, _ := strconv.Atoi(s)

	switch action {
	case "D":
		for i := 1; i <= value; i++ {
			line = append(line, newPoint(x, y-i))
		}
	case "U":
		for i := 1; i <= value; i++ {
			line = append(line, newPoint(x, y+i))
		}
	case "L":
		for i := 1; i <= value; i++ {
			line = append(line, newPoint(x-i, y))
		}
	case "R":
		for i := 1; i <= value; i++ {
			line = append(line, newPoint(x+i, y))
		}
	default:
		panic("Unknown command")
	}
	return line
}

func wire(steps []string) []*point {
	first := newPoint(0, 0)
	stepPoints := []*point{}
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
	wires := [][]*point{}
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

		wires = append(wires, wirePoints)

	}
	var distance, steps int

	for i1, pointW1 := range wires[0] {
		for i2, pointW2 := range wires[1] {
			if pointW1.x == pointW2.x && pointW1.y == pointW2.y && pointW1.x != 0 && pointW1.y != 0 {
				d := abs(pointW1.x) + abs(pointW2.y)
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
