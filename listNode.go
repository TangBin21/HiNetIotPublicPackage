// listNode
package HiNetEdgeDAS

import (
	// "container/list"
	"fmt"
)

//结点 结构体
type Node struct {
	X    interface{} //节点
	Next *Node       //下一个指针
}

//链表 结构体
type List struct {
	Head   *Node
	Length int
}

//链表 接口
type Method interface {
	Insert(i int, v interface{})
	InsertLast(v interface{})
	Delete(i int)
	DeleteHead()
	GetLength() int
	//SearchOfNode(v interface{}) int
	SearchOfIndex(i int) interface{}
	IsNull() bool
}

func CreateNode(v interface{}) *Node {
	//golang的传参是值传递，如果你传入的不是指针就是传值了，两者是不同的地址，你要改变值就得传一个指针进来
	return &Node{v, nil}
}

func CreateList() *List {
	return &List{CreateNode(nil), 0}
}

func (list *List) Insert(i int, v interface{}) {
	if list == nil {
		return
	}

	s := CreateNode(v)
	pre := list.Head
	for count := 0; count <= i-1; count++ {
		if count == i-1 {
			s.Next = pre.Next
			pre.Next = s
			list.Length++
		}
		pre = pre.Next
	}
}

func (list *List) InsertLast(v interface{}) {
	if list == nil {
		return
	}

	s := CreateNode(v)
	pre := list.Head
	for pre.Next != nil {
		pre = pre.Next
	}

	list.Length++
	pre.Next = s
}

func (list *List) DeleteHead() {
	if list == nil {
		return
	}

	if list.Length == 0 {
		return
	}

	pre := list.Head
	s := pre.Next
	pre.Next = s.Next
	list.Length--
}

func (list *List) Clean() {
	for {
		if list.Length > 0 {
			list.DeleteHead()
		} else {
			break
		}
	}
}

func (list *List) Delete(i int) {
	if i <= 0 || i > list.Length {
		return
	}
	if list == nil {
		return
	}

	var pre, s, t *Node
	var count int
	pre = list.Head
	for count = 0; count < list.Length; count++ {
		if count == i-1 {
			break
		}
		t = pre.Next
		s = t.Next
		pre = t
	}
	s = pre.Next
	pre.Next = s.Next
	list.Length--
}

func (list *List) GetLength() int {
	pre := list.Head
	for pre.Next != nil {
		list.Length++
	}

	return list.Length
}

// func (list *List) SearchOfNode(v interface{}) int {
// 	pre := list.Head.Next
// 	for i := 1; i <= list.Length; i++ {
// 		if pre.X == v {
// 			return i
// 		}
// 		pre = pre.Next
// 	}
// 	return 0
// }

//这里是要返回元素吧
func (list *List) SearchOfIndex(i int) interface{} {
	if list == nil {
		return nil
	}

	pre := list.Head
	for count := 0; count <= list.Length-1; count++ {
		s := pre.Next
		if count == i-1 {
			return s.X
		}
		pre = pre.Next
	}
	return nil
}

func (list *List) IsNull() bool {
	pre := list.Head.Next
	if pre == nil {
		return true
	}
	return false
}

func PrintList(list *List) {
	pre := list.Head //*Node
	fmt.Println("List shows as follow:...")
	for i := 0; i <= list.Length; i++ {
		fmt.Printf("%v\n", pre)
		pre = pre.Next
	}
}
