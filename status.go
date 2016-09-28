package mpdx

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fhs/gompd/mpd"
)

// A Status is the current status of an mpd player.
type Status struct {
	// Fields recognized by this package and parsed into correct types.
	Volume         int
	Repeat         bool
	Random         bool
	Single         bool
	Consume        bool
	Playlist       int
	PlaylistLength int
	State          State
	Song           int
	SongID         int
	NextSong       int
	NextSongID     int
	CurrentTime    time.Duration
	TotalTime      time.Duration
	Elapsed        time.Duration
	Duration       time.Duration
	Bitrate        int
	Crossfade      time.Duration
	SampleRate     int
	BitDepth       int
	Channels       int
	UpdatingDB     int
	Error          string

	// Fields which are not recognized by this package.
	Attrs mpd.Attrs

	// TODO(mdlayher): expose when better understood
	// MixrampdB      float64
	// MixrampDelay   float64
}

// NewStatus creates a Status from an input map of attributes.  All known
// attributes are parsed into Status fields, and any unknown attributes are
// copied into Status.Attrs.
func NewStatus(attrs mpd.Attrs) (*Status, error) {
	s := &Status{
		Attrs: make(mpd.Attrs, 0),
	}
	for k, v := range attrs {
		if err := s.parseAttr(k, v); err != nil {
			return nil, err
		}
	}

	return s, nil
}

// Raw key names returned by mpd for a status query.
const (
	volume         = "volume"
	repeat         = "repeat"
	random         = "random"
	single         = "single"
	consume        = "consume"
	playlist       = "playlist"
	playlistLength = "playlistlength"
	state          = "state"
	song           = "song"
	songID         = "songid"
	nextSong       = "nextsong"
	nextSongID     = "nextsongid"
	statusTime     = "time"
	elapsed        = "elapsed"
	duration       = "duration"
	bitrate        = "bitrate"
	xfade          = "xfade"
	mixrampdB      = "mixrampdb"
	mixrampDelay   = "mixrampdelay"
	audio          = "audio"
	updatingDB     = "updating_db"
	statusError    = "error"
)

// parseAttr parses a single attribute into a Status field using the
// input key and value.
func (s *Status) parseAttr(key string, value string) error {
	switch key {
	case repeat, random, single, consume:
		return s.parseBoolAttr(key, value)
	case volume, playlist, playlistLength, song, songID, nextSong, nextSongID, bitrate, updatingDB:
		return s.parseIntAttr(key, value)
	case statusTime, elapsed, duration, xfade:
		return s.parseDuration(key, value)
	case state:
		return s.parseState(value)
	case audio:
		return s.parseAudio(value)
	case statusError:
		s.Error = value
		return nil
	}

	// Attributes we aren't aware of are added to Attrs
	s.Attrs[key] = value

	return nil
}

// parseDuration parses a time.Duration attribute into a Status field
// using the input key and value.
func (s *Status) parseDuration(key string, value string) error {
	// All duration fields need a "s" suffix to be properly parsed
	// as a value in seconds by time.ParseDuration.

	if key == statusTime {
		// Field time is in format "1:20", where:
		//  -  1: current seconds elapsed
		//  - 20: total seconds in song
		vv := strings.Split(value, ":")
		if len(vv) != 2 {
			return fmt.Errorf("invalid %q field format: %q", key, value)
		}

		ct, err := time.ParseDuration(vv[0] + "s")
		if err != nil {
			return err
		}
		tt, err := time.ParseDuration(vv[1] + "s")
		if err != nil {
			return err
		}

		s.CurrentTime = ct
		s.TotalTime = tt
		return nil
	}

	d, err := time.ParseDuration(value + "s")
	if err != nil {
		return err
	}

	switch key {
	case elapsed:
		s.Elapsed = d
	case duration:
		s.Duration = d
	case xfade:
		s.Crossfade = d
	}

	return nil
}

// parseAudio parses the audio field into its three integer fields using
// the input value.
func (s *Status) parseAudio(value string) error {
	// Field audio is in format "44100:16:2", where:
	//  - 44100: sample rate in hertz
	//  -    16: bit depth
	//  -     2: number of channels
	vv := strings.Split(value, ":")
	if len(vv) != 3 {
		return fmt.Errorf("invalid audio field format: %q", value)
	}

	sample, err := strconv.Atoi(vv[0])
	if err != nil {
		return err
	}
	bits, err := strconv.Atoi(vv[1])
	if err != nil {
		return err
	}
	channels, err := strconv.Atoi(vv[2])
	if err != nil {
		return err
	}

	s.SampleRate = sample
	s.BitDepth = bits
	s.Channels = channels

	return nil
}

// parseState parses the state field into State, if value is valid.
func (s *Status) parseState(value string) error {
	switch State(value) {
	case StatePlay, StateStop, StatePause:
		s.State = State(value)
		return nil
	default:
		return fmt.Errorf("unknown state: %q", value)
	}
}

// parseIntAttr parses an integer attribute into a Status field using the
// input key and value.
func (s *Status) parseIntAttr(key string, value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	switch key {
	case volume:
		s.Volume = i
	case playlist:
		s.Playlist = i
	case playlistLength:
		s.PlaylistLength = i
	case song:
		s.Song = i
	case songID:
		s.SongID = i
	case nextSong:
		s.NextSong = i
	case nextSongID:
		s.NextSongID = i
	case bitrate:
		s.Bitrate = i
	case updatingDB:
		s.UpdatingDB = i
	}

	return nil
}

// pasreBoolAttr parses a boolean attribute into a Status field using the
// input key and value.
func (s *Status) parseBoolAttr(key string, value string) error {
	// Boolean values can only be 0 or 1.  No true/false values are
	// sent by mpd.
	switch value {
	case "0", "1":
		break
	default:
		return fmt.Errorf("value %q is not a valid boolean for key %q", value, key)
	}

	b, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}

	switch key {
	case repeat:
		s.Repeat = b
	case random:
		s.Random = b
	case single:
		s.Single = b
	case consume:
		s.Consume = b
	}

	return nil
}
