// Package mpdx is an extension of package mpd: github.com/fhs/gompd.
package mpdx

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
