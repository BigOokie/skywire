package client

import (
	"github.com/google/uuid"

	"github.com/skycoin/skywire/pkg/cipher"
	"github.com/skycoin/skywire/pkg/routing"
)

// MockClient implements mock route finder client.
type mockClient struct {
	err error
}

// NewMock constructs a new mock Client.
func NewMock() Client {
	return &mockClient{}
}

// SetError assigns error that will be return on the next call to a
// public method.
func (r *mockClient) SetError(err error) {
	r.err = err
}

// PairedRoutes implements Clien for MockClient
func (r *mockClient) PairedRoutes(src, dst cipher.PubKey, minHops, maxHops uint16) ([]routing.Route, []routing.Route, error) {
	if r.err != nil {
		return nil, nil, r.err
	}

	return []routing.Route{
			{
				&routing.Hop{
					From:      src,
					To:        dst,
					Transport: uuid.New(),
				},
			},
		}, []routing.Route{
			{
				&routing.Hop{
					From:      src,
					To:        dst,
					Transport: uuid.New(),
				},
			},
		}, nil
}