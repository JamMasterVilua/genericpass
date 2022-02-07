package genericpass

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPasswordFrom(t *testing.T) {
	data := `
localhost:5432:bobbies
*:5000:dogs
hosteroni:*:coppers
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

	_, err := PasswordFrom([]string{"telly", "5001"}, bytes.NewBufferString(data))
	if err == nil || err.Error() != fmt.Sprintf("found no key match for key %+v in the password file", []string{"telly", "5001"}) {
		t.Error(err)
	}
}
