package main
import "fmt"
func main(){
	var score int=101
	switch score/10{
		case 10 :{
			fmt.Println("您的等级为A等")
		}
		case 9 :{
			fmt.Println("您的等级为A等")
		}
		case 8 :{
			fmt.Println("您的等级为B等")
		}
		case 7 :{
			fmt.Println("您的等级为C等")
		}
		case 6 :{
			fmt.Println("您的等级为D等")
		}
		case 5 :{
			fmt.Println("您的等级为E等")
		}
		case 4 :{
			fmt.Println("您的等级为E等")
		}
		case 3 :{
			fmt.Println("您的等级为E等")
		}
		case 2 :{
			fmt.Println("您的等级为E等")
		}
		case 1 :{
			fmt.Println("您的等级为E等")
		}
		case 0 :{
			fmt.Println("您的等级为E等")
		}
	    default:{
			fmt.Println("您的成绩有误")
		}
	}
}