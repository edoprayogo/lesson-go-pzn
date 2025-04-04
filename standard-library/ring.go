package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data = ring.New(3)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(a any) {
		fmt.Println(a)
	})

}
