package main

import (
	"siddh.com/compiler"
	"siddh.com/graphter"
	"siddh.com/network"
	"fmt"
)

func CalculatePoints(response network.ResponseData) {
	// "((1x)^1 (+)^1 (1y)^1)^1 (+)^1 ((1x)^1 (+)^1 (1y)^1)^1 = (10)^1  (+)^1 (10x)^1"
	_ = response.Id
	
	querry := response.Message
	fmt.Println(querry)
	equation := compiler.Parse(querry)
	points := graphter.GetPoints(equation)
	_ = points
}

func ProcessEquation(equationStr string) ([][2]int, error) {
	equation := compiler.Parse(equationStr)
	points := graphter.GetPoints(equation)
	return points, nil
}

func main() {
	// equation := compiler.Parse("((1x)^1 (+)^1 (1y)^1)^1 (+)^1 ((1x)^1 (+)^1 (1y)^1)^1 = (10)^1  (+)^1 (10x)^1")
	equation := compiler.Parse("(x)^1 = (y)^1")
	fmt.Print(graphter.GetPoints(equation))
	fmt.Print(equation)
	// network.Start(CalculatePoints, ProcessEquation)
}
