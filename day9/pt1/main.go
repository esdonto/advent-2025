package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error!")
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')

	redTiles := [][2]int{}
	maxArea := 0

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], ",")

		tileX, _ := strconv.Atoi(lineSeparated[0])
		tileY, _ := strconv.Atoi(lineSeparated[1])

		newTile := [2]int{tileX, tileY}

		for i := range redTiles {
			newArea := getArea(redTiles[i], newTile)
			if maxArea < newArea {
				maxArea = newArea
			}
		}

		redTiles = append(redTiles, newTile)
	}

	println(maxArea)
}

func getArea(oldTile, newTile [2]int) int {
	dX := oldTile[0] - newTile[0]
	if dX < 0 {
		dX = -dX
	}
	dY := oldTile[1] - newTile[1]
	if dY < 0 {
		dY = -dY
	}
	return (dX + 1) * (dY + 1)
}
