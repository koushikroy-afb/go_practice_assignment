package main
import "fmt"


func main(){
	var name=" string in GO using var"
	var name2 string
	name2=" assigning after declaration"
	name3:= " shorthand declaration"
	var age int 
	fmt.Println("Enter Age")
	fmt.Scanln(&age);
	fmt.Println("Name1:",name)
	fmt.Println("Name2:",name2)
 fmt.Println("Name3:",name3)

}