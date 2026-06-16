package graphter

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"siddh.com/compiler"
)

type Chunk struct {
	id            int
	points        [][2]float64
	startingPoint [2]int
	endingPoint   [2]int
}

func GetTermValue(term compiler.Term, x float64, y float64) float64 {
	var value float64 = 0
	if term.Type == "N" {
		if len(term.Subterm) == 0 {
			value, _ := strconv.ParseFloat(term.Constant, 64)

			if term.Variable == "x" {
				value *= x
			}

			if term.Variable == "y" {
				value *= y
			}
			exponent, err := strconv.ParseFloat(term.Exponent, 64)
			if err != nil {
				// handle error
			}
			value = powerToTheExponent(value, exponent)
			return value
		} else {
			// when the term has subterms
			fmt.Print("erororor")
			value = 0
		}
	}

	// fmt.Println(value)
	return value
}

func GetComplexTermValue(term compiler.Term, x float64, y float64, activeTermsIds []string, expression map[string]compiler.Term) float64 { //all the terms which may have subterms are called complex terms
	var value float64 = 0

	if term.Type == "N" {

		if len(term.Subterm) == 0 {
			value = GetTermValue(term, x, y)
			return value
		} else {
			var subactiveTermsIds []string = make([]string, len(term.Subterm))
			copy(subactiveTermsIds, term.Subterm)
			value = calculateValueOfExpression(subactiveTermsIds, expression, x, y)

			exponent, err := strconv.ParseFloat(term.Exponent, 64)
			if err != nil {
				// handle error
			}

			value = powerToTheExponent(value, exponent)

			return value
		}

	}

	return value
}

// function to add two terms in an equation
func add(firstTerm compiler.Term, secondTerm compiler.Term, x float64, y float64, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 + a2
}

func subtract(firstTerm compiler.Term, secondTerm compiler.Term, x float64, y float64, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 - a2
}

func devide(firstTerm compiler.Term, secondTerm compiler.Term, x float64, y float64, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 - a2
}

func multiply(firstTerm compiler.Term, secondTerm compiler.Term, x float64, y float64, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 - a2
}

func powerToTheExponent(value float64, exponent float64) float64 {
	return math.Pow(value, exponent)
}

func GetPointsInChunk(equation [2]map[string]compiler.Term, accuracy int, c Chunk) Chunk {

	var leftHandSide map[string]compiler.Term = equation[0]
	var RightHandSide map[string]compiler.Term = equation[1]
	var points [][2]float64
	var resolution = int(math.Pow10(accuracy))
	// 1. Get a stable order of term IDs for deterministic evaluation
	var originalOrderLeft []string
	var originalOrderRight []string
	var initialPoint = c.startingPoint
	var finalPoint = c.endingPoint

	for id := range leftHandSide {
		if strings.Count(id, ".") == 1 {
			originalOrderLeft = append(originalOrderLeft, id)
		}
	}
	for id := range RightHandSide {
		if strings.Count(id, ".") == 1 {
			originalOrderRight = append(originalOrderRight, id)
		}
	}

	sort.Strings(originalOrderLeft)
	sort.Strings(originalOrderRight)
	// here x and y are multipled by resolution for calculation purposes
	for x := initialPoint[0] * resolution; x <= (finalPoint[0] * resolution); x++ {
		for y := initialPoint[1] * resolution; y <= finalPoint[1]*resolution; y++ {
			// 2. Create a fresh copy of IDs for every single (x, y) coordinate
			activeTermsIdsLeft := make([]string, len(originalOrderLeft))
			copy(activeTermsIdsLeft, originalOrderLeft)
			var lhsValue = calculateValueOfExpression(activeTermsIdsLeft, leftHandSide, float64(x)/float64(resolution), float64(y)/float64(resolution))
			activeTermsIdsRight := make([]string, len(originalOrderRight))
			copy(activeTermsIdsRight, originalOrderRight)

			var RhsValue = calculateValueOfExpression(activeTermsIdsRight, RightHandSide, float64(x)/float64(resolution), float64(y)/float64(resolution))

			if lhsValue == RhsValue {
				// fmt.Print("x = ")
				// fmt.Print(x)
				// fmt.Print(" y = ")
				// fmt.Print(y)
				// fmt.Print("\n")
				point := [2]float64{float64(x) / float64(resolution), float64(y) / float64(resolution)}
				points = append(points, point)
			}

		}
	}

	c.points = points
	return c
}

func GetPoints(equation [2]map[string]compiler.Term, accuracy int, initialPoint [2]int, finalPoint [2]int) [][2]float64 {

	var TotalPoints []Chunk
	var deltaX int = (finalPoint[0] - initialPoint[0]) / 3
	var deltaY int = (finalPoint[1] - initialPoint[1]) / 3
	for x := 1; x <= 3; x += 1 {
		for y := 1; y <= 3; y += 1 {
			var c Chunk = Chunk{
				id:            x + y - 2,
				points:        nil,
				startingPoint: [2]int{initialPoint[0] + deltaX*(x-1), initialPoint[0] + deltaY*(y-1)},
				endingPoint:   [2]int{initialPoint[0] + deltaX*x, initialPoint[1] + deltaY*y},
			}
			TotalPoints = append(TotalPoints, GetPointsInChunk(equation, accuracy, c))
		}

	}

	var points [][2]float64

	for _, C := range TotalPoints {
		points = append(points, C.points...)
	}
	return points
}
