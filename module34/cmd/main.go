package main

import (
	"fmt"
	"module34/pkg/calculate"
)

func main() {
	fmt.Println(calculate.Calculate("cmd/input.txt", "cmd/result.txt"))
}
