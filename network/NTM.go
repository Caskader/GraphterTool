package network

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var PointerId int64 = 1

type ResponseData struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Id      int64  `json:"id"`
}

type EquationRequest struct {
	Equation      string `json:"equation"`
	StartingPoint [2]int `json:"StartingPoint"`
	EndingPoint   [2]int `json:"EndingPoint"`
}

type PointsResponse struct {
	Points [][2]float64 `json:"points"`
	Status string       `json:"status"`
	Error  string       `json:"error,omitempty"`
}

var HandleInput func(ResponseData)
var HandleEquation func(string, [2]int, [2]int) ([][2]float64, error)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "React Developer"
	}

	response := ResponseData{
		Message: fmt.Sprintf("Hello, %s! Go received your React GET request.", name),
		Status:  "Success",
		Id:      PointerId,
	}

	PointerId += 1
	if PointerId >= 100000 {
		PointerId = 0
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	HandleInput(response)
}

func equationHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PointsResponse{
			Status: "error",
			Error:  "Method not allowed",
			Points: [][2]float64{},
		})
		return
	}

	var req EquationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PointsResponse{
			Status: "error",
			Error:  "Invalid request body",
			Points: [][2]float64{},
		})
		return
	}

	if HandleEquation == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PointsResponse{
			Status: "error",
			Error:  "Equation handler not configured",
			Points: [][2]float64{},
		})
		return
	}

	points, err := HandleEquation(req.Equation, req.StartingPoint, req.EndingPoint)
	fmt.Println(req.StartingPoint, req.EndingPoint)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PointsResponse{
			Status: "error",
			Error:  err.Error(),
			Points: [][2]float64{},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PointsResponse{
		Status: "success",
		Points: points,
	})
}

func sendData(id int64) {
	//assuming that we checked the id

}

func Start(q func(ResponseData), eq func(string, [2]int, [2]int) ([][2]float64, error)) {
	http.HandleFunc("/api/data", dataHandler)
	http.HandleFunc("/api/equation", equationHandler)

	HandleInput = q
	HandleEquation = eq

	fmt.Println("Go API Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
