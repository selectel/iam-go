package client

// AuthMethod is implemented by all authentication methods.
type AuthMethod interface {
	GetKeystoneToken() string
}

// KeystoneTokenAuth represents Keystone token authentication method.
// It conforms to AuthMethod interface.
type KeystoneTokenAuth struct {
	KeystoneToken string
}

func (k KeystoneTokenAuth) GetKeystoneToken() string {
	return k.KeystoneToken
}
