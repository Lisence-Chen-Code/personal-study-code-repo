package test

import (
	"fmt"
	"testing"
)

//链表节点
type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

// 循环链表
type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) initRing() *Ring {
	r.next = r
	r.prev = r
	return r
}

/*
模拟单链表结构，依照自定义指针寻址去打印链表中各个节点的data
*/
func TestLinkStructure(t *testing.T) {
	// 新的节点
	node := new(LinkNode)
	node.Data = 2

	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1 // node1 链接到 node 节点上

	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2 // node2 链接到 node1 节点上

	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			fmt.Println(nowNode.Data)
			// 获取下一个节点
			nowNode = nowNode.NextNode
		} else {
			// 如果下一个节点为空，表示链表结束了
			break
		}
	}
}

func TestRing(t *testing.T) {
	r := new(Ring)
	r.initRing()
	NewRing(10)
}

// 创建N个节点的循环链表
func NewRing(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
