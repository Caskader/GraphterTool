package graphter

import (
	"fmt"
	"strconv"

	"siddh.com/compiler"
)

func GetTermValue(term compiler.Term, x int, y int) float64 {

	value, _ := strconv.ParseFloat(term.Constant, 64)

	if term.Variable == "x" {
		value *= float64(x)
	}

	if term.Variable == "y" {
		value *= float64(y)
	}

	// fmt.Println(value)
	return value
}

// function to add two terms in an equation
func add(firstTerm compiler.Term, secondTerm compiler.Term, x int, y int) float64 {
	a1 := GetTermValue(firstTerm, x, y)
	a2 := GetTermValue(secondTerm, x, y)

	return a1 + a2
}

func calculateValueOfExpression(activeTermsIds []string, expression map[string]compiler.Term, x int, y int) float64 {
	var value float64 = 0
	for h := 0; h < len(activeTermsIds); h++ {
		currID := activeTermsIds[h]
		if currID == "D" {
			continue
		} // Skip already processed terms

		term := expression[currID]

		if term.Type == "O" {

			switch term.Operator {
			case "+":
				// Boundary check to prevent crashes
				if h > 0 && h < len(activeTermsIds)-1 {
					leftIdx := h - 1
					rightIdx := h + 1

					val := add(expression[activeTermsIds[leftIdx]], expression[activeTermsIds[rightIdx]], x, y)

					// Update the left term with the new calculated value
					t := expression[activeTermsIds[leftIdx]]
					t.Value = val
					expression[activeTermsIds[leftIdx]] = t

					// Mark operator and right operand as Done
					activeTermsIds[h] = "D"
					activeTermsIds[rightIdx] = "D"

					value = val

				}
			}
		}
	}
	return value
}

// func GetPoints(a map[string]compiler.Term) {

// 	var activeTermsIds []string = []string{}

// 	for i := range a {
// 		activeTermsIds = append(activeTermsIds, i)
// 	}
// 	fmt.Print(activeTermsIds)

// 	var activeTermsCache []string = activeTermsIds

// 	for x := 0; x < 10; x++ {
// 		for y := 0; y < 10; y++ {
// 			activeTermsIds = activeTermsCache
// 			for h := 0; h < len(activeTermsIds); h++ {
// 				var i string = activeTermsIds[h]
// 				var term compiler.Term = a[i]

// 				// checking for operators

// 				if term.Type == "O" {
// 					switch term.Operator {
// 					case "+":
// 						fmt.Print(h)
// 						var value float64 = add(a[activeTermsIds[h-1]], a[activeTermsIds[h+1]], x, y)

// 						activeTermsIds[h] = "D" // D for done
// 						activeTermsIds[h+1] = "D"

// 						t := a[activeTermsIds[h-1]]
// 						var newTerm compiler.Term = compiler.Term{
// 							Constant:     t.Constant,
// 							Variable:     t.Variable,
// 							Exponent:     t.Exponent,
// 							ExponentTerm: t.ExponentTerm,
// 							Type:         t.Type,
// 							Operator:     t.Operator,
// 							Subterm:      t.Subterm,
// 							ID:           t.ID,
// 							Value:        value, // added the value of addition in here
// 						}
// 						a[activeTermsIds[h-1]] = newTerm
// 						fmt.Print("for x:")
// 						fmt.Print(x)
// 						fmt.Print("y:")
// 						fmt.Print(y)
// 						fmt.Print("value is = ")
// 						fmt.Print(value)

// 					}
// 				}

// 			}
// 		}
// 	}

// }
func GetPoints(equation [2]map[string]compiler.Term) {

	var leftHandSide map[string]compiler.Term = equation[0]
	var RightHandSide map[string]compiler.Term = equation[1]

	// 1. Get the original order once
	var originalOrderLeft []string
	var originalOrderRight []string
	for i := range leftHandSide {
		originalOrderLeft = append(originalOrderLeft, i)
	}
	for i := range RightHandSide {
		originalOrderRight = append(originalOrderRight, i)
	}

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			// 2. Create a FRESH copy of IDs for every single (x, y) coordinate
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
