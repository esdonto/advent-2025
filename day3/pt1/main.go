package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
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
		tens := 0
		for i := range len(line)-2 {
			if line[i] > line[tens] {
				tens = i
			}
		}

		ones := tens + 1
		for i := ones; i < len(line)-1; i++ {
			if line[i] > line[ones] {
				ones = i
			}
		}

		joltage, err := strconv.Atoi(fmt.Sprintf("%c%c", line[tens], line[ones]))
		if err != nil {
			panic(err)
		}

		sum += joltage
	}
	println(sum)
}
