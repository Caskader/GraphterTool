package main

import (
	"encoding/json"
	"fmt"
	"os"

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

type Message struct {
	Latex string
}

func ProcessEquation(equationRaw json.RawMessage, startingPoint [2]int, endingPoint [2]int) ([][2]float64, error) {
	var m Message

	err := json.Unmarshal(equationRaw, &m)

	if err != nil {
		fmt.Print(err)
	}
	formattedEquation := compiler.Format(m.Latex)
	// pretty-print the incoming JSON equation payload to the terminal
	// var pretty bytes.Buffer
	// json.Indent(&pretty, equationRaw, "", "  ")
	// fmt.Println("Equation (raw):", string(equationRaw))
	// fmt.Println("Equation (JSON):\n" + pretty.String())
	fmt.Fprintln(os.Stdout, "raw equation ", m.Latex)
	fmt.Fprintln(os.Stdout, "formatted equation ", formattedEquation)
	equation := compiler.Parse(formattedEquation)
	points := graphter.GetPoints(equation, 2, startingPoint, endingPoint)

	// otherwise there's no points to return
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
