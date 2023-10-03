package simulatorengine

type QuantumGate struct {
	Qubits             []int
	BasisStatesActions map[int]map[int]complex128
}

// Method that simulates a quantum gate
func (g QuantumGate) ApplyGate(input Qudit) Qudit {
	output_amplitudes := make(map[int]complex128)
	for basis_state, amplitude := range input.Amplitudes {
		// Accumulate the contributions to the output state
		// for the current input basis state
		for contribution_basis_state, contribution_amplitude := range g.BasisStatesActions[basis_state] {
			cr_amplitude, exists := output_amplitudes[contribution_basis_state]
			if exists {
				output_amplitudes[contribution_basis_state] += amplitude * contribution_amplitude
			} else {
				output_amplitudes[contribution_basis_state] = cr_amplitude
			}
		}
	}
	output := Qudit{N_qubits: input.N_qubits, Amplitudes: output_amplitudes}
	return output
}

type QuantumCircuit struct {
	Gates []QuantumGate
}

func ApplyCircuit(input Qudit) Qudit {
	output := Qudit{input.N_qubits, input.Amplitudes}
	return output
}