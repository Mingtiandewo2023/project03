package main
import(
	"fmt"
	"unsafe"
)

func main(){
	var num1 =23
	fmt.Printf("num1的类型是：%T",num1)
    fmt.Println()
	fmt.Println(unsafe.Sizeof(num1))
}