package main

// Jordan Satterfield
import "fmt"

func main(){

	n:=7
	sum:=0

	for i:=0; i<=n; i++{
		sum = sum+i
	}

	fmt.Println(sum)
}