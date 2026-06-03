package main

import (
	"siddh.com/compiler"
	"siddh.com/graphter"
)

func main() {
	equation := compiler.Parse("((1x)^1 (+)^1 (1y)^1)^1 (+)^1 ((1x)^1 (+)^1 (1y)^1)^1 = (10)^1  (+)^1 (10x)^1")
	graphter.GetPoints(equation)
}
