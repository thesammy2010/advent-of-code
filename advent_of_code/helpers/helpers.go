package helpers

import (
	"bufio"
	"io/ioutil"
	"os"
)

type HashsetMap map[string]struct{}

type Hashset struct {
	set HashsetMap
}

func ReadFile(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func ReadFileLineByLine(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
