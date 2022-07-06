package main

import "fmt"

func main() {
   var color string = "Green"

   switch(color){
   case "Blue":
      fmt.Println("Blue like the sky")
   case "Red":
      fmt.Println("Red like the sun")
   case "Green":
      fmt.Println("Green like the trees")
   default:
      fmt.Println("Please choose a valid color.")
   }
}