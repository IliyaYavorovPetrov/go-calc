package main

import (
	"encoding/json"
	"github.com/IliyaYavorovPetrov/go-calc/app/math"
	"io/ioutil"
	"net"
	"net/http"
)

func getNumbers(r *http.Request) []int {
	var numbers []int

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return numbers
	}

	err = json.Unmarshal(body, &numbers)
	if err != nil {
		return numbers
	}

	return numbers
}

func writeResult(w http.ResponseWriter, result int) {
	response := map[string]int{"result": result}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		return
	}
}

func AppRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Sum(getNumbers(r))
		writeResult(w, result)
	})

	mux.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Sub(getNumbers(r))
		writeResult(w, result)
	})
	
	mux.HandleFunc("/mul", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Mul(getNumbers(r))
		writeResult(w, result)
	})

	return mux
}


func main() {
	addr := net.JoinHostPort("", "8080")

	s := http.Server{
		Addr:    addr,
		Handler: AppRouter(),
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
