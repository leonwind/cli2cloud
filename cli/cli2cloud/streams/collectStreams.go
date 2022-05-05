package streams

import (
	"bufio"
	"os"
)

func readFromStreams(messages chan string, f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		messages <- row
	}
}

func CreateStreams(messages chan string) {
	go readFromStreams(messages, os.Stdin)
	go readFromStreams(messages, os.Stderr)
}
