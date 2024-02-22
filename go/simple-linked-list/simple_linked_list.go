package linkedlist

import "errors"

// Define the List and Element types here.
type List struct {
    Head *Node
    size int
}

type Node struct {
    Value int
    Next *Node
}

func New(elements []int) *List {
    list := &List{}
    for i := range elements {
        list.Push(elements[i])
    }
    return list
}

func (l *List) Size() int {
    return l.size
}

func (l *List) Push(element int) {
    l.Head = &Node{Value: element, Next: l.Head}
    l.size++
}

func (l *List) Pop() (int, error) {
    if l.size < 1 {
        return 0, errors.New("Empty List")
    } 
    
    head, data := l.Head, l.Head.Value
    l.Head = head.Next
    head.Next = nil
    l.size--
    return data, nil
}

func (l *List) Array() []int {
    arr := make([]int, l.size)
    for i, head := l.size - 1, l.Head; head != nil; i, head = i - 1, head.Next {
        arr[i] = head.Value
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
    l.Head = prev
    return l
}
