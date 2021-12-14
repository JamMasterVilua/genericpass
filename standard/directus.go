package standard

import "github.com/jammasterz/genericpass"

type DirectusPasswordProvider struct {
	Host string
	Port string
}

var _ PasswordProvider = DirectusPasswordProvider{}

func (dpp DirectusPasswordProvider) GetPassword() (string, error) {
	return genericpass.Password(".directuspass", []string{dpp.Host, dpp.Port})
}
