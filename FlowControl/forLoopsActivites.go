package main

// Jordan Satterfield
import (
	"fmt"
	"os"
)

func main(){

	// Activity 7

	// var x,y, answer int;

	// fmt.Print("Please enter a number for X: ")
	// fmt.Scanf("%d", &x)
	// fmt.Print("Please enter a number for Y: ")
	// fmt.Scan( &y)

	// answer = x
	// for i:=2; i <= y; i++{
	// 	answer*= x
	// }

	// fmt.Println(x, "to the power of", y, "is", answer)


	////////////// Activity 8 e^x ///////////////////
	// Jordan Satterfield

	// var power,nTerms int	
	// var answer float64
	// fmt.Print("Please enter the desired power for e: ")
	// fmt.Scan(&power)
	// fmt.Print("Please enter the number of terms: ")
	// fmt.Scan(&nTerms)

	// answer = 1
	// var factorial int 
	// for i := nTerms; i > 0; i--{
	// 	factorial = i
	// 	for j := i-1; j > 0; j--{
	// 		factorial *= j
	// 	}
	// 	answer += math.Pow(float64(power), float64(i)) / float64(factorial)
	// }
	// fmt.Printf("e^%d = %.2f",power,answer)


	


	///////////////// Activity 9 /////////////////
	// Jordan Satterfield

	// var number int
	// fmt.Println("Please enter an integer: ")
	// fmt.Scan(&number)

	// // Number of digits in number
	// var numberString string = strconv.Itoa(number)
	// fmt.Println("Number of digits in", number, ":", len(numberString))

	// // First and Last digit
	// fmt.Println("First digit:", string(numberString[0]), "\nLast digit:", string(numberString[len(numberString)-1]))

	// // Sum of digits in number
	// var sum int
	// for i:=0; i < len(numberString);i++{
	// 	var char string = string(numberString[i])
	// 	num,_:= strconv.Atoi(char)
	// 	sum+=num
	// }
	// fmt.Println("Sum of numbers:", sum)

	// // Product of digits in number
	// product,_:= strconv.Atoi(string(numberString[0]))
	// for i:=1; i < len(numberString); i++{
	// 	num,_:= strconv.Atoi(string(numberString[i]))
	// 	fmt.Println(product * num)
	// 	product = product * num
	// 	fmt.Println(num, product)
	// }
	// if len(numberString) < 2{
	// 	product = 0
	// }
	// fmt.Println("Product of numbers:", product)

	// // Determine if number is prime
	// var isPrime bool = true
	// for i:=2; i < number; i++{
	// 	if (number % i == 0){
	// 		isPrime = false
	// 		break
	// 	}
	// }
	// if number == 1{
	// 	isPrime = false
	// }
	// if isPrime{
	// 	fmt.Println(number, "is prime.")
	// }else{
	// 	fmt.Println(number, "is not prime.")
	// }

	// // Calculate factorial of number
	// var factorial int = number
	// for i:=number - 1; i > 0; i--{
	// 	factorial*=i
	// }
	// fmt.Println("Factorial of", number,"is", factorial)


	///////////Activity 10 ///////////
	// Jordan Satterfield

	var rows int
	fmt.Println("Please enter the number of rows")
	fmt.Scan(&rows)
	if rows < 1{
		fmt.Println("You must enter a positive integer")
		os.Exit(1)
	}

	for i:=1; i <= rows; i++{
		for j:=1; j <= i; j++{
			fmt.Print(j)
		}
	}

	/////////// Activity 11 /////////////
	// Jordan Satterfield

	// var firstNum, secondNum, divisor, lowestInt int
	// fmt.Print("Please enter an integer: ")
	// fmt.Scan(&firstNum)
	// fmt.Print("Please enter a second integer: ")
	// fmt.Scan(&secondNum)
	// divisor = 1
	// if firstNum < secondNum || firstNum == secondNum{
	// 	lowestInt = firstNum
	// }else{
	// 	lowestInt = secondNum
	// }
	// if firstNum == 0 || secondNum == 0{
	// 	fmt.Println("Cannot find greatest common divisor if either integer = 0")
	// 	os.Exit(1)
	// }

	// // Find greatest common divisor
	// for i:=divisor; i < lowestInt; i++{
	// 	if (firstNum % i == 0) && (secondNum % i == 0){
	// 		divisor = i
	// 	}
	// }
	// fmt.Println("The lowest common divisor =", divisor)
}
