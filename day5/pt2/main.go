package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	ranges := [][2]int{}

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	for ; len(line) > 2; line, _ = reader.ReadString('\n') {
		scope_separated := strings.Split(line[:len(line)-1], "-")

		start, err := strconv.Atoi(scope_separated[0])
		if err != nil {
			panic(err)
		}

		stop, err := strconv.Atoi(scope_separated[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, [2]int{start, stop})
	}

	// Sorting slices by start

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// Merge intervals algorithm

	freshCount := 0
	buffer := ranges[0]

	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] > buffer[1] {
			freshCount += buffer[1] - buffer[0] + 1
			buffer = ranges[i]
		} else if ranges[i][1] > buffer[1] {
			buffer[1] = ranges[i][1]
		}
	}
	freshCount += buffer[1] - buffer[0] + 1

	println(freshCount)
}
