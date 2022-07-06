package main

import (
	"fmt"
	"strconv"
)

func main() {
   // var power2 int64  = 1
   // var  a int64 = 1

   // for {
   //    if (a > 10){
   //       break // exit the loop when we reach 10
   //    }

   //    fmt.Println("2 to the power of " + strconv.FormatInt(a,10) +" is equal to "+ strconv.FormatInt(power2,10))
         
   //    power2 += power2
   //    a++
   // }


   var a int = 20
   var b int = 30
   fmt.Println("a = " + strconv.Itoa(a))
   fmt.Println("b = " + strconv.Itoa(b))

   if (a > b){
      goto MESSAGE1 //this will jump the execution to where MESSAGE1 is defined
   } else {
      goto MESSAGE2
   }

   MESSAGE1: // We define a label that we can use in a goto statement
      fmt.Println("a is greater than b")

   MESSAGE2:
      fmt.Println("b is greater than a")
}