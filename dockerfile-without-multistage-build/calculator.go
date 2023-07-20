package main

//to use print function
import "fmt"

func main(){

	var num1 int
    var num2 int
    var choice int=0
    var result int=0

	fmt.Println("enter number 1")
	fmt.Scan(&num1)
	fmt.Println("enter number 2")
	fmt.Scan(&num2)
    fmt.Println("Select the operation you want to perform")
    fmt.Println("1: Addition") 
    fmt.Println("2: Subtraction") 
    fmt.Println("3: Multiplication") 
    fmt.Println("4: Division") 
    fmt.Print("Enter choice: ") 
    fmt.Scanf("%d",&choice) 
    
     switch choice{ 
       case 1: 
            result=num1+num2
            fmt.Printf("Addition is: %d",result) 
       case 2: 
            result=num1-num2
            fmt.Printf("Subtraction is: %d",result) 
       case 3: 
            result=num1*num2
            fmt.Printf("Multiplication is: %d",result) 
       case 4: 
            result=num1/num2
            fmt.Printf("Division is: %d",result) 
       default:  
            fmt.Println("Invalid value") 

}


