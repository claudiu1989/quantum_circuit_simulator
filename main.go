package main

import (
	"fmt"
	"math"

	simulatorengine "github.com/claudiu1989/quantum_circuit_simulator/simulator_engine"
)

func main() {
	amplitudes := map[string]complex128{"0": complex(1/math.Sqrt(2), 0), "1": complex(1/math.Sqrt(2), 0)}
	input_qudit := simulatorengine.Qudit{N_qubits: 2, Amplitudes: amplitudes}
	fmt.Println(input_qudit.Amplitudes)
	fmt.Println(input_qudit.N_qubits)
}
