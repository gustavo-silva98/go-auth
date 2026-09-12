package main

import (
	authclient "go-auth/internal/authClient"
)

type Backend struct {
	authClient authclient.AuthClient
}
