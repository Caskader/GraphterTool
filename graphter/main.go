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

	// fmt.Println(value)
	return value
}

// function to add two terms in an equation
func add(firstTerm compiler.Term, secondTerm compiler.Term, x int, y int) float64 {
	a1 := GetTermValue(firstTerm, x, y)
	a2 := GetTermValue(secondTerm, x, y)

	return a1 + a2
}

func GetPoints(a map[string]compiler.Term) {

	var activeTerms []string = []string{}
	for i := range a {
		activeTerms = append(activeTerms, i)

	}
	var activeTermsCache []string = activeTerms

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			activeTerms = activeTermsCache
			for h := range activeTerms {
				var i string = activeTerms[h]
				var term compiler.Term = a[i]

				// checking for operators

				if term.Type == "O" {
					switch term.Operator {
					case "+":
						var value float64 = add(a[activeTerms[h-1]], a[activeTerms[h+1]], 10, 10)

						activeTerms[h] = "D" // D for done
						activeTerms[h+1] = "D"

						t := a[activeTerms[h-1]]
						var newTerm compiler.Term = compiler.Term{
							Constant:     t.Constant,
							Variable:     t.Variable,
							Exponent:     t.Exponent,
							ExponentTerm: t.ExponentTerm,
							Type:         t.Type,
							Operator:     t.Operator,
							Subterm:      t.Subterm,
							ID:           t.ID,
							Value:        value, // added the value of addition in here
						}
						a[activeTerms[h-1]] = newTerm
						fmt.Println(value)

					}
				}

			}
			fmt.Println(activeTerms)
		}
	}
	fmt.Println(a["0.1"].Value)

}
