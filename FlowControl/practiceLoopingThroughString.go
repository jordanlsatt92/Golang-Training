package main

// Jordan Satterfield
import (
	"fmt"
	"unicode"
)

func main() {
   // var message string = "HELLO WORLD"

   // fmt.Println(message)

   // for i := 0; i < len(message); i++ {
   //    fmt.Println(string(message[i]))
   // }

   // s:=""
   // for i:='A'; i < 'E'; i++{
   //    s+=string(i)
   //    fmt.Println(s)
   // }

   // name:="Rober Pike"
   // reversedName:=""

   // for i:=len(name)-1; i >=0; i --{
   //    reversedName += string(name[i])
   // }

   // fmt.Println(reversedName)

   // for pos,v:=range name{
   //    fmt.Println("Pos ", pos,"   ", v, "   ", string(v))
   // }

   // var message string = "HELLO WORLD"

   // fmt.Println(message)

   // for i := len(message) - 1; i >=0; i-- {
   //    fmt.Print(string(message[i]))
   // }

   // for pos,v:=range message{
   //    fmt.Println("Pos ", pos,"   ", v, "   ", string(v))
   // }

   // Reverse robert pyke but keep names in original spot

   value:= "Hello World WELLcome"

   result:=""
   for pos,val:= range value{

      if unicode.IsSpace(val){
         for i:=pos-1; i >=0; i--{
            result += string(value[i])
         }
         result += " "
         for i:=len(value)-1; i> pos; i--{
            result += string(value[i])
         }
      }
   }
   fmt.Println(result)
}