time.Duration是time包定义的一个类型，代表的是两个时间点间经过的时间，以纳秒为单位，表示一段时间间隔

##### 时间操作

- Add
    ``` go
    func (t Time) Add(d Duration) Time
    ```
    > 求一个小时后的时间
    ``` go
    func main() {
        now := time.Now()
        later := now.Add(time.Hour)
        fmt.Println(later)
    }
    ```
- Sub
    ``` go
    func (t Time) Sub(u Time) Duration
    ```
- Equal
    ``` go
    func (t Time) Equal(u Time) bool
    ```

##### 定时器
> 本质上是一个通道(channel)
``` go
func tickDemo() {
    ticker := time.Tick(time.Second)  // 定义一个1秒时间的间
    for i := range ticker {
        fmt.Println(i) 
    }
}
```