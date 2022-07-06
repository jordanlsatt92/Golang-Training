package main

// Jordan Satterfield
import "fmt"

func main(){

	var paper1 bool = true
	var pen1 bool = true

	if paper1 && pen1 {
		fmt.Println("You can jot down a note")
	}else {
		fmt.Println("You cannot jot down a note")
	}

	var paper2 bool = false
	var pen2 bool = true

	if paper2 && pen2 {
		fmt.Println("You can jot down a note")
	}else {
		fmt.Println("You cannot jot down a note")
	}
}