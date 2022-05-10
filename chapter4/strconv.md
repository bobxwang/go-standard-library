strconv 包实现了基本数据类型与字符串类型的转换

##### string2int
``` go
func Atoi(s string) (i int, err error)
func Itoa(i int) string
```

##### ParseXXX
``` go
func ParseBool(str string) (value bool, err error)
// base 指定进制(2到36)，bitSize 指定结果必须能无溢出赋值的整数类型
func ParseInt(s string, base int, bitSize int) (i int64, err error)
// 类似ParseInt但不接受正负号，用于无符号整形 
func ParseUnit(s string, base int, bitSize int) (i uint64, err error)
func ParseFloat(s string, bitSize int) (f float64, err error)
```

##### FormatXXX
``` go
func FormatBool(b bool) string
func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
```