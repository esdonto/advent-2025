package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		lineSeparated := strings.Split(line, " ")
		lights := parseLights(lineSeparated[0])
		buttons := parseButtons(lineSeparated[1 : len(lineSeparated)-1])
		//joltage := lineSeparated[len(lineSeparated)-1]

		for i := range len(buttons) + 1 {
			if evaluateAllCombinations(lights, buttons, i) {
				sumButtons += i
				break
			}
		}
	}

	println(sumButtons)
}

func evaluateAllCombinations(lightsTarget []bool, buttons [][]int, nPressed int) bool {
	buttonCombinations := getAllButtonCombinations(len(buttons), nPressed)
	for _, pressed := range buttonCombinations {
		if evaluateCombination(lightsTarget, buttons, pressed) {
			return true
		}
	}

	return false
}

func evaluateCombination(lightsTarget []bool, buttons [][]int, buttonsPressed []bool) bool {
	lights := make([]bool, len(lightsTarget))

	for i := range buttonsPressed {
		if buttonsPressed[i] {
			for _, v := range buttons[i] {
				lights[v] = !lights[v]
			}
		}
	}

	for i := range lights {
		if lights[i] != lightsTarget[i] {
			return false
		}
	}

	return true
}

func getAllButtonCombinations(nButtons, nPressed int) [][]bool {
	if nPressed == 0 {
		return [][]bool{make([]bool, nButtons)}
	} else if nPressed == nButtons {
		return [][]bool{slices.Repeat([]bool{true}, nButtons)}
	} else {
		// Append false
		combFalse := getAllButtonCombinations(nButtons-1, nPressed)
		for i := range combFalse {
			combFalse[i] = append(combFalse[i], false)
		}

		// Append true
		combTrue := getAllButtonCombinations(nButtons-1, nPressed-1)
		for i := range combTrue {
			combTrue[i] = append(combTrue[i], true)
		}

		return append(combFalse, combTrue...)
	}
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
