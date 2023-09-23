package main
import "fmt"
func Cal(num1 int,num2 int) (int){
	var sum =1
	sum *= num1
	sum *= num2
	return sum 
}
func main(){
	sum := cal(10,20)
	fmt.Println (sum)
}