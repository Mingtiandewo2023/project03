package main
import "fmt"
func test(args...int){
	for i := 0;i < len(args);i++{
		fmt.Println(args[i])
	}
}
func main(){
	test(3333)
}