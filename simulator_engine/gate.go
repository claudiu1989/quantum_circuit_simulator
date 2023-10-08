package simulatorengine

import "slices"

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
		partial_state := getStateOnQubitsSubset(basis_state, g.Qubits)
		// Apply the gate on the specified qubits
		partial_state_qudit := g.ApplyGateContiguousQubits(Qudit{N_qubits: len(g.Qubits), Amplitudes: map[string]complex128{partial_state: amplitude}})
		// Iterate through the base states defined on the gate's qubits
		for partial_basis_state, partial_amplitude := range partial_state_qudit.Amplitudes {
			new_basis_state := ""
			// Construct the states on all qubits, after the gate was applied
			for i := range basis_state {
				// The qubit state is either from the output of the gate, or the original basis state value
				if slices.Contains(g.Qubits, i) {
					index_in_partial_basis_state := slices.Index(g.Qubits, i)
					new_basis_state += partial_basis_state[index_in_partial_basis_state : index_in_partial_basis_state+1]
				} else {
					new_basis_state += basis_state[i : i+1]
				}
			}
			cr_amplitude, exists := output_amplitudes[new_basis_state]
			if exists {
				output_amplitudes[new_basis_state] = cr_amplitude + partial_amplitude
			} else {
				output_amplitudes[new_basis_state] = partial_amplitude
			}
		}

	}
	output := Qudit{N_qubits: input.N_qubits, Amplitudes: output_amplitudes}
	return output
}

// Get the state for a subset of qubits (qubits_indices),
// from an arbitrary state (input_state)
func getStateOnQubitsSubset(input_state string, qubits_indices []int) string {
	partial_state := ""
	for _, qubit_index := range qubits_indices {
		partial_state += input_state[qubit_index : qubit_index+1]
	}
	return partial_state
}
