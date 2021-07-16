package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	node1 := &Node{Val: 5}
	node2 := &Node{Val: 3}
	node3 := &Node{Val: 7}
	node4 := &Node{Val: 2}
	node5 := &Node{Val: 6}
	node6 := &Node{Val: 8}
	node7 := &Node{Val: 1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node2
	fmt.Println(IsCycle(node1))
	fmt.Println(CycleNum(node1))
	fmt.Println(firstMeetNode(node1))
}

// 快慢指针法 时间复杂度O(n) 空间复杂度是 O(1)
// IsCycle 判断链表是否有环
func IsCycle(head *Node) bool {
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1.Val == p2.Val {
			return true
		}
	}
	return false
}

// CycleNum 链表环长
func CycleNum(head *Node) int {
	p1 := head
	p2 := head
	n := 0
	meet := false
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if meet {
			n++
		}
		if p1.Val == p2.Val {
			if meet {
				return n
			}
			meet = true
		}
	}
	return n
}

// firstMeetNode 判断入环点
func firstMeetNode(head *Node) *Node {
	p1 := head
	p2 := head
	var firstMeetNode *Node
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1.Val == p2.Val {
			firstMeetNode = p1
			break
		}
	}
	p1 = head
	for firstMeetNode != nil &&  firstMeetNode.Next != nil {
		p1 = p1.Next
		firstMeetNode = firstMeetNode.Next
		if p1.Val == firstMeetNode.Val {
			return p1
		}
	}
	return nil
}

