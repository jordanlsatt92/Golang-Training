//Jordan Satterfield
package main

import "fmt"

// Create 1 dimensional slice of size 3
// Return slice of user input
func returnSlice() []int{
	row:=make([]int,3)
	for j:=0; j < 3;j++{
		var input int
		fmt.Println("Please enter an integer:")
		fmt.Scan(&input)
		row[j] = input
	}
	return row
}

func main(){

	var matrix [][] int

	//Append user input slice to matrix
	for i:=0; i < 3; i++{
		matrix = append(matrix, returnSlice())
	}

	// Print the matrix ints by row
	for i:=0; i < len(matrix);i++{
		for j:=0; j <= i; j++{
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}