package simulatorengine

import (
	"testing"
)

type testCaseParseGetInputQudit struct {
	FileLines          []string
	ExpectedInputQudit Qudit
}

var testCasesParseGetInputQudit = []testCaseParseGetInputQudit{
	{
		[]string{"Input", "1|0>+0|1>"},
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(1, 0), "1": complex(0, 0)}},
	},
	{
		[]string{"Input", "0.707|0>+0.707|1>"},
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(0.707, 0), "1": complex(0.707, 0)}},
	},
	{
		[]string{"Input", "0.707i|0>+0.707|1>"},
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(0, 0.707), "1": complex(0.707, 0)}},
	},
	{
		[]string{"Input", "0.707i|0>+0.707i|1>"},
		Qudit{N_qubits: 1, Amplitudes: map[string]complex128{"0": complex(0, 0.707), "1": complex(0, 0.707)}},
	},
	{
		[]string{"Input", "0.707i|001>+0.707i|011>"},
		Qudit{N_qubits: 3, Amplitudes: map[string]complex128{"001": complex(0, 0.707), "011": complex(0, 0.707)}},
	},
}

func TestParseGetInputQudit(t *testing.T) {
	for _, test_case := range testCasesParseGetInputQudit {
		input_qudit := parseGetInputQudit(test_case.FileLines)
		if input_qudit.N_qubits != test_case.ExpectedInputQudit.N_qubits {
			t.Errorf("The number of qubits is %d instead of %d", input_qudit.N_qubits, test_case.ExpectedInputQudit.N_qubits)
		}
		for input_state, input_amplitude := range input_qudit.Amplitudes {
			if input_amplitude != test_case.ExpectedInputQudit.Amplitudes[input_state] {
				t.Errorf("The amplitude for state %s is %g instead of %g", input_state, input_amplitude, test_case.ExpectedInputQudit.Amplitudes[input_state])
			}
		}
	}

}
