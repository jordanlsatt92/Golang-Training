package main

import "fmt"

func main(){
	beg:='a'
	end:='d'

	fmt.Println(int(beg))
	for i:=beg; i <= end; i++{
		fmt.Println(string(i))
	}
}