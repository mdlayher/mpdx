package mpdx

import (
	"github.com/fhs/gompd/mpd"
)

// A Client is a wrapped mpd.Client, which provides higher-level
// functionality and parsing of raw mpd attribute output.
//
// The underlying mpd.Client's Close method must be called to free
// up resources.  This Client does not provide a Close method.
type Client struct {
	c client
}

// New wraps an input mpd.Client in a Client.
func New(c *mpd.Client) *Client {
	return &Client{
		c: c,
	}
}

// Stats retrieves statistics from an mpd server.
func (c *Client) Stats() (*Stats, error) {
	attrs, err := c.c.Stats()
	if err != nil {
		return nil, err
	}

	return NewStats(attrs)
}

// Status retrieves the current status of an mpd server.
func (c *Client) Status() (*Status, error) {
	attrs, err := c.c.Status()
	if err != nil {
		return nil, err
	}

	return NewStatus(attrs)
}

// Ensure mpd.Client implements client.
var _ client = &mpd.Client{}

// A client is a client which can communicate with mpd or a mocked
// version of it.
type client interface {
	Stats() (mpd.Attrs, error)
	Status() (mpd.Attrs, error)
}
