package main

import (
	"fmt"
	"learngo/oop/treeNode"
)

type myTreeNode struct {
	node *treeNode.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root treeNode.Node
	root = treeNode.Node{Value: 3}
	root.Left = &treeNode.Node{}
	root.Right = &treeNode.Node{Value: 5}
	root.Right.Left = new(treeNode.Node)
	root.Left.Right = &treeNode.Node{Value: 2}
	root.Traverse()

	node := myTreeNode{&root}
	node.postOrder()

	root.Print()
	(&root).SetValue(0)
	root.Print()

	pnode := root
	pnode.SetValue(1)
	pnode.Print()

	var proot treeNode.Node
	proot.SetValue(100)
	proot.Print()

	var proot1 *treeNode.Node
	proot1.SetValue(100)
	proot1.Print()

	nodes := []treeNode.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
