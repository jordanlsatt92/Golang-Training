package main

import "fmt"

type Node struct{
	val int
	next *Node
}

type Queue struct{
	head *Node
	tail *Node
	count int
}

func NewQueue() *Queue{
	return &Queue{}
}

func (q *Queue) enqueue(n *Node){
	if q.count==0{
		q.head = n
		q.tail = n
		q.head.next = q.tail
		q.count++
		fmt.Println(q.count)
		return
	}
	if q.count > 0{
		q.tail.next = n
		q.tail = n
		q.tail.next = nil
		q.count++
		fmt.Println(q.count)
		return
	}
}

func (q *Queue) printQueue(){
	var temp *Node = q.head
	for temp != nil{
		fmt.Print(temp.val, " ")
		temp = temp.next
	}
	fmt.Print("\n")
}

func (q *Queue) dequeue() *Node{
	if q.count == 0 {
		fmt.Println("Queue is empty")
		return nil
	}else if q.count == 1{
		var temp *Node = q.head
		q.head,q.tail = nil, nil
		q.count--
		return temp
	}else{
		var temp *Node = q.head
		q.head = q.head.next
		q.count--
		return temp
	}
}

func main(){

	var q *Queue = NewQueue()
	q.enqueue(&Node{val: 1})
	q.enqueue(&Node{val: 2})
	q.enqueue(&Node{val: 3})
	q.enqueue(&Node{val: 4})
	q.enqueue(&Node{val: 5})
	//q.printQueue()
	fmt.Println(q.dequeue().val,q.dequeue().val,q.dequeue().val,q.dequeue().val,q.dequeue().val)//, q.dequeue())
}