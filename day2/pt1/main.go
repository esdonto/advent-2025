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

			// If it has an odd number of digits i can skip to when it has an even amount
			if decimalCases % 2 == 1 {
				i = int(math.Pow10(decimalCases)) - 1
				continue
			}

			halfPow := int(math.Pow10(decimalCases/2))
			if i % halfPow == i / halfPow {
				sum += i
				println(" ", i)
			}
		}
	}
	println("sum:", sum)
}
