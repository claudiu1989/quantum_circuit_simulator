package simulatorengine

type QuantumGate struct {
	qubits       []int
	base_actions map[int]complex128
}

type QuantumCircuit struct {
	gates []QuantumGate
}

func ApplyCircuit(input Qudit) Qudit {
	output := Qudit{input.N_qubits, input.Amplitudes}
	return output
}
