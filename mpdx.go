// Package mpdx is an extension of package mpd: github.com/fhs/gompd.
package mpdx

import (
	"time"
)

// A State is the state of the current audio track.
type State string

const (
	// StatePlay indicates that a track is currently playing.
	StatePlay State = "play"

	// StateStop indicates that a track is stopped.
	StateStop State = "stop"

	// StatePause indicates that a track is paused..
	StatePause State = "pause"
)

// parseDurationSeconds parses a string value in seconds into a
// time.Duration.
func parseDurationSeconds(value string) (time.Duration, error) {
	// All duration fields need a "s" suffix to be properly parsed
	// as a value in seconds by time.ParseDuration.
	return time.ParseDuration(value + "s")
}
