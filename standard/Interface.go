package standard

type PasswordProvider interface {
	GetPassword() (string, error)
}
