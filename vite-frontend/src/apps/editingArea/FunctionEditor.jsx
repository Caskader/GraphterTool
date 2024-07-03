import React, { useEffect, useState } from 'react'


class Terms {
  constructor(){
    self.term = []
  }
}

// boilerplate for function (F(x))
const testfunc = {
  "metadata": {},
  "terms": [
    [{
      "type": "simplified",
      "variable": "x",
      "constant": "10",
      "after": ""
    }],
    [
      {
        "type": "simplified",
        "variable": "y",
        "constant": "1",
        "after": ""
      }

    ]
  ]
}

class FunctionHandler { // class for handleing functions
  constructor(h) {
    this.h = h
    this.variables = "qwertyuiopasdfghjklzxcvbnm"
  }

  parse(text) { // method to parse functions (string) to an acceptable format
    var Func = {
      "metadata": {

      },
      "terms": [
        [[], []]
      ]
    }
    const operators = "+-*/"
    const numbers = '1234567890=' // all valid numers
    const alphabets = "abcdefghejklmnopqrstuvwxyz" // all valid alphabets
    const numArr = numbers.split("") // spliting numbers for looping
    let constant = '' // ariable for storing the constant
    var raw = text.split("")

    var term = {}

    term = {
      "type": "simplified",
      "variable": "",
      "constant": "",
      "after": ""
    }
    let terms = []
    let equation = [terms, terms]
    let t = 0 // slector for the termbefore and after in an equation
    let termno = 0


    for (let i = 0; i < raw.length; i++) { // for loop for all the slices
      const element = raw[i];

      if (raw[i - 1] == "(") {
        term = new Terms()
      }
        for (let a = 0; a < numArr.length; a++) {
          const v = numArr[a];

          if (v == element) {
            if (v == "=") {
              equation[t] = terms
              t = 1
              console.log("changed constant to zero")
              
            }
            else if (raw[i - 1] != "^") {
              constant += v;
            }
          }
        }

        for (let a = 0; a < alphabets.length; a++) {
          const v = alphabets[a];

          if (element == v) {
            term["constant"] = constant
            term['n'] = v
            terms[termno] =
            {
              "type": "simplified",
              "variable": element,
              "constant": constant,
            }
            // termno += 1
            constant = ""
          }

        }


        // numArr.forEach((v, a, arr) => {
        //   if (v == element) {
        //     if (raw[i - 1] != "^") {
        //       term['constant'] = "1"
        //       return term
        //     }
        //     return term
        //   }
        //   return term
        // })
    }
    console.log(equation)
  }

  render(func) {
    var output = ""

    for (const key in this.func['terms']) {
      const element = this.func['terms'][key];

      if (key != 0) {
        output += " = "
      }
      for (const i in element) {

        output += element[i]['constant']
        output += element[i]['variable']
        output += " "
        output += element[i]['after']
        output += " "
      }
    }

    return output
  }
}

class SocketHandler{
  constructor(){
    this.url ="ws://localhost:2000/w"
    this.webSocket = new WebSocket(this.url);
    this.open = false
    this.webSocket.onopen = (e)=>{
      this.open = true
    }
  }
  send(msg){
    this.webSocket.onopen = (e)=>{
      this.webSocket.send("(10x)^1  = (2y)^1")
    }
  }
  get(){
    this.webSocket.onmessage = (e)=>{
      console.log(e)
    }
  }
}

export const FunctionEditor = () => {
  let c = new FunctionHandler(10)

  let h,seth = useState("")

  useEffect(() => {
    let f = new SocketHandler()
    f.send("hi")
    seth("2")
  }, [])


  return (
    <div id='Function-Editor' >
      {h}
    </div>
  )
}


