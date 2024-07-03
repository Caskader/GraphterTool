package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"siddh.com/compiler"
	"siddh.com/graphter"
)

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
		fmt.Print(compiler.Parse((string(msg))))
		// err = conn.WriteMessage(websocket.TextMessage, msg)
		err = conn.WriteMessage(websocket.TextMessage, []byte("got it"))

		if err != nil {
			log.Println(err)
			break
		}

	}
}

func K(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html>hi</html>"))
}

func main() {
	equation := compiler.Parse("(10x)^1 (+)^1 (10)^1 = (10y)^2")
	graphter.GetPoints(equation[0])
}
