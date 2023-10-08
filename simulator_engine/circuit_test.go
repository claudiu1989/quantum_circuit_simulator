package simulatorengine

import (
	"math"
	"math/cmplx"
	"testing"
)

// TESTS FOR ApplyCircuit

type testCaseApplyCircuit struct {
	Input               Qudit
	Expected_amplitudes map[string]complex128
}

var testCasesApplyCircuitOneGateOneQubit = []testCaseApplyCircuit{
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
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(1/math.Sqrt(2), 0), "1": complex(1/math.Sqrt(2), 0)}},
		map[string]complex128{"0": complex(0.5+1/math.Sqrt(2), 0), "1": complex(0.5, 0)},
	},
	{
		Qudit{N_qubits: 2, Amplitudes: map[string]complex128{"00": complex(1/math.Sqrt(2), 0), "11": complex(1/math.Sqrt(2), 0)}},
		map[string]complex128{"00": complex(0.5, 0), "01": complex(1/math.Sqrt(2), 0), "10": complex(0.5, 0)},
	},
	{
		Qudit{N_qubits: 3, Amplitudes: map[string]complex128{"011": complex(1/math.Sqrt(2), 0), "111": complex(1/math.Sqrt(2), 0)}},
		map[string]complex128{"011": complex(0.5+1/math.Sqrt(2), 0), "111": complex(0.5, 0)},
	},
}

func TestApplyCircuitOneGateOneQubit(t *testing.T) {
	for _, test_case := range testCasesApplyCircuitOneGateOneQubit {
		input := test_case.Input
		expected_amplitudes := test_case.Expected_amplitudes

		amplitudes_0 := map[string]complex128{"0": complex(1/math.Sqrt(2), 0), "1": complex(1/math.Sqrt(2), 0)}
		amplitudes_1 := map[string]complex128{"0": complex(1, 0), "1": complex(0, 0)}
		gate_actions := map[string]map[string]complex128{"0": amplitudes_0, "1": amplitudes_1}
		qubits := make([]int, 1)
		qubits[0] = 0

		some_gate := QuantumGate{Qubits: qubits, BasisStatesActions: gate_actions}
		some_circuit := QuantumCircuit{Gates: []QuantumGate{some_gate}}
		output := some_circuit.ApplyCircuit(input)
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

var testCasesApplyCircuitOneGateMultipleQubit = []testCaseApplyCircuit{
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"00": complex(1, 0)}},
		map[string]complex128{"00": complex(1, 0)},
	},
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"01": complex(1, 0)}},
		map[string]complex128{"01": complex(1, 0)},
	},
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"10": complex(1, 0)}},
		map[string]complex128{"11": complex(1, 0)},
	},
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"11": complex(1, 0)}},
		map[string]complex128{"10": complex(1, 0)},
	},
	{
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"11": complex(1/math.Sqrt(2), 0), "01": complex(1/math.Sqrt(2), 0)}},
		map[string]complex128{"10": complex(1/math.Sqrt(2), 0), "01": complex(1/math.Sqrt(2), 0)},
	},
}

// CNOT gate is used
func TestApplyCircuitOneGateMultipleQubit(t *testing.T) {
	for _, test_case := range testCasesApplyCircuitOneGateMultipleQubit {
		input := test_case.Input
		expected_amplitudes := test_case.Expected_amplitudes

		amplitudes_00 := map[string]complex128{"00": complex(1, 0)}
		amplitudes_01 := map[string]complex128{"01": complex(1, 0)}
		amplitudes_10 := map[string]complex128{"11": complex(1, 0)}
		amplitudes_11 := map[string]complex128{"10": complex(1, 0)}
		gate_actions := map[string]map[string]complex128{
			"00": amplitudes_00,
			"01": amplitudes_01,
			"10": amplitudes_10,
			"11": amplitudes_11}
		qubits := []int{0, 1}

		some_gate := QuantumGate{Qubits: qubits, BasisStatesActions: gate_actions}
		some_circuit := QuantumCircuit{Gates: []QuantumGate{some_gate}}
		output := some_circuit.ApplyCircuit(input)
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
