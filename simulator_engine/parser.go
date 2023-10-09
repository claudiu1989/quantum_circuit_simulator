package simulatorengine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const specialSplitString = "***"

func readFile(file_path string) []string {
	read_file, err := os.Open(file_path)

	if err != nil {
		fmt.Println(err)
	}

	var file_lines []string
	file_scanner := bufio.NewScanner(read_file)
	file_scanner.Split(bufio.ScanLines)

	for file_scanner.Scan() {
		file_lines = append(file_lines, file_scanner.Text())
	}

	read_file.Close()
	return file_lines
}

func parseGetInputQudit(file_lines []string) Qudit {
	input_qudit_string := file_lines[1]
	var amplitudes map[string]complex128
	var augumente_qudit_sb strings.Builder
	for i, rune_qudit := range input_qudit_string {
		if i < len(input_qudit_string)-1 && (input_qudit_string[i+1:i+2] == "+" || input_qudit_string[i+1:i+2] == "-") {
			augumente_qudit_sb.WriteString(specialSplitString)
		}
		augumente_qudit_sb.WriteRune(rune_qudit)
	}
	var n_qubits int
	amplitude_state_pairs := strings.Split(augumente_qudit_sb.String(), specialSplitString)
	for _, cr_amplitude_state_pair := range amplitude_state_pairs {
		amplitude_state_arr := strings.Split(cr_amplitude_state_pair, "|")
		cr_amplitude, err := strconv.ParseComplex(amplitude_state_arr[0], 128)
		cr_state := amplitude_state_arr[1][:len(amplitude_state_arr[1])-1]
		n_qubits = len(cr_state)
		amplitudes[cr_state] = cr_amplitude
	}

	return Qudit{N_qubits: n_qubits, Amplitudes: amplitudes}
}

func ParseFile(file_path string) {
	file_lines := readFile(file_path)
	fmt.Println(file_lines)
}
