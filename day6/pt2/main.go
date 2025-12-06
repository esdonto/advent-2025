package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error!")
		panic(err)
	}
	defer file.Close()

	lines := []string{}

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	for ; len(line) > 0 && line[0] != '*' && line[0] != '+'; line, _ = reader.ReadString('\n') {
		lines = append(lines, line[:len(line)-1])
	}

	sum := 0

	operators := strings.Fields(line[:len(line)-1])
	columnLines := 0

	for i := range operators {
		numbers := []int{}

		for {
			if columnLines >= len(lines[0]) {
				break
			}

			newNumber := 0
			for rowLines := range lines {
				if lines[rowLines][columnLines] != ' ' {
					newNumber *= 10
					newNumber += int(lines[rowLines][columnLines] - '0')
				}
			}

			columnLines++

			if newNumber != 0 {
				numbers = append(numbers, newNumber)
			} else {
				break
			}
		}

		switch operators[i] {
		case "+":
			sum += addAll(numbers)
		case "*":
			sum += multAll(numbers)
		default:
			panic("Unknown operator in last line")
		}
	}

	println("sum:", sum)
}

func addAll(numbers []int) int {
	sum := 0

	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

func multAll(numbers []int) int {
	mult := 1

	for i := range numbers {
		mult *= numbers[i]
	}

	return mult
}
