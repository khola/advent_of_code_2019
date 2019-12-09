package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func getInput() []byte {
	if len(os.Args) < 2 {
		log.Fatal("No input filename provided")
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Can't open the file")
	}

	fileContent := string(file)
	stringsSlice := []byte(fileContent)
	return stringsSlice
}

func findValueForPixel(index int, image [][]byte) int {
	one, zero := byte('1'), byte('0')

	for _, layer := range image {
		if layer[index] == one {
			return 1
		}
		if layer[index] == zero {
			return 0
		}
	}
	return 2
}

func main() {
	bytes := getInput()

	columns := 25
	rows := 6
	layerSize := columns * rows
	layersCount := len(bytes) / layerSize
	image := [][]byte{}

	for i := 0; i < layersCount; i++ {
		start, end := i*layerSize, (i+1)*layerSize
		layer := bytes[start:end]
		image = append(image, layer)
	}
	finalImage := []int{}
	for i := 0; i < layerSize; i++ {
		finalImage = append(finalImage, findValueForPixel(i, image))
	}

	for i, pixel := range finalImage {
		if pixel == 0 {
			fmt.Print("█")
		} else {
			fmt.Print("░")
		}
		if (i+1)%columns == 0 {
			fmt.Println()
		}
	}

}
