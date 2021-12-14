package genericpass

import (
	"bytes"
	"testing"
)

func TestPasswordFrom(t *testing.T) {
	data := `
localhost:5432:bobbies
*:5000:dogs
*:*:coppers
`

	tests := []struct {
		host, port, pass string
	}{
		{"localhost", "5432", "bobbies"},
		{"localhostess", "5000", "dogs"},
		{"hosteroni", "5001", "coppers"},
	}

	for _, tt := range tests {
		pass, err := PasswordFrom([]string{tt.host, tt.port}, bytes.NewBufferString(data))
		if err != nil {
			t.Error(tt, err)
		}
		if pass != tt.pass {
			t.Error(tt, "Wrong password: ", pass)
		}
	}
}
