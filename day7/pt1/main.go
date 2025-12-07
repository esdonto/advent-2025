package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	line, err := reader.ReadString('\n')

	buffer := []rune(strings.ReplaceAll(line[:len(line)-1], "S", "|"))
	newBuffer := slices.Repeat([]rune{'.'}, len(buffer))
	splitCount := 0

	for ; err == nil; line, err = reader.ReadString('\n') {
		for i := range buffer {
			if buffer[i] == '|' {
				if line[i] == '^' {
					splitCount++
					if i > 0 {
						newBuffer[i-1] = '|'
					}
					if i < len(buffer)-1 {
						newBuffer[i+1] = '|'
					}
				} else {
					newBuffer[i] = '|'
				}
			}
		}
		buffer = newBuffer
		newBuffer = slices.Repeat([]rune{'.'}, len(buffer))
	}
	println(splitCount)
}
