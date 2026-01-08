package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// THERES NO LOOPBACK!!!

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
		}
	}

	svrTOfft := findPaths(nodes["svr"], nodes["fft"], map[*Node]int{})
	svrTOdac := findPaths(nodes["svr"], nodes["dac"], map[*Node]int{})
	fftTOdac := findPaths(nodes["fft"], nodes["dac"], map[*Node]int{})
	dacTOfft := findPaths(nodes["dac"], nodes["fft"], map[*Node]int{})
	fftTOout := findPaths(nodes["fft"], nodes["out"], map[*Node]int{})
	dacTOout := findPaths(nodes["dac"], nodes["out"], map[*Node]int{})

	println((svrTOfft * fftTOdac * dacTOout) + (svrTOdac * dacTOfft * fftTOout))
}

func findPaths(node *Node, target *Node, cache map[*Node]int) int {
	if node == target {
		return 1
	}
	if value, ok := cache[node]; ok {
		return value
	}

	sum := 0
	for _, nextNode := range node.output {
		sum += findPaths(nextNode, target, cache)
	}
	cache[node] = sum
	return sum
}

type Node struct {
	name   string
	output []*Node
}
