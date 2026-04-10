package graphter

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"siddh.com/compiler"
)

func GetTermValue(term compiler.Term, x int, y int) float64 {
	var value float64 = 0
	if term.Type == "N" {
		if len(term.Subterm) == 0 {
			value, _ := strconv.ParseFloat(term.Constant, 64)

			if term.Variable == "x" {
				value *= float64(x)
			}

			if term.Variable == "y" {
				value *= float64(y)
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

func GetComplexTermValue(term compiler.Term, x int, y int, activeTermsIds []string, expression map[string]compiler.Term) float64 { //all the terms which may have subterms are called complex terms
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
func add(firstTerm compiler.Term, secondTerm compiler.Term, x int, y int, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 + a2
}

func subtract(firstTerm compiler.Term, secondTerm compiler.Term, x int, y int, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 - a2
}

func devide(firstTerm compiler.Term, secondTerm compiler.Term, x int, y int, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 - a2
}

func multiply(firstTerm compiler.Term, secondTerm compiler.Term, x int, y int, activeTermsIds []string, expression map[string]compiler.Term) float64 {
	a1 := GetComplexTermValue(firstTerm, x, y, activeTermsIds, expression)
	a2 := GetComplexTermValue(secondTerm, x, y, activeTermsIds, expression)

	return a1 - a2
}

func powerToTheExponent(value float64, exponent float64) float64 {
	return math.Pow(value, exponent)
}

func GetPoints(equation [2]map[string]compiler.Term) {

	var leftHandSide map[string]compiler.Term = equation[0]
	var RightHandSide map[string]compiler.Term = equation[1]

	// 1. Get a stable order of term IDs for deterministic evaluation
	var originalOrderLeft []string
	var originalOrderRight []string

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

	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100; y++ {
			// 2. Create a fresh copy of IDs for every single (x, y) coordinate
			activeTermsIdsLeft := make([]string, len(originalOrderLeft))
			copy(activeTermsIdsLeft, originalOrderLeft)
			var lhsValue = calculateValueOfExpression(activeTermsIdsLeft, leftHandSide, x, y)

			activeTermsIdsRight := make([]string, len(originalOrderRight))
			copy(activeTermsIdsRight, originalOrderRight)

			var RhsValue = calculateValueOfExpression(activeTermsIdsRight, RightHandSide, x, y)

			if lhsValue == RhsValue {
				fmt.Print("x = ")
				fmt.Print(x)
				fmt.Print(" y = ")
				fmt.Print(y)
				fmt.Print("\n")
			}

		}
	}
}
