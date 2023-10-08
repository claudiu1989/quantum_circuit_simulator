package simulatorengine

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(file_path string) []string {
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
