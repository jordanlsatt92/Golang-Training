package main

// Jordan Satterfield
import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// add array of transactions
type bankAccount struct{
	name string
	accountNum int
	balance float64
	transactions [] *transaction
}

type transaction struct{
	time string
	date string
	amount float64
}

// deposit; negative should show error
func (a *bankAccount) deposit(amount float64) float64{
	if amount < 0.0 {
		fmt.Println("You cannot deposit a negative amount of money.")
		return 0
	}else{
		currentTime:=time.Now()
		var date string = (currentTime.Month().String()) + " " + fmt.Sprint(currentTime.Day()) + " " + fmt.Sprint(currentTime.Year())
		time:=fmt.Sprint(currentTime.Hour()) + ":" + fmt.Sprint(currentTime.Minute())
		a.balance += amount
		a.transactions = append(a.transactions, &transaction{time, date, amount})
		return a.balance
	}
}
// withdrawel; must maintain minimun balance
func (a *bankAccount) withdrawel(amount float64) float64{
	if a.balance - amount < 0{
		fmt.Println(a.name,"\nYou cannot withdraw an amount greater than your balance")
		fmt.Println("Your account balance:", a.balance, "\nYour attempted withdrawel:", amount)
		return a.balance
	}else{
		currentTime:=time.Now()
		var date string = (currentTime.Month().String()) + " " + fmt.Sprint(currentTime.Day()) + " " + fmt.Sprint(currentTime.Year())
		time:= fmt.Sprint(currentTime.Hour()) + ":" + fmt.Sprint(currentTime.Minute())
		a.balance -= amount
		a.transactions = append(a.transactions, &transaction{time, date, -amount})
		return a.balance
	}
}

func generateAccount() func() int{
	var number int = 1
	var temp int
	return func() int {
		temp = number
		number += 1
		return temp
	}
}

func sortByName(array []*bankAccount){
	sort.Slice(array, func(i, j int) bool {
		str1:= strings.Split(array[i].name, " ")
		name1:= str1[0]
		str2:= strings.Split(array[j].name, " ")
		name2:= str2[0]
		return name1 < name2
	})
}

func sortByBalance(array []*bankAccount){
	sort.Slice(array, func(i, j int) bool{
		return array[i].balance < array[j].balance
	})
}

func main(){

	accounts:=[]*bankAccount{}
	f:= generateAccount()
	var input string
	var firstName string
	var lastName string
	var accountNum int
	var balance float64
	for input != "done"{
		fmt.Print("First name of account holder (string):")
		fmt.Scanln(&firstName)
		fmt.Print("Last name of account holder (string):")
		fmt.Scanln(&lastName)
		accountNum = f()
		fmt.Print("Initial deposit (float):")
		fmt.Scanln(&balance)
		fmt.Print("Type \"done\" if you do not want create more accounts, hit enter if you do:")
		fmt.Scanln(&input)
		a:=bankAccount{name:firstName + " " + lastName, accountNum:accountNum, balance:balance}
		accounts = append(accounts, &a)
	}

	fmt.Println("All Accounts:")
	for i:= 0; i < len(accounts);i++{
		fmt.Println(*accounts[i])
	}

	sortByName(accounts)
	fmt.Println("\nSorted by first name:")
	for i:= 0; i < len(accounts);i++{
		fmt.Println(*accounts[i])
	}

	sortByBalance(accounts)
	fmt.Println("\nSorted by account balance:")
	for i:= 0; i < len(accounts);i++{
		fmt.Println(*accounts[i])
	}

	// Transactions
	for i:= 0; i < len(accounts);i++{
		accounts[i].deposit(50.66)
		accounts[i].withdrawel(30.32)
		accounts[i].deposit(100.32)
		accounts[i].withdrawel(75.88)
		accounts[i].deposit(10.21)
		accounts[i].withdrawel(300)
	}

	fmt.Println("\nTransactions for each user:")
	for i:=0; i < len(accounts);i++{
		fmt.Printf("%s's transactions\n", accounts[i].name)
		for j:=0; j < len(accounts[i].transactions); j++{
			fmt.Printf("Time: %s, Date: %s, Transaction: %.2f\n",accounts[i].transactions[j].time, accounts[i].transactions[j].date, accounts[i].transactions[j].amount)
		}
	}

	

	// for i:=0; i< len(accounts[0].transactions);i++{
	// 	fmt.Println(accounts[0].transactions[i])
	// }
	
}