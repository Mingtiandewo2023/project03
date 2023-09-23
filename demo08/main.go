package main
import "fmt"
func main(){
	var age =23
	fmt.Println (&age)
	var ptr *int =&age
	fmt.Println (ptr)
	fmt.Println (*ptr)
}