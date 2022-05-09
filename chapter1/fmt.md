fmt 实现了类似C语言中的printf跟scanf的格式化I/O，主要分为向外输出及获取输入内容两部分


##### Print
> 将内容输出到系统的标准输出
``` go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

##### FPrint
> 将内容输出到一个io.Writer接口类型的变量w中
``` go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```

##### Sprint
> 把传入的数据生成一个字符串返回
``` go
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```

##### Errorf
> 根据format参数生成格式化字符串并返回一个包含该字符串的错误 
``` go
func Errorf(format string, a ...interface{}) error
```

##### Scan
> 从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符
``` go
func Scan(a ...interface{}) (n int, err error)
```

##### Scanf
> 从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中
``` go
func Scanf(format string, a ...interface{}) (n int, err error)
```

##### Scanln
> 类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
``` go
func Scanln(a ...interface{}) (n int, err error)
```

###### bufio.NewReader
``` go
func bufioDemo() {
    reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
    fmt.Print("请输入内容：")
    text, _ := reader.ReadString('\n') // 读到换行
    text = strings.TrimSpace(text)
    fmt.Printf("%#v\n", text)
}
```