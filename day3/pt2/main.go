package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error!")
		panic(err)
	}
	defer file.Close()

	sum := 0

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	for ; err==nil; line, err = reader.ReadString('\n'){
		joltage := 0

		size := 12
		for i := range size {
			digit_pos := getDigitPos(line[:len(line)-size+i])

			joltage *= 10
			joltage += int(line[digit_pos] - '0')

			line = line[digit_pos+1:]
		}

		sum += joltage
	}
	println(sum)
}

func getDigitPos(bank string) int {
	pointer := 0
	for i := range len(bank) {
		if bank[i] > bank[pointer] {
			pointer = i
		}
	}
	return pointer
}
