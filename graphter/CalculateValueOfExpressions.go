package graphter

import (
	"siddh.com/compiler"
)

func calculateValueOfExpression(activeTermsIds []string, expression map[string]compiler.Term, x int, y int) float64 {
	var value float64
	var foundValue bool = false
	var FoundOperator bool = false

	for h := 0; h < len(activeTermsIds); h++ {
		currID := activeTermsIds[h]
		if currID == "D" {
			continue
		} // Skip already processed terms

		term := expression[currID]

		if term.Type == "O" {
			FoundOperator = true
			switch term.Operator {
			case "+":
				// Boundary check to prevent crashes
				if h > 0 && h < len(activeTermsIds)-1 {
					leftIdx := h - 1
					rightIdx := h + 1

					val := add(expression[activeTermsIds[leftIdx]], expression[activeTermsIds[rightIdx]], x, y, activeTermsIds, expression)

					// Update the left term with the new calculated value
					t := expression[activeTermsIds[leftIdx]]
					t.Value = val
					expression[activeTermsIds[leftIdx]] = t

					// Mark operator and right operand as Done
					activeTermsIds[h] = "D"
					activeTermsIds[rightIdx] = "D"

					value = val
					foundValue = true

				}
			case "-":
				// Boundary check to prevent crashes
				if h > 0 && h < len(activeTermsIds)-1 {
					leftIdx := h - 1
					rightIdx := h + 1

					val := subtract(expression[activeTermsIds[leftIdx]], expression[activeTermsIds[rightIdx]], x, y, activeTermsIds, expression)

					// Update the left term with the new calculated value
					t := expression[activeTermsIds[leftIdx]]
					t.Value = val
					expression[activeTermsIds[leftIdx]] = t

					// Mark operator and right operand as Done
					activeTermsIds[h] = "D"
					activeTermsIds[rightIdx] = "D"

					value = val
					foundValue = true

				}

			case "*":
				// Boundary check to prevent crashes
				if h > 0 && h < len(activeTermsIds)-1 {
					leftIdx := h - 1
					rightIdx := h + 1

					val := multiply(expression[activeTermsIds[leftIdx]], expression[activeTermsIds[rightIdx]], x, y, activeTermsIds, expression)

					// Update the left term with the new calculated value
					t := expression[activeTermsIds[leftIdx]]
					t.Value = val
					expression[activeTermsIds[leftIdx]] = t

					// Mark operator and right operand as Done
					activeTermsIds[h] = "D"
					activeTermsIds[rightIdx] = "D"

					value = val
					foundValue = true
				}
			case "/":
				// Boundary check to prevent crashes
				if h > 0 && h < len(activeTermsIds)-1 {
					leftIdx := h - 1
					rightIdx := h + 1

					val := devide(expression[activeTermsIds[leftIdx]], expression[activeTermsIds[rightIdx]], x, y, activeTermsIds, expression)

					// Update the left term with the new calculated value
					t := expression[activeTermsIds[leftIdx]]
					t.Value = val
					expression[activeTermsIds[leftIdx]] = t

					// Mark operator and right operand as Done
					activeTermsIds[h] = "D"
					activeTermsIds[rightIdx] = "D"

					value = val
					foundValue = true

				}
			}
		}
	}

	if foundValue == false {
	}

	if FoundOperator == false {
		for i := 0; i < len(activeTermsIds); i++ {
			value += GetComplexTermValue(expression[activeTermsIds[i]], x, y, activeTermsIds, expression)
		}
	}

	return value
}
