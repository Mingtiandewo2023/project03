package main
import "fmt"
func exchange(num1 int,num2 int){
	var t int
	t = num1
	num1 = num2
	num2 = t
	fmt.Println("num1的数值为：",num1)
	fmt.Println("num2的数值为:",num2)
}
func main(){
	var num1,num2 = 10,20
	exchange(num1,num2)
}