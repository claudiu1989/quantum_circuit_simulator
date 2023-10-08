package simulatorengine

type QuantumCircuit struct {
	Gates []QuantumGate
}

// Method that simulates a quantum circuit
func (qc QuantumCircuit) ApplyCircuit(input Qudit) Qudit {
	qudit_after_gate := input
	for _, gate := range qc.Gates {
		qudit_after_gate = gate.ApplyGate(qudit_after_gate)
	}
	return qudit_after_gate
}
