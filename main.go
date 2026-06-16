package main

import (
	"fmt"

	"siddh.com/compiler"
	"siddh.com/graphter"
	"siddh.com/network"
)

func CalculatePoints(response network.ResponseData) {
	// "((1x)^1 (+)^1 (1y)^1)^1 (+)^1 ((1x)^1 (+)^1 (1y)^1)^1 = (10)^1  (+)^1 (10x)^1"
	_ = response.Id

	query := response.Message
	fmt.Println(query)
	//equation := compiler.Parse(query)
	var points [][2]float64
	_ = points
}

func ProcessEquation(equationStr string, startingPoint [2]int, endingPoint [2]int) ([][2]float64, error) {
	equation := compiler.Parse(equationStr)
	points := graphter.GetPoints(equation, 2, startingPoint, endingPoint)
	return points, nil
}

func main() {
	// equation := compiler.Parse("((1x)^1 (+)^1 (1y)^1)^1 (+)^1 ((1x)^1 (+)^1 (1y)^1)^1 = (10)^1  (+)^1 (10x)^1")
	//equation := compiler.Parse("(1x)^2 = (1y)^1")
	//var p = graphter.GetPoints(equation, 3)
	//for i := 0; i < len(p); i++ {
	//	fmt.Println(p[i])
	//}
	//fmt.Print(equation)
	network.Start(CalculatePoints, ProcessEquation)
}
