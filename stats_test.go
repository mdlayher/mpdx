package mpdx

import (
	"reflect"
	"testing"
	"time"

	"github.com/fhs/gompd/mpd"
)

func TestClientStats(t *testing.T) {
	tests := []struct {
		name  string
		attrs mpd.Attrs
		s     *Stats
	}{
		{
			name: "no attributes",
			s:    &Stats{Attrs: mpd.Attrs{}},
		},
		{
			name: "unknown attributes",
			attrs: mpd.Attrs{
				"foo": "bar",
			},
			s: &Stats{Attrs: mpd.Attrs{
				"foo": "bar",
			}},
		},
		{
			name: "all attributes with some unknown",
			attrs: mpd.Attrs{
				"uptime":     "120",
				"playtime":   "30",
				"artists":    "10",
				"albums":     "20",
				"songs":      "30",
				"dbplaytime": "600",
				"dbupdate":   "1475107200",
				"foo":        "bar",
			},
			s: &Stats{
				Artists:    10,
				Albums:     20,
				Songs:      30,
				Uptime:     2 * time.Minute,
				DBPlayTime: 10 * time.Minute,
				DBUpdate:   time.Date(2016, time.September, 29, 0, 0, 0, 0, time.UTC),
				PlayTime:   30 * time.Second,
				Attrs: mpd.Attrs{
					"foo": "bar",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClient(tt.attrs, nil)
			s, err := c.Stats()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if want, got := tt.s.DBUpdate, s.DBUpdate; !want.Equal(got) {
				t.Fatalf("unexpected database update times:\n- want: %v\n-  got: %v",
					want, got)
			}

			// Zero out times after time-specific comparison
			tt.s.DBUpdate = time.Time{}
			s.DBUpdate = time.Time{}

			if want, got := tt.s, s; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected Stats:\n- want: %+v\n-  got: %+v",
					want, got)
			}
		})
	}
}
