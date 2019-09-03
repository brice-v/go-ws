package main

import (
	"container/list"
	"fmt"
)

type NodeType int

const (
	Added NodeType = iota
	Original
)

type PieceTable struct {
	original []byte
	added    []byte
	nodes    *list.List
}

func (n *Node) String() (result string) {
	if n.typ == Original {
		result = fmt.Sprintf("{NodeType: Original, start: %d, length: %d}", n.start, n.length)
	} else {
		result = fmt.Sprintf("{NodeType: Added, start: %d, length: %d}", n.start, n.length)
	}
	return
}

type Node struct {
	typ    NodeType
	start  int
	length int
}

func (PT *PieceTable) newNode(typ NodeType, start, length int) {
	PT.nodes.PushBack(&Node{typ: typ, start: start, length: length})
}

func (PT *PieceTable) Insert(data []byte) {
	dataLen := len(data)
	dataStart := len(PT.added)

	PT.added = append(PT.added, data...)
	PT.newNode(Added, dataStart, dataLen)
}

func (PT *PieceTable) Display() {
	for e := PT.nodes.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		n := e.Value.(*Node)
		if n.typ == Original {
			for i:= n.start; i< n.start+n.length; i++ {
			fmt.Print(string(PT.original[i]))
			}
		} else {
			for i:= n.start; i< n.start+n.length; i++ {
			fmt.Print(string(PT.added[i]))
			}
		}

	}
}

func NewPT(optBuf []byte) *PieceTable {
	if optBuf != nil {
		pt := &PieceTable{original: optBuf, added: []byte(""), nodes: list.New()}
		pt.newNode(Original, 0, len(optBuf))
		return pt
	}
	return &PieceTable{original: []byte(""), added: []byte(""), nodes: list.New()}
}

func main() {

	x := NewPT([]byte("HEllo\n\n WOrld"))
	x.Insert([]byte("More Text Here:"))
	x.Insert([]byte("\n\n"))
	x.Insert([]byte("\tMore Text Over Here"))
	x.Display()

}
