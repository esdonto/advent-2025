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
	for range 30 {
		reader.ReadString('\n')
	}
	line, err := reader.ReadString('\n')

	sumValid := 0

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], ": ")

		regionShape := strings.Split(lineSeparated[0], "x")
		length, err := strconv.Atoi(regionShape[0])
		if err != nil {
			panic(err)
		}
		width, err := strconv.Atoi(regionShape[1])
		if err != nil {
			panic(err)
		}

		sum := 0
		for _, v := range strings.Split(lineSeparated[1], " ") {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			sum += vInt
		}

		if 9*(length/3)*(width/3) >= sum*9 {
			sumValid++
		}
	}

	println(sumValid)
}
