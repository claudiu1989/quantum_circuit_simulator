package simulatorengine

import (
	"math"
	"testing"
)

func TestApplyGate(t *testing.T) {
	amplitudes_0 := map[int]complex128{0: complex(1/math.Sqrt(2), 0), 1: complex(1/math.Sqrt(2), 0)}
	amplitudes_1 := map[int]complex128{0: complex(1, 0), 1: complex(0, 0)}
	gate_actions := map[int]map[int]complex128{0: amplitudes_0, 1: amplitudes_1}
	some_gate := QuantumGate{N_qubits: 1, BasisStatesActions: gate_actions}
	input := Qudit{N_qubits: 1, Amplitudes: map[int]complex128{0: complex(0, 0), 1: complex(1, 0)}}
	output := some_gate.ApplyGate(input)
	if output.N_qubits != input.N_qubits {
		t.Errorf("The output has %d qubits instead of %d", output.N_qubits, input.N_qubits)
	}

	expected_amplitudes := map[int]complex128{0: complex(1, 0), 1: complex(0, 0)}
	for basis_state, out_amplitude := range output.Amplitudes {
		if out_amplitude != expected_amplitudes[basis_state] {
			t.Errorf("For basis state %d, the amplitude is %g, instead of %g", basis_state, out_amplitude, expected_amplitudes[basis_state])
		}
	}

}
