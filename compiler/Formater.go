package compiler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// get exponent when on ^
func getExponent(char string, Arrystring []string, i *int) (string, error) {
	var ConstantOfTheTerm string = ""

	for h := 0; h < 3; h++ {
		if *i+h < len(Arrystring) {
			if Arrystring[*i+h] == "^" {
				char = Arrystring[*i+h]
				*i += h
				break
			}
		} else {
			break
		}
	}

	var endCharIndex int = 0
	if "^" == char {
		char2 := Arrystring[*i+1]
		if isInteger(char2) {
			ConstantOfTheTerm = char2
			endCharIndex += 1
		} else if char2 == "{" {
			var a string = ""
			for n := 1; n < 100; n++ {
				if (n + *i) < len(Arrystring) {
					char3 := Arrystring[*i+n+1]

					if char3 == "}" {
						endCharIndex += n + 1
						break
					} else {
						a += char3
					}
				}
			}
			if a != "" {
				ConstantOfTheTerm += "( " + Format(a) + ")^1 "

			} else {
				// sometimes only giving a blank space in the exponent give a raw text ^{} which should be converted to ^1
				ConstantOfTheTerm += "1"
			}
		} else {
			return "", errors.New("exponent cannot be resolved ")
		}

		// fmt.Println("caharcter " + char + "was considered a number")

		// if we get an constant then check if there is another num further 10 digits
		*i += endCharIndex
	} else {
		return "1", nil
	}

	return ConstantOfTheTerm, nil
}

func getConstants(char string, ConstantOfTheTerm string, Arrystring []string, i int) (int, string) {
	var Constants []string = strings.Split("0123456789", "")
	ansString := ""
	var endCharIndex int = 0
	for h := 0; h < len(Constants); h++ {

		if Constants[h] == char {
			// fmt.Println("caharcter " + char + "was considered a number")

			ConstantOfTheTerm += char
			endCharIndex = 1
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

			ansString = ansString + " (" + ConstantOfTheTerm
			i += endCharIndex
		}
	}
	return i, ConstantOfTheTerm
}

func getVariables(char string, ConstantOfTheTerm string, Arrystring []string, i int) (int, string) {
	var Constants []string = strings.Split("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM", "")
	ansString := ""
	for h := 0; h < len(Constants); h++ {

		if Constants[h] == char {
			// fmt.Println("caharcter " + char + "was considered a number")

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

			ansString = ansString + " (" + ConstantOfTheTerm
			i += endCharIndex
		}
	}
	return i, ConstantOfTheTerm
}

func Format(s string) string {
	var Arrystring []string = strings.Split(s, "")
	var Variables []string = strings.Split("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM", "") // all the variables that we support
	var Operations []string = strings.Split("+-/*=", "")
	var ansString string = ""

	for i := 0; i < len(s); i++ {
		var char string = Arrystring[i]
		var ConstantOfTheTerm = ""
		//TODO: add power operations
		fmt.Println(char)
		if char == "\\" {
			fmt.Print("hi")
			var c string = ""
			i, c = getVariables(Arrystring[i+1], "", Arrystring, i+1)
			if c == "" {
				// if we didnt find any command after the slash just ignore the slash
				i--
			} else if c == "right" {
				ansString += ")^1 "
			} else if c == "left" {
				ansString += "("
			}

		}

		// checking for Constants
		i, ConstantOfTheTerm = getConstants(char, ConstantOfTheTerm, Arrystring, i)
		if ConstantOfTheTerm != "" {
			ansString += "(" + ConstantOfTheTerm

		}
		if i < len(Arrystring) {
			char = Arrystring[i]
		}
		//TODO: add 1 as constant in a variable x is same as 1x
		VariableOfTheTerm := ""
		for h := 0; h < len(Variables); h++ {
			if Variables[h] == char {
				// fmt.Println("caharcter " + char + "was considered a number")

				VariableOfTheTerm += char
				var endCharIndex = 0
				// if we get an constant then check if there is another num further 10 digits
				for j := 1; j <= 11; j++ {
					if i+j < len(Arrystring) {
						var char2 string = Arrystring[i+j]
						var charWasFound bool = false
						for k := 0; k < len(Variables); k++ {
							if Variables[k] == char2 {
								VariableOfTheTerm += char2
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

				exp, err := getExponent(char, Arrystring, &i)

				if err != nil {
					fmt.Println("WARNING: ", err)
				}

				if ConstantOfTheTerm != "" {
					ansString = ansString + VariableOfTheTerm + ")^" + exp + " "
				} else {
					ansString = ansString + " (1" + VariableOfTheTerm + ")^" + exp + " "
				}
				i += endCharIndex
			}
		}
		if ConstantOfTheTerm != "" {
			if VariableOfTheTerm == "" {
				exp, err := getExponent(char, Arrystring, &i)

				if err != nil {
					fmt.Println("Warning: ", err)
				}

				ansString += ")^" + exp + " "
			}
		}

		for h := 0; h < len(Operations); h++ {
			if char == Operations[h] {
				if char == "=" {
					ansString += " " + char + " "

				} else {
					ansString += " (" + char + ")^1 "

				}

				break
			}
		}

	}
	return (ansString)
}
