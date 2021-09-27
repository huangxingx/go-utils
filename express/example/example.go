package main

import (
	"fmt"

	"github.com/huangxingx/go-utils/express"
)

func main() {
	expressStr := "1+2+3"
	exp := express.NewExpress(expressStr)
	result := exp.Execute(nil)
	fmt.Printf("%s = %.2f", expressStr, result.(float64))
}
