package main

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

var (
	nodeG Node = Node{data: "g", left: nil, right: nil}
	nodeF Node = Node{data: "f", left: &nodeG, right: nil}
	nodeE Node = Node{data: "e", left: nil, right: nil}
	nodeD Node = Node{data: "d", left: &nodeE, right: nil}
	nodeC Node = Node{data: "c", left: nil, right: nil}
	nodeB Node = Node{data: "b", left: &nodeD, right: &nodeF}
	nodeA Node = Node{data: "a", left: &nodeB, right: &nodeC}
)

/**
广度优先: a -> b -> c -> d -> f -> e -> g

先序遍历: a -> b -> d -> e -> f -> g -> c

中序遍历: e -> d -> b -> g -> f -> a -> c

后序遍历: e -> d -> g -> f -> b -> c -> a
*/

//广度优先
func breadthFirstSearch(node Node) []string {
	var result []string
	var nodes []Node = []Node{node}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		result = append(result, node.data)
		if node.left != nil {
			nodes = append(nodes, *node.left)
		}
		if node.right != nil {
			nodes = append(nodes, *node.right)
		}
	}
	return result
}
func preOrderRecursive(node Node) {
	// 在这里输出就是先序
	fmt.Println(node.data)
	if node.left != nil {
		preOrderRecursive(*node.left)
	}
	// 在这里输出就是中序
	fmt.Println(node.data)
	if node.right != nil {
		preOrderRecursive(*node.right)
	}
	// 在这里输出是后序
	fmt.Println(node.data)
}
func main() {
	preOrderRecursive(nodeA)
	//breadthFirstSearch(nodeA)
}
