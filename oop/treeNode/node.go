package treeNode

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//局部变量也可以给外面用  返回局部变量的地址
//创建在堆上还是栈上 根据编译器选择
//在这个函数中是放在堆中
//在堆中参与垃圾回收
//在栈中 函数退出后局部变量随之消失
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		println("node is nil")
		return
	}
	node.Value = value
}
