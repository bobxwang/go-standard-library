##### 空接口
- 使用空接口可实现接收任意类型的函数参数
  ``` go
  func show(a interface{}) {
      fmt.Printf("type:%T value:%v\n", a, a)
  }
  ```
- 使用空接口可保存任意值的字典
  ``` go
  var studentInfo = make(map[string]interface{})
  studentInfo["name"] = "李白"
  studentInfo["age"] = 18
  studentInfo["married"] = false
  fmt.Println(studentInfo)
  ```