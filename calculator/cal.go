package main

import "fmt"
func main(){
	 var num1,num2 int
var	 operator string
	fmt.Println("Enter Number 1:")
	fmt.Scanln(&num1)
	fmt.Println("Enter Number 2:")
	fmt.Scanln(&num2)
	fmt.Println("Enter Operator (+,-,*,/):")
	fmt.Scanln(&operator)

	switch operator{
	case "+":
		fmt.Println("Result is :" , num1+num2)
	case "-":
		fmt.Println("Result is :" , num1-num2)
	case "*":
		fmt.Println("Result is :" , num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result is :" , num1/num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
	
}