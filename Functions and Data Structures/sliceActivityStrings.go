// Jordan Satterfield

package main

import (
	"fmt"
	"strings"
)

func takeInput(n int) []string{
	inputArray:=make([]string,1)
	var input string
	for i:=0; i < n; i++{
		fmt.Print("Please enter a string:")
		fmt.Scan(&input)
		inputArray = append(inputArray, strings.ToLower( input))
	}
	return inputArray
}

func average(array []string) float64{
	var count float64
	for i:=0; i< len(array); i++{
		count += float64(len(array[i]))
	}
	return count / float64(len(array))
}

func idLongWords(array []string, avg float64) [] string{
	longWords:=make([]string,1)
	for i:=0; i < len(array); i++{
		if float64(len(array[i])) > avg{
			longWords = append(longWords, array[i])
		}
	}
	return longWords
}

func idShortWords(array []string, avg float64) [] string{
	shortWords:=make([]string,1)
	for i:=0; i < len(array); i++{
		if float64(len(array[i])) < avg{
			shortWords = append(shortWords, array[i])
		}
	}
	return shortWords
}

func main(){
	inputArray:= takeInput(10)
	fmt.Printf("\nThe average length of words entered : %.2f\n",average(inputArray))
	fmt.Println("Original slice:", inputArray)
	fmt.Println("Words longer than average:", idLongWords(inputArray, average(inputArray)))
	fmt.Println("Words shorter than average:", idShortWords(inputArray, average(inputArray)))
}