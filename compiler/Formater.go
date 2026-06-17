package compiler

import (
	"fmt"
	"strings"
)

func Format(s string) string {
	var Arrystring []string = strings.Split(s, "")
	var Variables []string = strings.Split("qwertyuiopasdfghjklzxcvbnm", "") // all the variables that we support
	var Constants []string = strings.Split("0123456789", "")
	var ansString string = "("

	for i := 0; i < len(s); i++ {
		var char string = Arrystring[i]
		// checking for variables
		for h := 0; h < len(Constants); h++ {
			if Constants[h] == char {
				// fmt.Println("caharcter " + char + "was considered a number")
				var ConstantOfTheTerm = ""
				ConstantOfTheTerm += char
				var endCharIndex = 0
				// if we get an constant then check if there is another num further 10 digits
				for j := 1; j <= 11; j++ {
					if i+j < len(Arrystring) {
						var char2 string = Arrystring[i+j]
						var charWasFound bool = false
						for k := 0; k < len(Constants); k++ {
							if Constants[k] == char2 {
								ConstantOfTheTerm += char2
								charWasFound = true
								endCharIndex = j + 1
								break
							}
						}
						if charWasFound == false {
							break
						}
					}
				}

				fmt.Println(ConstantOfTheTerm)
				i += endCharIndex
			}
		}

		char = Arrystring[i]

		for h := 0; h < len(Variables); h++ {
			if Variables[h] == char {
				// fmt.Println("caharcter " + char + "was considered a number")
				var ConstantOfTheTerm = ""
				ConstantOfTheTerm += char
				var endCharIndex = 0
				// if we get an constant then check if there is another num further 10 digits
				for j := 1; j <= 11; j++ {
					if i+j < len(Arrystring) {
						var char2 string = Arrystring[i+j]
						var charWasFound bool = false
						for k := 0; k < len(Variables); k++ {
							if Variables[k] == char2 {
								ConstantOfTheTerm += char2
								charWasFound = true
								endCharIndex = j + 1
								break
							}
						}
						if charWasFound == false {
							break
						}
					}
				}

				fmt.Println(ConstantOfTheTerm)
				i += endCharIndex
			}
		}

	}
	return (ansString)
}
