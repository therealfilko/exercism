package linkedlist

import "errors"

// Define the List and Element types here.
type List struct {
    Head *Node
}

type Node struct {
    Value int
    Next *Node
}

func New(elements []int) *List {
    list := &List{}
    for i := len(elements) - 1; i >= 0; i-- {
        newNode := &Node{Value: elements[i], Next: list.Head}
        list.Head = newNode
    }
    return list
}

func (l *List) Size() int {
    sum := 0
    cur := l.Head
    for cur != nil {
        sum++
        cur = cur.Next
    }
    return sum
}

func (l *List) Push(element int) {
    newNode := &Node{Value: element, Next: l.Head}
    l.Head = newNode
}

func (l *List) Pop() (int, error) {
    if l.Head == nil {
        return -1, errors.New("Empty List")
    } 
    cur := l.Head.Value
    l.Head = l.Head.Next

    return cur, nil
}

func (l *List) Array() []int {
    var arr []int
    cur := l.Head
    for cur != nil {
        arr = append(arr, cur.Value)
        cur = cur.Next
    }
    return arr
}

func (l *List) Reverse() *List {
    cur := l.Head
    var prev *Node = nil
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return l
}
