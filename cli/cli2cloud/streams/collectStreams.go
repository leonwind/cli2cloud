package streams

import (
	"bufio"
	"os"
)

func readFromStreams(messages chan interface{}, f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		messages <- row
	}
	close(messages)
}

func CreateStreams(messages chan interface{}) {
	go readFromStreams(messages, os.Stdin)
	//go readFromStreams(messages, os.Stderr)
}
