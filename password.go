package genericpass

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// PasswordFrom reads password for given host and user from r, which should
// be in a valid pgpass format. Host should be of the form "hostname:port".
func PasswordFrom(keyParts []string, r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")

		if checkKey(keyParts, splits) {
			return splits[len(splits)-1], nil
		}
	}

	return "", fmt.Errorf("no match found")
}

func checkKey(keyParts, splits []string) bool {
	if len(keyParts) != len(splits)-1 {
		return false
	}

	for index, part := range keyParts {
		if !eq(part, splits[index]) {
			return false
		}
	}

	return true
}

func eq(s, pattern string) bool {
	return pattern == "*" || s == pattern
}

// Password reads password for given host and user from a default pgpass file.
// Host should be of the form "hostname:port".
func Password(fileName string, keyParts []string) (string, error) {
	f, err := OpenDefault(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return PasswordFrom(keyParts, f)
}
