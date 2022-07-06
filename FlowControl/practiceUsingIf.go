package main

import "fmt"

func main(){
	var accountValue float32 = 320.11

	if accountValue < 0.01 {
		fmt.Println("The transaction cannot be completed.")
	}else{
		fmt.Println("Transaction completed")
	}
}
