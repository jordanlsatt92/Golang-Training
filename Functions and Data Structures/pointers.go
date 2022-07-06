// package main

// // Jordan Satterfield
// import "fmt"

// func changeArray(array []int){
// 	var arrayptr [] *int
// 	for i:= 0; i < len(array); i++{
// 		arrayptr = append(arrayptr,&array[i])
// 		*arrayptr[i] ++
// 	}
// }

// func main(){

// 	array := []int {1,2,3}
// 	fmt.Println(array)
// 	changeArray(array)
// 	fmt.Println(array)
// 	array2 := []int {1,2,3,4,5,6,7}
// 	fmt.Println(array2)
// 	changeArray(array2)
// 	fmt.Println(array2)
// }

// package main

// // Jordan Satterfield
// import "fmt"

// func main(){

// 	x:=10
// 	px:=&x
// 	ppx:=&px
// 	fmt.Println("x=", x,"px=", *px, "ppx=", **ppx)
// 	*px=*px+2
// 	fmt.Println("x=", x,"px=", *px, "ppx=", **ppx)
// 	**ppx=**ppx + 10
// 	fmt.Println("x=", x,"px=", *px, "ppx=", **ppx)
// }

// package main

// // Jordan Satterfield
// import "fmt"

// func findDataType(i interface{}){
// 	switch x:=i.(type){
// 	case int:
// 		fmt.Println("Integer", x)
// 	case string:
// 		fmt.Println("String", x)
// 	case float64:
// 		fmt.Println("float", x)
// 	case nil:
// 		fmt.Println("nil type", x)
// 	default:
// 		fmt.Println("default", x)
// 	}
// }

// func main(){

// 	findDataType(10)
// 	findDataType(102.34343)
// 	findDataType("Hello")
// 	var k interface{}
// 	findDataType(k)

// }

// package main

// // Jordan Satterfield
// import "fmt"

// type account struct {
// 	number string
// 	balance float64
//  }

// func main(){

// 	a:= account{}
// 	a.balance = 123.456
// 	a.number = "123456"
// 	fmt.Println(a)

// 	b:=new(account)
// 	b.balance = 123.45
// 	b.number = "54321"
// 	fmt.Println(b)

// }

package main

// Jordan Satterfield
import "fmt"

type bankAccount struct{
	name string
	accountNum int
	balance float64
}

// deposit; negative should show error

// withdrawel; must maintain minimun balance

func main(){

	a:= account{}
	a.balance = 123.456
	a.number = "123456"
	fmt.Println(a)

	b:=new(account)
	b.balance = 123.45
	b.number = "54321"
	fmt.Println(b)
	
}