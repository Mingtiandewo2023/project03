package main
import "fmt"

func main(){
	var a,b int
	var pwd =1025011314
	fmt.Print("请输入您的密码:")
	fmt.Scan(&a)
	if a==pwd {
		fmt.Print("请再次输入您的密码：")
		fmt.Scan(&b)
		if b==pwd {
			fmt.Println("登录成功")
		}else{
			fmt.Println("第二次密码错误！")
		}
	}else{
		fmt.Println("密码错误！")
	}
}