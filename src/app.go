package main

import (
	"math"
	"net/http"
    "github.com/gorilla/mux"
)

func calculo(w http.ResponseWriter, r *http.Request) {
    x := 0.001;
    for i := 0; i < 1000000; i++ {
        x = x + math.Sqrt(x)
    }
    w.Write([]byte("Code.education Rocks!"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", calculo)
    http.ListenAndServe(":8001", r)
}