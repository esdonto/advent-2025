package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error!")
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')

	// Number of possible paths for the tachyon to take to arrive at this position
	buffer := make([]int, len(line)-1)
	for i := range line {
		if line[i] == 'S' {
			buffer[i] = 1
		}
	}
	newBuffer := make([]int, len(line)-1)

	for ; err == nil; line, err = reader.ReadString('\n') {
		for i := range buffer {
			if buffer[i] > 0 {
				if line[i] == '^' {
					if i > 0 {
						newBuffer[i-1] += buffer[i]
					}
					if i < len(buffer)-1 {
						newBuffer[i+1] += buffer[i]
					}
				} else {
					newBuffer[i] += buffer[i]
				}
			}
		}
		buffer = newBuffer
		newBuffer = make([]int, len(line)-1)
	}

	sum := 0
	for i := range buffer {
		sum += buffer[i]
	}
	println(sum)
}
