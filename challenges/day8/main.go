package day8

import (
	"fmt"
	"strings"

	"github.com/daniel-moya/aoc/challenges/utils"
)

func Run(path string) int {
	return Part2(path)
}

func Part1(path string) int {
	lines := utils.GetInputFromFile(path)

	directions := strings.Trim(lines[0], " ")
	nodes := lines[2:]

	nodeMap := MapNodes(nodes)

	next := "AAA"
	i := 0
	steps := 0
	for next != "ZZZ" {
		node := nodeMap[next]
		next = node.getNext(directions[i])

		if i == len(directions)-1 {
			i = 0
		} else {
			i += 1
		}
		steps++
	}

	return steps
}

func Part2(path string) int {
	lines := utils.GetInputFromFile(path)

	directions := strings.Trim(lines[0], " ")
	nodes := lines[2:]

	nodeMap := MapNodes(nodes)

	nextPoints := []string{}

	for key := range nodeMap {
		if key[2] == 'A' {
			nextPoints = append(nextPoints, key)
		}
	}
	steps := 0
	fmt.Println(nextPoints)
	i := 0
	for !nextPointsHaveZ(nextPoints) {
		u := []string{}
		for _, np := range nextPoints {
			node := nodeMap[np]
			u = append(u, node.getNext(directions[i]))
		}

		nextPoints = u

		if i == len(directions)-1 {
			i = 0
		} else {
			i += 1
		}
		steps++
	}

	return steps
}

func nextPointsHaveZ(nextPoints []string) bool {
	for _, np := range nextPoints {
		if np[2] != 'Z' {
			return false
		}
	}
	return true
}

func MapNodes(lines []string) map[string]Node {
	nodeMap := make(map[string]Node)
	for _, line := range lines {
		node := NewNode(line)
		nodeMap[node.Key] = node
	}
	return nodeMap
}

type Node struct {
	Key   string
	Right string
	Left  string
}

func NewNode(line string) Node {
	node := &Node{}
	node.ParseRaw(line)
	return *node
}

func (n *Node) ParseRaw(line string) {
	bits := strings.Split(line, " = ")
	key := bits[0]
	directions := strings.Split(strings.Trim(bits[1], "()"), ", ")
	left := directions[0]
	right := directions[1]

	n.Key = key
	n.Left = left
	n.Right = right
}

func (n *Node) getNext(direction byte) string {
	if string(direction) == "R" {
		return strings.Trim(n.Right, " ")
	}
	return strings.Trim(n.Left, " ")
}
