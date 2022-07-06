// package main

// import "fmt"

// type Node struct{
// 	Value int

// }

// type Stack struct{
// 	nodes[] *Node
// 	count int
// }

// func NewStack() *Stack{
// 	return &Stack{}
// }

// func (n *Node) String() string{
// 	return fmt.Sprint(n.Value)
// }

// func (s* Stack) Push(n *Node){
// 	s.nodes = append(s.nodes[:s.count], n)
// 	s.count++
// }

// func (s* Stack) Pop(n *Node) *Node{
// 	if s.count == 0{return nil}
// 	s.count--
// 	return s.nodes[s.count]
// }

// func main(){

// 	s:=NewStack()
// 	s.Push(&Node{10})
// 	s.Push(&Node{12})
// 	s.Push(&Node{14})
// 	s.Push(&Node{15})
// 	fmt.Println(s.Pop(), s.Pop,s.Pop,s.Pop)
// }

package main

import "fmt"
type Node struct{
    Value int
}

type Queue struct{
    nodes[] *Node
    count int
}
func NewQueue() *Queue{
    return &Queue{}
}
func (s *Queue)  Push(n *Node){
    s.nodes=append(s.nodes[:s.count],n)
         
    s.count++//pont to next position
		 
}

func (s *Queue)  Pop()  *Node{
    if s.count==0{ 
        return nil
    }
    if s.count==1{
        s.count--
        var temp *Node = s.nodes[0]
        s.nodes[0] = nil
        return temp
    }
	var temp *Node = s.nodes[0]
    s.count--
	s.nodes = s.nodes[1:]
    return  temp
}
func (n  *Node) String()string{
    return fmt.Sprint(n.Value)
}

func main(){
   s:=NewQueue()
   s.Push( &Node{10} )
   s.Push( &Node{12} )
   s.Push( &Node{14} )
   s.Push( &Node{16} )
    fmt.Println((s.nodes))
    fmt.Println(s.Pop()  ,s.Pop(),s.Pop(),s.Pop())

}
