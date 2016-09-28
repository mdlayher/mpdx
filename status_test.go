package mpdx

import (
	"reflect"
	"testing"
	"time"

	"github.com/fhs/gompd/mpd"
)

func TestClientStatus(t *testing.T) {
	tests := []struct {
		name  string
		attrs mpd.Attrs
		s     *Status
		ok    bool
	}{
		{
			name: "invalid time field format",
			attrs: mpd.Attrs{
				"time": "foo",
			},
		},
		{
			name: "invalid audio field format",
			attrs: mpd.Attrs{
				"audio": "foo",
			},
		},
		{
			name: "invalid state",
			attrs: mpd.Attrs{
				"state": "foo",
			},
		},
		{
			name: "invalid boolean value",
			attrs: mpd.Attrs{
				"repeat": "false",
			},
		},
		{
			name: "no attributes",
			s:    &Status{Attrs: mpd.Attrs{}},
			ok:   true,
		},
		{
			name: "unknown attributes",
			attrs: mpd.Attrs{
				"foo": "bar",
			},
			s: &Status{Attrs: mpd.Attrs{
				"foo": "bar",
			}},
			ok: true,
		},
		{
			name: "some attributes with some unknown",
			attrs: mpd.Attrs{
				"volume": "100",
				"random": "1",
				"state":  "pause",
				"song":   "10",
				"foo":    "bar",
			},
			s: &Status{
				Volume: 100,
				Random: true,
				State:  StatePause,
				Song:   10,
				Attrs: mpd.Attrs{
					"foo": "bar",
				},
			},
			ok: true,
		},
		{
			name: "all attributes with some unknown",
			attrs: mpd.Attrs{
				"volume":         "100",
				"repeat":         "0",
				"random":         "1",
				"single":         "0",
				"consume":        "0",
				"playlist":       "20",
				"playlistlength": "30",
				"state":          "pause",
				"song":           "10",
				"songid":         "40",
				"nextsong":       "50",
				"nextsongid":     "60",
				"time":           "10:60",
				"elapsed":        "10.240",
				"duration":       "59.820",
				"bitrate":        "1024",
				"xfade":          "0",
				"audio":          "44100:16:2",
				"updating_db":    "1",
				"error":          "some error",
				"foo":            "bar",
			},
			s: &Status{
				Volume:         100,
				Repeat:         false,
				Random:         true,
				Single:         false,
				Consume:        false,
				Playlist:       20,
				PlaylistLength: 30,
				State:          StatePause,
				Song:           10,
				SongID:         40,
				NextSong:       50,
				NextSongID:     60,
				CurrentTime:    10 * time.Second,
				TotalTime:      1 * time.Minute,
				Elapsed:        10*time.Second + 240*time.Millisecond,
				Duration:       59*time.Second + 820*time.Millisecond,
				Bitrate:        1024,
				Crossfade:      0 * time.Second,
				SampleRate:     44100,
				BitDepth:       16,
				Channels:       2,
				UpdatingDB:     1,
				Error:          "some error",
				Attrs: mpd.Attrs{
					"foo": "bar",
				},
			},
			ok: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClient(tt.attrs, nil)
			s, err := c.Status()

			if err != nil && tt.ok {
				t.Fatalf("unexpected error: %v", err)
			}
			if err == nil && !tt.ok {
				t.Fatal("no error occurred, but expected an error")
			}

			if want, got := tt.s, s; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Status:\n- want: %+v\n-  got: %+v",
					want, got)
			}
		})
	}
}
