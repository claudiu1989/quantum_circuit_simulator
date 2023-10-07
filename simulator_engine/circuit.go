package simulatorengine

type QuantumGate struct {
	Qubits             []int
	BasisStatesActions map[string]map[string]complex128
}

// Method that simulates a quantum gate applied to all input qubits
func (g QuantumGate) ApplyGateContiguousQubits(input Qudit) Qudit {
	output_amplitudes := make(map[string]complex128)
	for basis_state, amplitude := range input.Amplitudes {
		// Accumulate the contributions to the output state
		// for the current input basis state
		for contribution_basis_state, contribution_amplitude := range g.BasisStatesActions[basis_state] {
			_, exists := output_amplitudes[contribution_basis_state]
			if exists {
				output_amplitudes[contribution_basis_state] += amplitude * contribution_amplitude
			} else {
				output_amplitudes[contribution_basis_state] = amplitude * contribution_amplitude
			}
		}
	}
	output := Qudit{N_qubits: input.N_qubits, Amplitudes: output_amplitudes}
	return output
}

// Method that simulates a quantum gate applied to a subset of the
// input qubits
func (g QuantumGate) ApplyGate(input Qudit) Qudit {
	output_amplitudes := make(map[string]complex128)
	for basis_state, amplitude := range input.Amplitudes {
		// Accumulate the contributions to the output state
		// for the current input basis state
		for contribution_basis_state, contribution_amplitude := range g.BasisStatesActions[basis_state] {
			_, exists := output_amplitudes[contribution_basis_state]
			if exists {
				output_amplitudes[contribution_basis_state] += amplitude * contribution_amplitude
			} else {
				output_amplitudes[contribution_basis_state] = amplitude * contribution_amplitude
			}
		}
	}
	output := Qudit{N_qubits: input.N_qubits, Amplitudes: output_amplitudes}
	return output
}

/*
func getStateOnSpecifiedQubits(qubits_indices []int, full_state Qudit) Qudit {

		for basis_state, amplitude := range full_state.Amplitudes {

		}

		return full_state
	}
*/

func getStateOnQubitsSubset(input_state string, qubit_indices []int) string {
	partial_state := ""
	for _, qubit_index := range qubit_indices {
		partial_state += input_state[qubit_index : qubit_index+1]
	}
	return partial_state
}

type QuantumCircuit struct {
	Gates []QuantumGate
}

func ApplyCircuit(input Qudit) Qudit {
	output := Qudit{input.N_qubits, input.Amplitudes}
	return output
}
