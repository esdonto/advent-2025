package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// THIS ISNT A GRAPH, ITS A TREE!!!

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error!")
		panic(err)
	}
	defer file.Close()

	nodes := map[string]*Node{}
	connections := map[string][]string{}

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')

	for ; err == nil; line, err = reader.ReadString('\n') {
		lineSeparated := strings.Split(line[:len(line)-1], ": ")
		nodeName := lineSeparated[0]
		nodes[nodeName] = &Node{name: nodeName}
		connections[nodeName] = strings.Split(lineSeparated[1], " ")
	}
	nodes["out"] = &Node{name: "out"}

	for nodeName := range connections {
		for _, connectedNode := range connections[nodeName] {
			nodes[nodeName].output = append(nodes[nodeName].output, nodes[connectedNode])
			nodes[connectedNode].input = append(nodes[connectedNode].input, nodes[nodeName])
		}
	}

	checkReachability(nodes["fft"], nodes["dac"])

	svt_to_fft := 0
	pathsFIFO := [][]*Node{{nodes["svr"]}}

	for len(pathsFIFO) > 0 {
		currentPath := pathsFIFO[len(pathsFIFO)-1]
		pathsFIFO = pathsFIFO[:len(pathsFIFO)-1]

		//println(len(pathsFIFO), len(currentPath))
		//printPath(currentPath)

		for _, nextNode := range currentPath[len(currentPath)-1].output {
			if nextNode.name == "fft" || nextNode.name == "dac" || nextNode.name == "out" {
				svt_to_fft++
				//println(svt_to_fft)
			} else {
				alreadyInPath := slices.Contains(currentPath, nextNode)
				if alreadyInPath {
					println("oi")
				}
				if !alreadyInPath {
					newPath := make([]*Node, len(currentPath))
					copy(newPath, currentPath)
					newPath = append(newPath, nextNode)
					pathsFIFO = append(pathsFIFO, newPath)
				}
			}
		}
	}
}

func checkReachability(fft, dac *Node) {
	FIFO := []*Node{fft}
	for len(FIFO) > 0 {
		currentNode := FIFO[len(FIFO)-1]
		FIFO = FIFO[:len(FIFO)-1]

		currentNode.fromFFT = true

		for _, nextNode := range currentNode.output {
			if !nextNode.fromFFT {
				FIFO = append(FIFO, nextNode)
			}
		}
	}

	FIFO = []*Node{fft}
	for len(FIFO) > 0 {
		currentNode := FIFO[len(FIFO)-1]
		FIFO = FIFO[:len(FIFO)-1]

		currentNode.toFFT = true

		for _, nextNode := range currentNode.input {
			if !nextNode.toFFT {
				FIFO = append(FIFO, nextNode)
			}
		}
	}

	FIFO = []*Node{dac}
	for len(FIFO) > 0 {
		currentNode := FIFO[len(FIFO)-1]
		FIFO = FIFO[:len(FIFO)-1]

		currentNode.fromDAC = true

		for _, nextNode := range currentNode.output {
			if !nextNode.fromDAC {
				FIFO = append(FIFO, nextNode)
			}
		}
	}

	FIFO = []*Node{dac}
	for len(FIFO) > 0 {
		currentNode := FIFO[len(FIFO)-1]
		FIFO = FIFO[:len(FIFO)-1]

		currentNode.toDAC = true

		for _, nextNode := range currentNode.input {
			if !nextNode.toDAC {
				FIFO = append(FIFO, nextNode)
			}
		}
	}
}

type Node struct {
	name    string
	output  []*Node
	input   []*Node
	fromFFT bool //FFT output reaches it
	toFFT   bool //Its output can reachs FFT
	fromDAC bool //DAC output reaches it
	toDAC   bool //Its output can reachs DAC
}

func printPath(path []*Node) {
	print("path: [ ")
	for _, node := range path {
		print(node.name, " ")
	}
	println("]")
}

func printFIFO(fifo [][]*Node) {
	print("fifo: [ ")
	for _, path := range fifo {
		print("[ ")
		for _, node := range path {
			print(node.name, " ")
		}
		print("] ")
	}
	println("]")
}
