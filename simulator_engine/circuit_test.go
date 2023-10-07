package simulatorengine

import (
	"math"
	"math/cmplx"
	"testing"
)

const equalityThreshold = 1e-5

// TESTS FOR ApplyGateContiguousQubits

type testCaseApplyGateContiguousQubits struct {
	Input               Qudit
	Expected_amplitudes map[string]complex128
}

var testCasesApplyGateContiguousQubits = []testCaseApplyGateContiguousQubits{
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

func TestApplyGateContiguousQubits(t *testing.T) {
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

type testCaseForGetStateOnQubitsSubset struct {
	Input_state           string
	Qubits_indices        []int
	Expected_output_state string
}

var testCasesForGetStateOnQubitsSubset = []testCaseForGetStateOnQubitsSubset{
	{
		"00101",
		[]int{1, 2, 4},
		"011",
	},
	{
		"1",
		[]int{0},
		"1",
	},
	{
		"0010101000",
		[]int{2},
		"1",
	},
	{
		"011110110",
		[]int{0, 5, 8},
		"000",
	},
}

// TEST getStateOnQubitsSubset
func TestGetStateOnQubitsSubset(t *testing.T) {
	for _, test_case := range testCasesForGetStateOnQubitsSubset {
		partial_state := getStateOnQubitsSubset(test_case.Input_state, test_case.Qubits_indices)
		if partial_state != test_case.Expected_output_state {
			t.Errorf("The output state is %s, but the state %s was expected", partial_state, test_case.Expected_output_state)
		}
	}
}
