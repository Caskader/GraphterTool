package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

// in this algorithm all entities including operators are counted as terms
// each term has a constant(0,1,2,3,4,...) variable(x,y,z,...) exponent(if it is a single no ) exponentTerm (if it is another term)
type Term struct {
	constant     string
	variable     string
	exponent     string
	exponentTerm string // the exponent terms is actually an address of the term in TermContainer
	Type         string // can be "O" as operator or "N" as Normal term
	operator     string
	subterm      []string
	ID           string // address of the term in termContainer
}

var TermContainerBefore = map[string]Term{} // a map to store all the terms
var TermContainerAfter = map[string]Term{}

// function to get terms from plain string
func GetTerm(i int, raw []string, idPointer uint16, prefix string, TermContainer map[string]Term) (uint16, map[string]Term, int, string) {
	var nums [10]string = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var alphaets [2]string = [2]string{"x", "y"}
	var operators [4]string = [4]string{"/", "*", "+", "-"}
	var Type = "N"

	h := i + 1
	var constant string = ""
	var variable string = ""
	var expo string = ""
	var op string = "" //operation
	var nidPointer uint16 = 0
	var subtermsId []string = []string{}
	for true {
		a := raw[h]
		// var subTermPointer uint16 = 0
		for m := 0; m < 10; m++ {
			if raw[h] == "(" {
				var id string = ""
				nidPointer, TermContainer, h, id = GetTerm(h, raw, nidPointer, ".0", TermContainer)
				subtermsId = append(subtermsId, id)
			} else if raw[h] == ")" {
				expo = raw[h+2] // get the value after two places of the term
				if expo == "(" {
					expo = "Feature not available" // currently the feature of giving a term as an exponent is not availale
				}
				break // end the task
			}
		}
		if a == ")" { // this means that the current term has ended
			expo = raw[h+2] // get the value after two places of the term
			if expo == "(" {
				expo = "Feature not available" // currently the feature of giving a term as an exponent is not availale
			}
			break // end the task
		}

		// checking for constants
		var num string = ""
		for x := 0; x < len(nums); x++ {
			num = nums[x]
			if a == num {
				constant += num
			}
		}

		var j string = ""
		// checking for variables
		for y := 0; y < len(alphaets); y++ {
			j = alphaets[y]
			if a == j {
				variable = a
			}

		}

		// checking for operators
		var k string = ""
		for c := 0; c < len(operators); c++ {
			k = operators[c]
			if a == k {
				Type = "O"
				op = a
			}
		}

		if h >= 100 {
			break
		}
		h++
	}
	var t Term = Term{
		constant:     constant,
		variable:     variable,
		exponent:     expo,
		exponentTerm: "",
		Type:         Type,
		operator:     op,
		subterm:      subtermsId,
		ID:           prefix + "." + strconv.Itoa(int(idPointer)),
	}
	// return t
	//
	TermContainer[prefix+"."+strconv.Itoa(int(idPointer))] = t
	// fmt.Println(string(idPointer))
	idPointer += 1
	return idPointer, TermContainer, h + 2, t.ID // skip 3 as we have to skip the exponent too
}

func Parse(y string) [2]map[string]Term {
	var raw []string = strings.Split(y, "")
	var idPionter uint16 = 0
	var gotEqualtoSign bool = false
	var equation [2]map[string]Term = [2]map[string]Term{}
	for i := 0; i < len(raw); i++ {
		var element = raw[i]
		if element == "=" {
			idPionter = 0
			gotEqualtoSign = true
		}
		if element == "(" {
			if !gotEqualtoSign {
				idPionter, TermContainerBefore, i, _ = GetTerm(i, raw, idPionter, "", TermContainerBefore)

			} else if gotEqualtoSign {
				idPionter, TermContainerAfter, i, _ = GetTerm(i, raw, idPionter, "", TermContainerAfter)
			}

		}
	}
	equation = [2]map[string]Term{TermContainerBefore, TermContainerAfter}
	return equation

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	// data := map[string]interface{}{
	// 	"ni": 10,
	// }
	// jsonData, err := json.Marshal(data)

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		// var jsonMap map[string]interface{}
		// json.Unmarshal(msg, &jsonMap)

		// log.Printf("got message %s \n", msg)
		Parse(string(msg))
		err = conn.WriteMessage(websocket.TextMessage, []byte("got msg"))

		if err != nil {
			log.Println(err)
			break
		}

	}
}

func main() {
	// going to add serer connection when parsing is over
	// http.HandleFunc("/w", webSocketHandler)
	// log.Fatal(http.ListenAndServe(":2000", nil))

	fmt.Println(Parse("(10x)(+)^2(10y)^2"))
}
