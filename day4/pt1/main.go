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

	//sum := 0

	// Assembiling the matrix

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')

	werehouse := [][]bool{}
	werehouse = append(werehouse, make([]bool, len(line)+1))

	for ; err == nil; line, err = reader.ReadString('\n') {
		row := make([]bool, len(line)+1)

		for i := range len(row) - 1 {
			if line[i] == '@' {
				row[i+1] = true
			}
		}

		werehouse = append(werehouse, row)
	}

	werehouse = append(werehouse, make([]bool, len(werehouse[0])))

	// Checking the matrix

	count := 0

	for i := 1; i < len(werehouse)-1; i++ {
		for j := 1; j < len(werehouse[i])-1; j++ {
			if werehouse[i][j] {
				if count3x3(werehouse, i, j) < 5 {
					//print("X")
					count++
				} else {
					//print("@")
				}
			} else {
				//print(".")
			}
		}
		//println()
	}

	println("count:", count)
}

func count3x3(werehouse [][]bool, row int, column int) int {
	count := 0

	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if werehouse[i][j] {
				count++
			}
		}
	}

	return count
}
