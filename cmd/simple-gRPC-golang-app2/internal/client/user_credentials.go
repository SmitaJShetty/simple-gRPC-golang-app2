package main

import "context"

// UserCredential authenticator construct
type UserCredential struct {
	username string
	password string
}

// GetRequestMetadata returns meta data
func (u *UserCredential) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    u.username,
		"password": u.password,
	}, nil
}

//RequireTransportSecurity confirms if this option will be used
func (u *UserCredential) RequireTransportSecurity() bool {
	return true
}
