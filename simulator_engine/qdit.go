package simulatorengine

type Qudit struct {
	N_qubits   int
	Amplitudes map[int]complex128
}
