package linkedlist

import "fmt"
import "sync"

type Node struct {
	Val        interface{}
	next, prev *Node
}
type List struct {
	head *Node
	tail *Node
	len  int
}

var ErrEmptyList error

func (e *Node) Next() *Node {
	return e.next
}
func (e *Node) Prev() *Node {
	return e.prev
}
func NewList(args ...interface{}) *List {
	newL := &List{}
	if len(args) != 0 {
		for _, val := range args {
			newL.PushBack(val)
		}
	}
	return newL
}
func (l *List) PushFront(v interface{}) {
	newNode := &Node{Val: v, next: l.head, prev: nil}
	l.head = newNode
	if l.len == 0 {
		l.tail = l.head
		l.len++
		return
	}
	l.head.next.prev = l.head
	l.len++
	return
}
func (l *List) PushBack(v interface{}) {
	newNode := &Node{Val: v, next: nil, prev: l.tail}
	l.tail = newNode
	if l.len == 0 {
		l.head = newNode
		l.len++
		return
	}
	l.tail.prev.next = l.tail
	l.len++
	return
}
func (l *List) PopFront() (interface{}, error) {
	if l.len == 0 {
		return nil, ErrEmptyList
	}
	d := l.head.Val
	if l.head.next != nil {
		l.head.next.prev = nil
	}
	l.head = l.head.next
	if l.len == 1 {
		l.tail = l.head
	}
	l.len--
	return d, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.len == 0 {
		return nil, ErrEmptyList
	}
	v := l.tail.Val
	if l.tail.prev != nil {
		l.tail.prev.next = nil
	}
	l.tail = l.tail.prev
	if l.len == 1 {
		l.head = l.tail
	}
	l.len--
	return v, nil
}
func (l *List) Reverse() *List {
	pointerH := l.head
	pointerT := l.tail
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(n *Node) {
		for ll := 0; ll < l.len/2; ll++ {
			swap(n)
			n = n.prev
		}
		wg.Done()
	}(pointerH)
	go func(n *Node) {
		for ll := l.len; ll > l.len/2; ll-- {
			swap(n)
			n = n.next
		}
		wg.Done()
	}(pointerT)
	wg.Wait()
	l.head = pointerT
	l.tail = pointerH
	return l
}
func swap(n *Node) *Node {
	tmp := n.next
	n.next = n.prev
	n.prev = tmp
	return n
}
func (l *List) First() *Node {
	return l.head
}
func (l *List) Last() *Node {
	return l.tail
}
func (l *List) Show() {
	for list := l.head; list != nil; list = list.Next() {
		fmt.Printf("%+v", list)
	}

}
