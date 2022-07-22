package test

//遍历二叉树
import (
	"fmt"
	"testing"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//为 TraverseFunc 方法提供 实现
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

//为 Node 结构增加一个方法 TraverseFunc ，
//此方法 传入一个方法参数，在遍历是执行
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func TestTraverse(t *testing.T) {
	var root Node
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse() // 进行了 打印封装

	//以下通过匿名函数，实现了 自定义实现
	nodeCount := 0
	root.TraverseFunc(func(node *Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount) //Node count: 5
}
