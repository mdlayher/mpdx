package mpdx

import (
	"strconv"
	"time"

	"github.com/fhs/gompd/mpd"
)

// Stats are statistics returned by an mpd server.
type Stats struct {
	// Fields recognized by this package and parsed into correct types.
	Artists    int
	Albums     int
	Songs      int
	Uptime     time.Duration
	DBPlayTime time.Duration
	DBUpdate   time.Time
	PlayTime   time.Duration

	// Fields which are not recognized by this package.
	Attrs mpd.Attrs
}

// NewStats creates a Stats from an input map of attributes.  All known
// attributes are parsed into Stats fields, and any unknown attributes are
// copied into Stats.Attrs.
func NewStats(attrs mpd.Attrs) (*Stats, error) {
	s := &Stats{
		Attrs: make(mpd.Attrs, 0),
	}
	for k, v := range attrs {
		if err := s.parseAttr(k, v); err != nil {
			return nil, err
		}
	}

	return s, nil
}

// Raw key names returned by mpd for a stats query.
const (
	uptime     = "uptime"
	playTime   = "playtime"
	artists    = "artists"
	albums     = "albums"
	songs      = "songs"
	dbPlayTime = "dbplaytime"
	dbUpdate   = "dbupdate"
)

// parseAttr parses a single attribute into a Stats field using the
// input key and value.
func (s *Stats) parseAttr(key string, value string) error {
	switch key {
	case uptime, playTime, dbPlayTime:
		return s.parseDurationAttr(key, value)
	case artists, albums, songs:
		return s.parseIntAttr(key, value)
	case dbUpdate:
		return s.parseTimeAttr(key, value)
	}

	// Attributes we aren't aware of are added to Attrs
	s.Attrs[key] = value

	return nil
}

// parseDurationAttr parses a time.Duration attribute into a Stats field
// using the input key and value.
func (s *Stats) parseDurationAttr(key string, value string) error {
	d, err := parseDurationSeconds(value)
	if err != nil {
		return err
	}

	switch key {
	case uptime:
		s.Uptime = d
	case playTime:
		s.PlayTime = d
	case dbPlayTime:
		s.DBPlayTime = d
	}

	return nil
}

// parseIntAttr parses an integer attribute into a Stats field using the
// input key and value.
func (s *Stats) parseIntAttr(key string, value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	switch key {
	case artists:
		s.Artists = i
	case albums:
		s.Albums = i
	case songs:
		s.Songs = i
	}

	return nil
}

// parseTimeAttr parses a time.Time attribute into a Stats field
// using the input key and value.
func (s *Stats) parseTimeAttr(key string, value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	t := time.Unix(int64(i), 0)

	switch key {
	case dbUpdate:
		s.DBUpdate = t
	}

	return nil
}
