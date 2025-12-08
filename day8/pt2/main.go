package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type JunctionBox struct {
	x int
	y int
	z int
}

// Distance squared
func (b1 *JunctionBox) distanceTo(b2 *JunctionBox) int {
	dX := b1.x - b2.x
	dY := b1.y - b2.y
	dZ := b1.z - b2.z
	return dX*dX + dY*dY + dZ*dZ
}

type JunctionDistance struct {
	distance  int
	indexBox1 int
	indexBox2 int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error!")
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')

	boxes := []JunctionBox{}
	distances := []JunctionDistance{}

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], ",")

		newX, _ := strconv.Atoi(lineSeparated[0])
		newY, _ := strconv.Atoi(lineSeparated[1])
		newZ, _ := strconv.Atoi(lineSeparated[2])

		newBox := JunctionBox{newX, newY, newZ}

		for i := range boxes {
			distances = append(distances, JunctionDistance{
				distance:  newBox.distanceTo(&boxes[i]),
				indexBox1: i,
				indexBox2: len(boxes),
			})
		}

		boxes = append(boxes, newBox)
	}

	slices.SortFunc(distances, func(a, b JunctionDistance) int {
		return a.distance - b.distance
	})

	circuits := make([]int, len(boxes))
	for i := range circuits {
		circuits[i] = i
	}

	for i := range distances {
		circuitFusing := circuits[distances[i].indexBox1]
		circuitFused := circuits[distances[i].indexBox2]

		if circuitFused != circuitFusing {
			println(boxes[distances[i].indexBox1].x * boxes[distances[i].indexBox2].x)
			for j := range circuits {
				if circuits[j] == circuitFused {
					circuits[j] = circuitFusing
				}
			}
		}
	}
}
