package main

import "fmt"

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func (tree *Tree) addNode(data int) *Tree {
	if data <= tree.data {
		if tree.left == nil {
			tree.left = &Tree{data: data}
		} else {
			tree.left.addNode(data)
		}
	} else {
		if tree.right == nil {
			tree.right = &Tree{data: data}
		} else {
			tree.right.addNode(data)
		}
	}
	return tree
}

func (tree *Tree) inOrder() {
	temp := tree
	if temp == nil {
		return
	} else {
		if temp.left != nil {
			temp.left.inOrder()
		}
		fmt.Println(temp.data)
		if temp.right != nil {
			temp.right.inOrder()
		}
	}
}
func main() {
	root := &Tree{data: 10}
	root = root.addNode(5)
	root = root.addNode(2)
	root = root.addNode(100)
	root.inOrder()
}
