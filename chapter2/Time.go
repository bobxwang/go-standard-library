package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	var timestamp int64 = 1651918243
	timeObj := time.Unix(timestamp, 0) // 将时间戳转成时间
	fmt.Printf("timeObj:%v\n", timeObj)

	fmt.Println(now.Format("2006-01-02 15:04:05")) // 24小时制

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	o, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(o)
	fmt.Println(o.Sub(now))
}
