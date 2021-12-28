package transport

import "net/http"

type ConnectOptions struct {
	Insecure     bool
	CustomClient *http.Client
}
