package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	timestamp1 := now.Unix()
	fmt.Printf("%d",timestamp1)
// 	now := time.Now() //获取当前时间
// 	fmt.Printf("current time:%v\n",now)

// 	year := now.Year()
// 	month := now.Month()
// 	day := now.Day()
// 	hour := now.Hour()
// 	minute := now.Minute()
// 	second := now.Second()

// 	fmt.Printf("%d-%d-%d-%d:%d:%d\n",year,month,day,hour,minute,second)
}