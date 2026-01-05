package main

import (
	"bufio"
	"fmt"
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

	sumButtons := 0

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], " ")
		//lights := parseLights(lineSeparated[0])
		buttons := parseButtons(lineSeparated[1 : len(lineSeparated)-1])
		joltage := parseJoltage(lineSeparated[len(lineSeparated)-1])

		//sumButtons += getMinButtonPresses(buttons, joltage)

		print(getMinButtonPresses(buttons, joltage))
	}

	println(sumButtons)
}

func getMinButtonPresses(buttons [][]int, joltageTarget []int) int {
	bufferPresses := make([]int, len(buttons))

	var minPresses int
	//foundFirst := false
	increaseIdx := 0
	for {
		//fmt.Println(bufferPresses)
		exactMatch, surpassesTarget := evaluateCombination(joltageTarget, buttons, bufferPresses)
		if exactMatch {
			nPresses := sumPresses(bufferPresses)
			println(nPresses)
		}
		// Kind of a base-m +1 algorithm
		if exactMatch || surpassesTarget {
			if increaseIdx == len(bufferPresses)-1 {
				// End of the line, all possible combinations have been checked
				break
			}

			bufferPresses[increaseIdx] = 0
			increaseIdx++
		} else {
			increaseIdx = 0
		}
		bufferPresses[increaseIdx]++
		/*if exactMatch {
			nPresses := sumPresses(bufferPresses)
			if !foundFirst || minPresses < nPresses {
				minPresses = nPresses
			}
		} else {

		}*/
	}
	return minPresses
}

func sumPresses(buttonsPressed []int) int {
	sum := 0
	for _, v := range buttonsPressed {
		sum += v
	}
	return sum
}

// 1st bool evaluates if the combination gets the exact target
// 2nd bool evaluates if the combination meets the target
func evaluateCombination(joltageTarget []int, buttons [][]int, buttonsPressed []int) (bool, bool) {
	joltages := make([]int, len(joltageTarget))

	for i := range buttonsPressed {
		for _, v := range buttons[i] {
			joltages[v] += buttonsPressed[i]
		}
	}

	exactMatch := true

	for i := range joltages {
		if joltages[i] != joltageTarget[i] {
			exactMatch = false
			if joltages[i] > joltageTarget[i] {
				return false, true
			}
		}
	}

	return exactMatch, false
}

func parseButtons(buttonsString []string) [][]int {
	buttonsInt := make([][]int, len(buttonsString))

	for i := range buttonsInt {
		button := strings.Split(buttonsString[i][1:len(buttonsString[i])-1], ",")

		buttonsInt[i] = make([]int, len(button))
		for j := range buttonsInt[i] {
			v, err := strconv.Atoi(button[j])
			if err != nil {
				panic("Error reading button value")
			}
			buttonsInt[i][j] = v
		}
	}

	return buttonsInt
}

func parseJoltage(joltageString string) []int {
	joltageStringSeparated := strings.Split(joltageString[1:len(joltageString)-1], ",")
	joltageInt := make([]int, len(joltageStringSeparated))

	for i := range joltageInt {
		v, err := strconv.Atoi(joltageStringSeparated[i])
		if err != nil {
			panic("Error reading joltage value")
		}
		joltageInt[i] = v
	}

	return joltageInt
}
