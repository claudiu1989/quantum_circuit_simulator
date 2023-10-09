package simulatorengine

import (
	"testing"
)

func TestParseGetInputQudit(t *testing.T) {
	file_lines := []string{"Input", "1|0>+0|1>"}
	expected_input_qudit := Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(1, 0), "1": complex(0, 0)}}
	input_qudit := parseGetInputQudit(file_lines)
	if input_qudit.N_qubits != expected_input_qudit.N_qubits {
		t.Errorf("The number of qubits is %d instead of %d", input_qudit.N_qubits, expected_input_qudit.N_qubits)
	}
	for input_state, input_amplitude := range input_qudit.Amplitudes {
		if input_amplitude != expected_input_qudit.Amplitudes[input_state] {
			t.Errorf("The amplitude for state %s is %g instead of %g", input_state, input_amplitude, expected_input_qudit.Amplitudes[input_state])
		}
	}

}
