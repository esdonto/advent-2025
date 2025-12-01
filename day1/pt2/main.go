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

	pointer := 50
	counter := 0

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	for ; err==nil; line, err = reader.ReadString('\n'){
		dist, err := strconv.Atoi(line[1:len(line)-1])
		if err != nil {
			panic(err)
		}

		if line[0] == 'L' {
			counter += (dist - (pointer - 100) % 100) / 100
			pointer -= dist
		} else if line[0] == 'R' {
			pointer += dist
			counter += pointer / 100
		} else {
			panic("Weird start of line")
		}

		pointer %= 100
		if pointer < 0 {
			pointer += 100
		}
	}

	println(counter)
}
