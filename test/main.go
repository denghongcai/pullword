package main

import (
	"fmt"

	"github.com/denghongcai/pullword"
)

func main() {
	req := pullword.NewRequest("你的姿势水平还远远不够", 1, false)
	result, err := req.Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", result)
}
