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
