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

	numbers := [][]int{}

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	for ; len(line) > 0 && line[0] != '*' && line[0] != '+'; line, _ = reader.ReadString('\n') {
		lineSeparated := strings.Fields(line[:len(line)-1])

		newNumbersLine := make([]int, len(lineSeparated))

		for i := range lineSeparated {
			value, err := strconv.Atoi(lineSeparated[i])
			if err != nil {
				panic(err)
			}

			newNumbersLine[i] = value
		}

		numbers = append(numbers, newNumbersLine)
	}

	sum := 0

	operators := strings.Fields(line[:len(line)-1])

	for i := range operators {
		switch operators[i] {
		case "+":
			sum += addAll(numbers, i)
		case "*":
			sum += multAll(numbers, i)
		default:
			panic("Unknown operator in last line")
		}
	}

	println("sum:", sum)
}

func addAll(numbers [][]int, index int) int {
	sum := 0

	for i := range numbers {
		sum += numbers[i][index]
	}

	return sum
}

func multAll(numbers [][]int, index int) int {
	mult := 1

	for i := range numbers {
		mult *= numbers[i][index]
	}

	return mult
}
