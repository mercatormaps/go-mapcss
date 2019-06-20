package mapcss_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mercatormaps/go-mapcss"

	"github.com/stretchr/testify/require"
)

func TestParseCanvasAntialiasing(t *testing.T) {
	tests := []struct {
		value    string
		expected mapcss.Antialiasing
	}{
		{"", mapcss.AntialiasingFull},
		{"none", mapcss.AntialiasingNone},
		{"text", mapcss.AntialiasingText},
		{"full", mapcss.AntialiasingFull},
	}

	for _, tt := range tests {
		run(t, tt.value, func(s *mapcss.Stylesheet, err error) {
			require.NoError(t, err)
			require.Equal(t, tt.expected, s.Canvas.Antialiasing)
		}, `
			canvas {
				antialiasing: %s;
			}
		`, tt.value)
	}
}

func TestParseCanvasFillOpacity(t *testing.T) {
	tests := []struct {
		value    string
		expected float32
		ok       bool
	}{
		{"", 1, true},
		{"0", 0, true},
		{"0.5", 0.5, true},
		{"1", 1, true},
		{"2", 0, false},
	}

	for _, tt := range tests {
		run(t, tt.value, func(s *mapcss.Stylesheet, err error) {
			if tt.ok {
				require.NoError(t, err)
				require.Equal(t, tt.expected, s.Canvas.FillOpacity)
			} else {
				require.Error(t, err)
			}
		}, `
			canvas {
				fill-opacity: %s;
			}
		`, tt.value)
	}
}

func run(t *testing.T, name string, f func(*mapcss.Stylesheet, error), format string, args ...interface{}) {
	if name == "" {
		name = "empty"
	}

	t.Run(name, func(t *testing.T) {
		s, err := mapcss.Parse(strings.NewReader(fmt.Sprintf(format, args...)))
		f(s, err)
	})
}
