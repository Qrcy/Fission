package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	X := [][]float32{
		[]float32{8, 3},
		[]float32{2, 4},
		[]float32{3, 6},
	}

	Y := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 6, 8},
	}

	out := multiply(X, Y)
	w.Write([]byte(fmt.Sprintf("v1: %b", out)))
}

func multiply(x, y [][]float32) ([][]float32) {

	out := make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return out
}
