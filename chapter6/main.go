package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	file, err := os.Open("/Users/wangx/go/src/ehr/main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dir, fileName := filepath.Split("/Users/wangx/go/src/ehr/main.go")
	fmt.Println(dir)
	fmt.Println(fileName)
}
