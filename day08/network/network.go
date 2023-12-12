package network

import (
	"bufio"
	"io"
	"strings"
)

type Node struct {
	Value string
	Left  string
	Right string
}

type Network struct {
	instruction string
	nodes       map[string]Node
}

func Build(r io.Reader) *Network {
	network := &Network{nodes: make(map[string]Node)}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	network.instruction = strings.TrimSpace(scanner.Text())
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		node := Node{
			Value: line[0:3],
			Left:  line[7:10],
			Right: line[12:15],
		}
		network.nodes[node.Value] = node
	}

	return network
}

func (n *Network) Navigate1() int {
	return n.walk(n.nodes["AAA"], n.instruction, 0)
}

func (n *Network) walk(node Node, instruction string, step int) int {
	if node.Value == "ZZZ" {
		return 0
	}

	var next Node
	if instruction[step%len(instruction)] == 'R' {
		next = n.nodes[node.Right]
	} else {
		next = n.nodes[node.Left]
	}
	return 1 + n.walk(next, instruction, step+1)
}

func (n *Network) Navigate2() int {
	var count []int
	for _, node := range n.nodes {
		if node.Value[2] != 'A' {
			continue
		}

		steps := 0
		i := 0
		for node.Value[2] != 'Z' {
			if n.instruction[i] == 'R' {
				node = n.nodes[node.Right]
			} else {
				node = n.nodes[node.Left]
			}

			steps++
			i = (i + 1) % len(n.instruction)
		}

		count = append(count, steps)
	}

	result := count[0]
	for i := 1; i < len(count); i++ {
		result *= count[i] / gcd(result, count[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
