package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {

	if think == "sb" {
		talk = "你是个大傻比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{}
	fmt.Println(peo.Speak("bitch"))
}
