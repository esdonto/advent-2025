package main

import (
	"bufio"
	"fmt"
	"math"
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

	sum := 0

	reader := bufio.NewReader(file)
	scope, _ := reader.ReadString(',')
	for ; len(scope)>0; scope, _ = reader.ReadString(','){
		scope_separated := strings.Split(scope[:len(scope)-1], "-")

		start, err := strconv.Atoi(scope_separated[0])
		if err != nil {panic(err)}

		stop, err :=  strconv.Atoi(scope_separated[1])
		if err != nil {panic(err)}

		println(start, "-", stop)

		for i := start; i <= stop; i++ {
			decimalCases := int(math.Ceil(math.Log10(float64(i+1))))

			for subdivisionSize := 1 ; subdivisionSize <= decimalCases / 2; subdivisionSize++ {
				// If the subdivision repetiton fits into the ID
				if decimalCases % subdivisionSize != 0 {continue}

				subdivision := i % int(math.Pow10(subdivisionSize))

				// If the subdivision has any leading zeros
				if subdivision < int(math.Pow10(subdivisionSize - 1)) {continue}

				if i == createInvalidID(decimalCases, subdivisionSize, subdivision) {
					sum += i
					println(" ", i)
					break
				}
			}
		}
	}
	println("sum:", sum)
}

// Creating what the invalid ID would be given the repeating subdivision and ID size
func createInvalidID(sizeID int, sizeSubdivision int, subdivision int) int {
	invalidID := subdivision

	for i := 1; i <= sizeID - sizeSubdivision; i++ {
		invalidID *= 10
		if i % sizeSubdivision == 0 {
			invalidID += subdivision
		}
		if subdivision == 10 {
		}
	}

	return invalidID
}
