package main

import (
	"fmt"
	"strconv"
)

func main(){

	var a int64 = 68
	fmt.Printf("%c\n", a)

	for i:=1; i <=68; i++{
		fmt.Println("Number base 10 for decimals "+ strconv.FormatInt(int64(i),10))
		fmt.Println("Number base 2 for decimals " + strconv.FormatInt(int64(i),2))
	}
}