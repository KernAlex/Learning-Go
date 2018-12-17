package testProj

import (
	"fmt"
	"strings"
)

type Node struct{
	value interface{}
	next *Node
}
type LinkList struct{
	front *Node
}

func MakeNew() LinkList{
	return *new(LinkList)
}

func (l *LinkList) InsertFront(x interface{}) {
	temp := Node{}
	temp.value = x
	temp.next = l.front
	l.front = &temp
}

func (l LinkList) String() string{
	temp := l.front
	var s []interface{}
	for temp != nil {
		s = append(s, temp.value)
		temp = temp.next
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), "->"), "[]")
}

