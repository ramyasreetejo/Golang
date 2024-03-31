package main

import "fmt"

func main() {
	fmt.Println("Implementing bst in go.")
	newBst := bst{nil}
	newBst.insert(6)
	newBst.insert(4)
	newBst.insert(5)
	newBst.insert(9)
	newBst.insert(7)
	newBst.insert(10)
	newBst.insert(2)
	newBst.levelordertraversal()
	newBst.delete(9)
	newBst.levelordertraversal()
	newBst.delete(6)
	newBst.levelordertraversal()
	newBst.delete(5)
	newBst.levelordertraversal()
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bst struct {
	root *Node
}

func (bst *bst) insert(data int) {
	newNode := &Node{data, nil, nil}
	if bst.root == nil {
		bst.root = newNode
		return
	}
	if data < bst.root.data {
		bst.root.left = bst.insertNode(bst.root.left, data)
	} else {
		bst.root.right = bst.insertNode(bst.root.right, data)
	}
}

func (bst *bst) insertNode(root *Node, data int) *Node {
	newNode := &Node{data, nil, nil}
	if root == nil {
		return newNode
	}
	if data < root.data {
		root.left = bst.insertNode(root.left, data)
	} else {
		root.right = bst.insertNode(root.right, data)
	}
	return root
}

func (bst *bst) delete(data int) {
	bst.root = bst.deleteNode(bst.root, data)
}

func (bst *bst) deleteNode(root *Node, data int) *Node {
	if root == nil {
		return nil
	}
	if root.data == data {
		if root.right == nil {
			root = root.left
		} else if root.left == nil {
			root = root.right
		} else {
			inordersucc := bst.inorder_succ(root.right)
			root.data = inordersucc.data
			root.right = bst.deleteNode(root.right, inordersucc.data)
		}
	} else {
		if data < root.data {
			root.left = bst.deleteNode(root.left, data)
		} else {
			root.right = bst.deleteNode(root.right, data)
		}
	}
	return root
}

func (bst *bst) inorder_succ(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func (bst *bst) levelordertraversal() {
	if bst.root == nil {
		return
	}
	queue := []*Node{}
	queue = append(queue, bst.root)
	for len(queue) != 0 {
		temp := queue[0]
		fmt.Print(temp.data, " ")
		queue = queue[1:]
		if temp.left != nil {
			queue = append(queue, temp.left)
		}
		if temp.right != nil {
			queue = append(queue, temp.right)
		}
	}
	fmt.Println()
}
