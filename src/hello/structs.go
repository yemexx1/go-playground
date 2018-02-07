package main

import (
	"fmt"
	"hello/model"
)

func main() {
	jumperList := model.GetList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}
