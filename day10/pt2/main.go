package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

// Based on the "Bifurcate your way to victory" by u/tenthmascot solution on https://reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/

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

		found, minButtons := evaluateAllCombinations(joltage, buttons)

		if !found {
			fmt.Println("No possible button presses found!")
			panic("")
		}

		sumButtons += minButtons
	}

	println(sumButtons)
}

func iterateCombinations(nButtons int) iter.Seq[[]bool] {
	pressed := make([]bool, nButtons)
	return func(yield func([]bool) bool) {
		if !yield(pressed) {
			return
		}
		for binaryAddOne(nButtons, pressed) {
			if !yield(pressed) {
				return
			}
		}
	}
}

// Returns if it has found a possible combination and the amount of presses of the minimal one
func evaluateAllCombinations(joltageTarget []int, buttons [][]int) (bool, int) {
	//fmt.Println(joltageTarget)
	if allZeroes(joltageTarget) {
		return true, 0
	}

	firstFound := false
	var minPresses int

	for pressed := range iterateCombinations(len(buttons)) {
		isValid, newJoltageTarget := evaluateCombination(joltageTarget, buttons, pressed)
		if isValid {
			found, nPresses := evaluateAllCombinations(divideJoltages(newJoltageTarget), buttons)
			if found {
				totalPresses := 2*nPresses + sumPresses(pressed)
				if !firstFound || minPresses > totalPresses {
					minPresses = totalPresses
					firstFound = true
				}
			}
		}
	}

	if !firstFound {
		return false, 0
	}

	return true, minPresses
}

// Returns if the result of the button combination subtratcion results in a valid joltage read (i.e. all evens and >0) and the result of the subtraction
func evaluateCombination(joltageTarget []int, buttons [][]int, buttonsPressed []bool) (bool, []int) {
	joltages := make([]int, len(joltageTarget))
	copy(joltages, joltageTarget)

	for i := range buttonsPressed {
		if buttonsPressed[i] {
			for _, v := range buttons[i] {
				joltages[v] -= 1
			}
		}
	}

	for _, v := range joltages {
		if v < 0 || v%2 > 0 {
			return false, []int{}
		}
	}

	return true, joltages
}

func parseLights(lightsString string) []bool {
	lightsBool := make([]bool, len(lightsString)-2)

	for i := range lightsBool {
		if lightsString[i+1] == '#' {
			lightsBool[i] = true
		}
	}

	return lightsBool
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

// Returns false if it has reached the end of the line
func binaryAddOne(nButtons int, buttons []bool) bool {
	for i := range nButtons {
		buttons[i] = !buttons[i]
		if buttons[i] {
			return true
		}
	}
	return false
}

func sumPresses(buttons []bool) int {
	sum := 0
	for _, v := range buttons {
		if v {
			sum++
		}
	}
	return sum
}
func allZeroes(joltages []int) bool {
	for _, v := range joltages {
		if v != 0 {
			return false
		}
	}
	return true
}
func divideJoltages(joltages []int) []int {
	for i := range joltages {
		joltages[i] /= 2
	}
	return joltages
}
