package mpdx

import (
	"github.com/fhs/gompd/mpd"
)

var _ client = &memoryClient{}

// A memoryClient is a client implementation which returns attributes
// stored in memory.
type memoryClient struct {
	attrs mpd.Attrs
	err   error
}

// testClient creates and configures a Client backed by in-memory attributes.
func testClient(attrs mpd.Attrs, err error) *Client {
	if attrs == nil {
		attrs = make(mpd.Attrs, 0)
	}

	return &Client{
		c: &memoryClient{
			attrs: attrs,
			err:   err,
		},
	}
}

func (c *memoryClient) Status() (mpd.Attrs, error) { return c.attrs, c.err }
