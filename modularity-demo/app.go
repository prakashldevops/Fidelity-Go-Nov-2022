package main

import (
	"fmt"

	"github.com/tkmagesh/fidelity-go-nov-2022/modularity-demo/calculator"
)

func main() {
	fmt.Println("modularity-demo executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.PrintStats())
}
