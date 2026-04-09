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
func GetPoints(a map[string]compiler.Term) {
	// 1. Get the original order once
	var originalOrder []string
	for i := range a {
		originalOrder = append(originalOrder, i)
	}

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			// 2. Create a FRESH copy of IDs for every single (x, y) coordinate
			activeTermsIds := make([]string, len(originalOrder))
			copy(activeTermsIds, originalOrder)

			for h := 0; h < len(activeTermsIds); h++ {
				currID := activeTermsIds[h]
				if currID == "D" {
					continue
				} // Skip already processed terms

				term := a[currID]

				if term.Type == "O" {

					switch term.Operator {
					case "+":
						// Boundary check to prevent crashes
						if h > 0 && h < len(activeTermsIds)-1 {
							leftIdx := h - 1
							rightIdx := h + 1

							val := add(a[activeTermsIds[leftIdx]], a[activeTermsIds[rightIdx]], x, y)

							// Update the left term with the new calculated value
							t := a[activeTermsIds[leftIdx]]
							t.Value = val
							a[activeTermsIds[leftIdx]] = t

							// Mark operator and right operand as Done
							activeTermsIds[h] = "D"
							activeTermsIds[rightIdx] = "D"

							fmt.Printf("At (%d,%d) Value: %f\n", x, y, val)
						}
					}
				}
			}
		}
	}
}
