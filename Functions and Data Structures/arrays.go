// package main

// // Jordan Satterfield
// import (
// 	"fmt"
// )

// func main() {

// var arr1 = [10] int {1,2,3,4,5,6,7,8,9,10}
// var arr2 [10] int

// for index, value := range arr1{
// 	arr2[9 - index] = value
// }

// fmt.Println(arr1)
// fmt.Println(arr2)

// x:=[]int{1,2,3,7,8,9}
// x=append(x,4)
// fmt.Println(x)
// x:=[]int{1,2,3,4,5,6,7}
// y:=[]int{-1,-2,-3}
// copy(x[3:],y)
// fmt.Println(x)

// 	matrix := [3] [3] int {
//         {1, 2, 3},
//         {4, 5, 6},
//         {7, 8, 9},
//     }

//   fmt.Println(matrix)
//   fmt.Println(matrix[0][0])
//   fmt.Println(matrix[1][2])

// arr1:=[]int{1,2,3,4,5,6,7,8,9}
// var position int
// var direction string
// fmt.Print("Please enter a position: ")
// fmt.Scan(&position)
// fmt.Print("Please enter a direction: ")
// fmt.Scan(&direction)

// slice1:=arr1[(position)%len(arr1):]
// slice2:=arr1[0:(position)%9]
// if direction == "left"{
// 	a:= [9]int{}
// 	slice1 = append(slice1,slice2...)
// 	a = append(a, slice1...)
// }

// fmt.Println(slice1, slice2)
// }

package main

// Jordan Satterfield
import "fmt"


func main() {
  // declare a two-dimensional array with two sizes instead of one
  matrix := [3] [3] int {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
	
  for i:=0; i<len(matrix);i++{
	var rowSum int = 0
	var colSum int = 0
	for j:=0; j<len(matrix);j++{
		rowSum+=matrix[i][j]
		colSum+=matrix[j][i]
	}
	fmt.Printf("Row #%d sum = %d\n",i+1, rowSum)
	fmt.Printf("Col #%d sum = %d\n",i+1, colSum)
  }
}