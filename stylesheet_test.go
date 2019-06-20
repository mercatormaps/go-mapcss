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
		name := tt.value
		if name == "" {
			name = "empty"
		}

		t.Run(name, func(t *testing.T) {
			s, err := mapcss.Parse(strings.NewReader(fmt.Sprintf(`
				canvas {
					antialiasing: %s;
				}
			`, tt.value)))

			require.NoError(t, err)
			require.Equal(t, tt.expected, s.Canvas.Antialiasing)
		})
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
		name := tt.value
		if name == "" {
			name = "empty"
		}

		t.Run(name, func(t *testing.T) {
			s, err := mapcss.Parse(strings.NewReader(fmt.Sprintf(`
				canvas {
					fill-opacity: %s;
				}
			`, tt.value)))

			if tt.ok {
				require.NoError(t, err)
				require.Equal(t, tt.expected, s.Canvas.FillOpacity)
			} else {
				require.Error(t, err)
			}
		})
	}
}
