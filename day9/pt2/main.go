package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortedSlice struct {
	slice []int
}

func (s *sortedSlice) add(v int) {
	i := sort.SearchInts(s.slice, v)
	if i < len(s.slice) {
		if s.slice[i] != v {
			s.slice = append(s.slice, 0)
			copy(s.slice[i+1:], s.slice[i:])
			s.slice[i] = v
		}
	} else {
		s.slice = append(s.slice, 0)
		s.slice[i] = v
	}
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

	redTiles := [][2]int{}
	axisX := sortedSlice{[]int{0}}
	axisY := sortedSlice{[]int{0}}

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], ",")

		tileX, _ := strconv.Atoi(lineSeparated[0])
		tileY, _ := strconv.Atoi(lineSeparated[1])

		axisX.add(tileX)
		axisY.add(tileY)

		newTile := [2]int{tileX, tileY}
		redTiles = append(redTiles, newTile)
	}

	for i := range redTiles {
		redTiles[i] = [2]int{sort.SearchInts(axisX.slice, redTiles[i][0]), sort.SearchInts(axisY.slice, redTiles[i][1])}
	}

	// 0 - Inside
	// 1 - Line
	// 2 - Outside
	floor := make([][]int, len(axisX.slice)+1)
	for i := range floor {
		floor[i] = make([]int, len(axisY.slice)+1)
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
				area := getArea(redTiles[i], redTiles[j], axisX.slice, axisY.slice)
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}

	println(maxArea)
}

func fillLoop(floor [][]int) {
	// Flood fill algo
	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		coord := queue[0]
		queue = queue[1:]
		if floor[coord[0]][coord[1]] == 0 {
			floor[coord[0]][coord[1]] = 2
			if coord[0] > 0 {
				queue = append(queue, [2]int{coord[0] - 1, coord[1]})
			}
			if coord[0] < len(floor)-1 {
				queue = append(queue, [2]int{coord[0] + 1, coord[1]})
			}
			if coord[1] > 0 {
				queue = append(queue, [2]int{coord[0], coord[1] - 1})
			}
			if coord[1] < len(floor[0])-1 {
				queue = append(queue, [2]int{coord[0], coord[1] + 1})
			}
		}
	}
}

func makePerimeter(prevTile, nextTile [2]int, floor [][]int) {
	if prevTile[0] == nextTile[0] {
		if prevTile[1] < nextTile[1] {
			for y := prevTile[1]; y < nextTile[1]; y++ {
				floor[prevTile[0]][y] = 1
			}
		} else {
			for y := prevTile[1]; y > nextTile[1]; y-- {
				floor[prevTile[0]][y] = 1
			}

		}
	} else if prevTile[1] == nextTile[1] {
		if prevTile[0] < nextTile[0] {
			for x := prevTile[0]; x < nextTile[0]; x++ {
				floor[x][prevTile[1]] = 1
			}
		} else {
			for x := prevTile[0]; x > nextTile[0]; x-- {
				floor[x][prevTile[1]] = 1
			}
		}
	} else {
		panic("Two tiles are not in the same line")
	}
}

func checkRectangle(tile1, tile2 [2]int, floor [][]int) bool {
	minX, maxX := tile1[0], tile2[0]
	if tile1[0] > tile2[0] {
		minX, maxX = maxX, minX
	}
	minY, maxY := tile1[1], tile2[1]
	if tile1[1] > tile2[1] {
		minY, maxY = maxY, minY
	}

	for x := minX; x < maxX; x++ {
		if floor[x][minY] == 2 {
			return false
		}
	}
	for y := minY; y < maxY; y++ {
		if floor[maxX][y] == 2 {
			return false
		}
	}
	for x := maxX; x > minX; x-- {
		if floor[x][maxY] == 2 {
			return false
		}
	}
	for y := maxY; y > minY; y-- {
		if floor[minX][y] == 2 {
			return false
		}
	}

	return true
}

func getArea(tile1, tile2 [2]int, axisX, axisY []int) int {
	dX := axisX[tile1[0]] - axisX[tile2[0]]
	if dX < 0 {
		dX = -dX
	}
	dY := axisY[tile1[1]] - axisY[tile2[1]]
	if dY < 0 {
		dY = -dY
	}
	return (dX + 1) * (dY + 1)
}
