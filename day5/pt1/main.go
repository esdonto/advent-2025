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

	freshCount := 0

	for line, _ = reader.ReadString('\n'); len(line) > 2; line, _ = reader.ReadString('\n') {
		id, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			panic(err)
		}

		for _, freshRange := range ranges {
			if id >= freshRange[0] && id <= freshRange[1] {
				freshCount++
				break
			}
		}
	}

	println(freshCount)
}
