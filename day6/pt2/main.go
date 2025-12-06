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

	indexOperators := 0
	numbers := []int{}

	for columnLines := range lines[0] {
		newNumber := 0

		for rowLines := range lines {
			if lines[rowLines][columnLines] != ' ' {
				newNumber *= 10
				newNumber += int(lines[rowLines][columnLines] - '0')
			}
		}

		if newNumber != 0 {
			numbers = append(numbers, newNumber)
		} else {
			sum += applyOperator(operators[indexOperators], numbers)

			indexOperators++
			numbers = nil
		}
	}

	sum += applyOperator(operators[indexOperators], numbers)

	println("sum:", sum)
}

func applyOperator(operator string, numbers []int) int {
	switch operator {
	case "+":
		return addAll(numbers)
	case "*":
		return multAll(numbers)
	default:
		panic("Unknown operator in last line")
	}

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
