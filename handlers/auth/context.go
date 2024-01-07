package auth

import (
	"errors"
	"net/http"

	"github.com/gorilla/context"
)

type contextKey int

const (
	sessionServiceKey contextKey = 0
	clientKey         contextKey = 1
)

type OauthClient struct {
	Key         string `sql:"type:varchar(254);unique;not null"`
	Secret      string `sql:"type:varchar(60);not null"`
	RedirectURI string `sql:"type:varchar(200)"`
}

var (
	CLIENT = OauthClient{
		Key:         "client-key",
		Secret:      "client-secret",
		RedirectURI: "/",
	}
)

var (
	// ErrSessionServiceNotPresent ...
	ErrSessionServiceNotPresent = errors.New("Session service not present in the request context")
	// ErrClientNotPresent ...
	ErrClientNotPresent = errors.New("Client not present in the request context")
)

// Returns *oauth.Client from the request context
func getClient(r *http.Request) (*OauthClient, error) {
	val, ok := context.GetOk(r, clientKey)
	if !ok {
		return nil, ErrClientNotPresent
	}

	client, ok := val.(*OauthClient)
	if !ok {
		return nil, ErrClientNotPresent
	}

	return client, nil
}
