package simulatorengine

import (
	"math"
	"math/cmplx"
	"testing"
)

const equalityThreshold = 1e-5

type testApplyGateContiguousQubits struct {
	Input               Qudit
	Expected_amplitudes map[string]complex128
}

var testCasesApplyGateContiguousQubits = []testApplyGateContiguousQubits{
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(0, 0), "1": complex(1, 0)}},
		map[string]complex128{"0": complex(1, 0), "1": complex(0, 0)},
	},
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(1, 0), "1": complex(0, 0)}},
		map[string]complex128{"0": complex(1/math.Sqrt(2), 0), "1": complex(1/math.Sqrt(2), 0)},
	},
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(1/math.Sqrt(2), 0), "1": complex(1/math.Sqrt(2), 0)}},
		map[string]complex128{"0": complex(0.5+1/math.Sqrt(2), 0), "1": complex(0.5, 0)},
	},
}

func TestApplyGate(t *testing.T) {
	for _, test_case := range testCasesApplyGateContiguousQubits {
		input := test_case.Input
		expected_amplitudes := test_case.Expected_amplitudes

		amplitudes_0 := map[string]complex128{"0": complex(1/math.Sqrt(2), 0), "1": complex(1/math.Sqrt(2), 0)}
		amplitudes_1 := map[string]complex128{"0": complex(1, 0), "1": complex(0, 0)}
		gate_actions := map[string]map[string]complex128{"0": amplitudes_0, "1": amplitudes_1}
		qubits := make([]int, 1)
		qubits[0] = 0

		some_gate := QuantumGate{Qubits: qubits, BasisStatesActions: gate_actions}
		output := some_gate.ApplyGateContiguousQubits(input)
		if output.N_qubits != input.N_qubits {
			t.Errorf("The output has %d qubits instead of %d", output.N_qubits, input.N_qubits)
		}

		for basis_state, out_amplitude := range output.Amplitudes {
			if cmplx.Abs(out_amplitude-expected_amplitudes[basis_state]) > equalityThreshold {
				t.Errorf("For basis state %s, the amplitude is %g, instead of %g", basis_state, out_amplitude, expected_amplitudes[basis_state])
			}
		}
	}
}
