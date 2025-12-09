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
	maxX, maxY := 0, 0

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], ",")

		tileX, _ := strconv.Atoi(lineSeparated[0])
		tileY, _ := strconv.Atoi(lineSeparated[1])

		if maxX < tileX {
			maxX = tileX
		}
		if maxY < tileY {
			maxY = tileY
		}

		newTile := [2]int{tileX, tileY}
		redTiles = append(redTiles, newTile)
	}

	floor := make([][]bool, maxX+1)
	for i := range floor {
		floor[i] = make([]bool, maxY+1)
	}

	for i := range len(redTiles) - 1 {
		makePerimeter(redTiles[i], redTiles[i+1], floor)
	}
	makePerimeter(redTiles[len(redTiles)-1], redTiles[0], floor)

	fillLoop(floor)

	maxArea := 0

	for i := range redTiles {
		for j := range redTiles {
			if i != j && checkRectangle(redTiles[i], redTiles[j], floor) {
				area := getArea(redTiles[i], redTiles[j])
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}

	println(maxArea)

}

func fillLoop(floor [][]bool) {
	// Raster fill algo
	for i := range floor {
		brush := false
		for j := range floor[i] {
			if floor[i][j] {
				brush = !brush
			} else if brush {
				floor[i][j] = true
			}
		}
	}
}

func makePerimeter(prevTile, nextTile [2]int, floor [][]bool) {
	if prevTile[0] == nextTile[0] {
		for y := prevTile[1]; y < nextTile[1]; y++ {
			floor[prevTile[0]][y] = true
		}
	} else if prevTile[1] == nextTile[1] {
		for x := prevTile[0]; x < nextTile[0]; x++ {
			floor[x][prevTile[1]] = true
		}
	} else {
		panic("Two tiles are not in the same line")
	}
}

func checkRectangle(tile1, tile2 [2]int, floor [][]bool) bool {
	minX, maxX := tile1[0], tile2[0]
	if tile1[0] > tile2[0] {
		minX, maxX = maxX, minX
	}
	minY, maxY := tile1[1], tile2[1]
	if tile1[1] > tile2[1] {
		minY, maxY = maxY, minY
	}

	for x := minX; x < maxX; x++ {
		if !floor[x][minY] {
			return false
		}
	}
	for y := minY; y < maxY; y++ {
		if !floor[maxX][y] {
			return false
		}
	}
	for x := maxX; x > minX; x-- {
		if !floor[x][maxY] {
			return false
		}
	}
	for y := maxY; y > minY; y-- {
		if !floor[minX][y] {
			return false
		}
	}

	return true
}

func getArea(tile1, tile2 [2]int) int {
	dX := tile1[0] - tile2[0]
	if dX < 0 {
		dX = -dX
	}
	dY := tile1[1] - tile2[1]
	if dY < 0 {
		dY = -dY
	}
	return (dX + 1) * (dY + 1)
}
