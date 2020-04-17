package main

import (
	"fmt"
)

func addressOf(p, a interface{}) {
	fmt.Printf("address of %s: %v\n", p, a)
}
func valOf(p, a interface{}) {
	fmt.Printf("val of %s: % \n", p, a)
}

func main() {
	x := 4
	p := &x
	addressOf("x", &x)
	addressOf("p", &p)
	// fmt.Println(&p)
	valOf("*p", *p)
	// test(p)
	// fmt.Println(p)
	// fmt.Println(*p)
}

func test(p *int) {
	fmt.Println(p)
	d := 5
	fmt.Println("&d: ", &d)
	*p = *(&d)
	fmt.Println(p)
	d = 7
	fmt.Println("&d: ", &d)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(val int) {
	if l.Next != nil {
		l.Next.add(val)
	} else {
		l.Next = &ListNode{Val: val}
	}
}

func (l *ListNode) print() {
	fmt.Printf("%d ", l.Val)
	if l.Next != nil {
		l.Next.print()
	} else {
		fmt.Println()
	}
}

// // (2 -> 4 -> 3) + (5 -> 6 -> 4)
// // 342 + 465 = 807.
// // 7 0 8
// func main() {

// 	// acc.print()

// 	l1 := &ListNode{Val: 5}
// 	// l1 := &ListNode{Val: 2}
// 	l1.add(4)
// 	// l1.add(3)
// 	// // l1.print()
// 	l2 := &ListNode{Val: 5}
// 	// l2 := &ListNode{Val: 5}
// 	// l2.add(6)
// 	// l2.add(4)

// 	// l2.print()

// 	addTwoNumbers(l1, l2).print()
// 	// fmt.Println(l1)
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add(l1, l2, 0)
	return l2
}

func add(l1, l2 *ListNode, carried int) {
	res := l1.Val + l2.Val + carried
	carried = 0
	if res >= 10 {
		res -= 10
		carried++
	}
	l2.Val = res
	switch {
	case l1.Next != nil && l2.Next != nil:
		add(l1.Next, l2.Next, carried)
	case l1.Next == nil && l2.Next != nil:
		l1.Next = &ListNode{}
		add(l1.Next, l2.Next, carried)
	case l1.Next != nil && l2.Next == nil:
		l2.Next = &ListNode{}
		add(l1.Next, l2.Next, carried)
	case carried > 0:
		l2.Next = &ListNode{Val: carried}
	}
}
