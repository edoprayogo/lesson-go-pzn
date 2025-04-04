package main

import (
	"container/list"
	"fmt"
)

func main() {

	var data *list.List = list.New()

	data.PushBack("edo")
	data.PushBack("prayogo")

	head := data.Front()
	fmt.Println("head value", head.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
