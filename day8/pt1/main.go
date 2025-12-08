package main

import (
	"bufio"
	"fmt"
	"math"
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

func (b1 *JunctionBox) distanceTo(b2 *JunctionBox) float64 {
	dX := b1.x - b2.x
	dY := b1.y - b2.y
	dZ := b1.z - b2.z
	return math.Sqrt(float64(dX*dX + dY*dY + dZ*dZ))
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
	distacesMatrix := [][]float64{}

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], ",")

		newX, _ := strconv.Atoi(lineSeparated[0])
		newY, _ := strconv.Atoi(lineSeparated[1])
		newZ, _ := strconv.Atoi(lineSeparated[2])

		newBox := JunctionBox{newX, newY, newZ}

		newRow := make([]float64, len(boxes)+1)
		for i := range boxes {
			newDist := newBox.distanceTo(&boxes[i])
			newRow[i] = newDist
		}

		for i := range distacesMatrix {
			distacesMatrix[i] = append(distacesMatrix[i], newRow[i])
		}
		distacesMatrix = append(distacesMatrix, newRow)

		boxes = append(boxes, newBox)
	}

	circuits := make([]int, len(boxes))
	for i := range circuits {
		circuits[i] = i
	}

	var previousDist float64 = 0
	for range 1000 {
		var minI, minJ int
		var minDistance float64 = -1

		// Going throught the matrix
		for i := range distacesMatrix {
			for j := range i {
				if distacesMatrix[i][j] > previousDist && (minDistance < 0 || minDistance > distacesMatrix[i][j]) {
					minI = i
					minJ = j
					minDistance = distacesMatrix[i][j]
				}
			}
		}

		circuitFusing := circuits[minI]
		circuitFused := circuits[minJ]

		if circuitFused != circuitFusing {
			for i := range circuits {
				if circuits[i] == circuitFused {
					circuits[i] = circuitFusing
				}
			}
		}

		previousDist = minDistance
		minDistance = -1
	}

	lenCircuits := []int{}

	for i := range boxes {
		len := 0
		for _, circuit := range circuits {
			if i == circuit {
				len++
			}
		}
		if len > 0 {
			lenCircuits = append(lenCircuits, len)
		}
	}

	slices.SortFunc(lenCircuits, func(a, b int) int {
		return b - a
	})

	for _, v := range circuits {
		print(v, "-")
	}

	println(lenCircuits[0] * lenCircuits[1] * lenCircuits[2])
}
