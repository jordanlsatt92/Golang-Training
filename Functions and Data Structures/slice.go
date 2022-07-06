package main

import (
	"fmt"
)

func main() {
  // define an array with 7 elements
//   numbers := [7]int{0,1,2,5,798,43,78}
//   fmt.Println(numbers)

//   // define a slice s based on numbers the array
//   s := numbers[1:5]
//   fmt.Println(s)

//   fmt.Println("Length of slice:", len(s))

//   fmt.Println("Slice capacity:", cap(s))

  // Using make to create slice

  // doubles the size everytime capacity is surpassed
//   slice1:=make([]int, 1,1)
//   for i:=0; i<=10; i++{
// 	slice1 = append(slice1, i)
// 	fmt.Printf("%d, %d\n", len(slice1), cap(slice1))
//   }  
  //fmt.Printf("%d, %d", len(slice1), cap(slice1))

  var records [][] string

  student1:=make([]string,3)
  student1[0] ="ABC"
  student1[1] ="XYZ"
  student1[2] = "AXZ"

  records = append(records, student1)

  fmt.Println(records)
  student2:= make([]string,3)
  student2[0]="123"
  student2[1]= "456"
  student2[2] = "789"
  records = append(records, student2)

  fmt.Println(records)
}