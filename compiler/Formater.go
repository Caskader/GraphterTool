	package compiler

import (
	"fmt"
	"strings"
)

func Format(s string) string {
	var Arrystring []string = strings.Split(s, "")
	var Variables [2]string = [2]string{"x", "y"}
	var ansString string = ""

	for i := 0; i < len(s); i++ {
		var char string = Arrystring[i]
		// checking for variables
		for h := 0; h < len(Variables); h++ {
			if Variables[h] == char {
				ansString += char
			}
		}
	}
	fmt.Print(ansString)
	return (ansString)
}
