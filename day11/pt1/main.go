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
		nodeOutput := make([]*Node, len(connections[nodeName]))
		for i, connectedNode := range connections[nodeName] {
			nodeOutput[i] = nodes[connectedNode]
		}
		nodes[nodeName].output = nodeOutput
	}

	outPaths := 0
	pathsFIFO := [][]*Node{{nodes["you"]}}

	for len(pathsFIFO) > 0 {
		currentPath := pathsFIFO[len(pathsFIFO)-1]
		pathsFIFO = pathsFIFO[:len(pathsFIFO)-1]

		for _, nextNode := range currentPath[len(currentPath)-1].output {
			if nextNode.name == "out" {
				outPaths++
			} else {
				alreadyInPath := slices.Contains(currentPath, nextNode)
				if !alreadyInPath {
					newPath := make([]*Node, len(currentPath))
					copy(newPath, currentPath)
					newPath = append(newPath, nextNode)
					pathsFIFO = append(pathsFIFO, newPath)
				}
			}
		}
	}

	fmt.Println(outPaths)

}

type Node struct {
	name   string
	output []*Node
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
