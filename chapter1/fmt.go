package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
递归列出路径下的所有文件
*/
func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}

type user struct {
	name string
}

func printFormat() {

	u := user{name: "abcdefg"}
	fmt.Printf("% + v\n", u)     //格式化输出结构
	fmt.Printf("%#v\n", u)       //输出值的 Go 语言表示方法
	fmt.Printf("%T\n", u)        //输出值的类型的 Go 语言表示
	fmt.Printf("%t\n", true)     //输出值的 true 或 false
	fmt.Printf("%b\n", 1024)     //二进制表示
	fmt.Printf("%c\n", 11111111) //数值对应的 Unicode 编码字符
	fmt.Printf("%d\n", 10)       //十进制表示
	fmt.Printf("%o\n", 8)        //八进制表示
	fmt.Printf("%q\n", 22)       //转化为十六进制并附上单引号
	fmt.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
	fmt.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
	fmt.Printf("%U\n", 1233)     //Unicode表示
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示
	fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
	fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
	fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
	fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
	fmt.Printf("%q\n", "dedede") //双引号括起来的字符串
	fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
	fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
	fmt.Printf("%p\n", &u)       //0x开头的十六进制数表示
}

func main() {
	if 1 > 3 {
		dir := "/Users/wangx/go/src/gopkg.in"
		listAll(dir, 0)
	}

	if 1 > 4 {
		// 输出
		printFormat()
	}

	if 1 > 6 {
		// 输入
		var (
			name    string
			age     int
			married bool
		)
		fmt.Scan(&name, &age, &married)
		fmt.Printf("扫描结果 name:%s, age:%d, married:%t \n", name, age, married)
		fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
		fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	}

	if 1 > 6 {
		// 字符串系列
		a := "gopher"
		b := "hello world"
		fmt.Println(strings.Compare(a, b))   // 比较两个字符串大小，如果相等，返回 0，如果 a<b，返回-1，否则为 1
		fmt.Println(strings.EqualFold(a, b)) // 比较 a,b忽略大小写后是否相等
		fmt.Printf("%q\n", strings.Split("abcd,def,agc,hee", ","))
		fmt.Println(strings.HasPrefix("Gopher", "go"))
		fmt.Println("ba" + strings.Repeat("na", 2))
		bb := bytes.NewBufferString("Hello World, 您好")
		fmt.Println(bb)

		n, err := strconv.ParseInt("128", 10, 8)
		if err != nil {
			fmt.Println("parse int error:", err)
		} else {
			fmt.Println(n)
		}
	}
}
