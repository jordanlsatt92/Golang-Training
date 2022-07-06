// package main

// import "fmt"

// func loop(beg , end int){
// 	for i:=beg; i<=end; i++{
// 		fmt.Println("", i)
// 	}
// }

// func sum(n int) int{
// 	s:=0
// 	for i:=1; i <=n; i++{
// 		s+=i
// 	}
// 	return s
// }

// func sum2(beg, end int) int{
// 	s:=beg
// 	for i:=beg+1; i <= end; i++{
// 		s+=i
// 	}
// 	return s
// }

// func multi(x, y int) int{
// 	return x*y
// }

// func add(a,b int) int{
// 	return a+b
// }

// type Operator func(x,y int) int

// func print(f Operator, x, y int){
// 	fmt.Println(f(x,y))
// }

// Closure
// 1 1 2 3 5 8 13
// func fibonacii() func() int{
// 	first, second:= 0, 1
// 	return func() int{
// 		temp:=first
// 		first, second = second, first+second
// 		return temp
// 	}
// }

// func prime() func() int{
// 	var number = 2
// 	return func() int{
// 		isPrime:=true
// 		for i:=2; i < number; i++{
// 			if number % i == 0{
// 				isPrime = false
// 				break
// 			}
// 		}
// 		temp:=
// 	}
// }

// func main(){
// print(add, 10,12)
// print(multi, 10,12)
// fmt.Println("hello world")
// loop(1,5)
// loop(10,20)
//fmt.Println(sum(5))
// a,b,c:=multi(1,2,3)
// fmt.Println(a,b,c)
//Anonymous function
// sumx:=func(a,b,c int) int{
// 	return a+b+c
// }(3,4,5)
// fmt.Println(sumx)

// 	f:=fibonacii()

// 	for i:=0; i<=15; i++{
// 		fmt.Println(f())
// 	}
// }

package main

import "fmt"

func factorial(n int) int{
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}

func prime(n int) func() int{
	
}

func main(){
	fmt.Println(factorial(4))
}