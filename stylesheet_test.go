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
		ok       bool
	}{
		{"none", mapcss.AntialiasingNone, true},
		{"text", mapcss.AntialiasingText, true},
		{"full", mapcss.AntialiasingFull, true},
		{"unknown", 0, false},
		{"", 0, false},
	}

	for _, tt := range tests {
		run(t, tt.value, func(t *testing.T, s *mapcss.Stylesheet, err error) {
			if tt.ok {
				require.NoError(t, err)
				require.Equal(t, &tt.expected, s.Canvas.Antialiasing)
			} else {
				require.Error(t, err)
			}
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
		{"0", 0, true},
		{"0.5", 0.5, true},
		{"1", 1, true},
		{"2", 0, false},
		{"-1", 0, false},
	}

	for _, tt := range tests {
		run(t, tt.value, func(t *testing.T, s *mapcss.Stylesheet, err error) {
			if tt.ok {
				require.NoError(t, err)
				require.Equal(t, &tt.expected, s.Canvas.FillOpacity)
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

func TestParseCanvasFillColor(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected mapcss.Color
		ok       bool
	}{
		{
			name:     "simple hex",
			value:    "#a1Ff09",
			expected: mapcss.Color{161, 255, 9, 1},
			ok:       true,
		},
		{
			name:     "hex with alpha",
			value:    "#a1Ff0933",
			expected: mapcss.Color{161, 255, 9, 0.2},
			ok:       true,
		},
		{
			name:     "shorthand hex",
			value:    "#f06",
			expected: mapcss.Color{255, 0, 102, 1},
			ok:       true,
		},
		{
			name:     "shorthand hex with alpha",
			value:    "#f063",
			expected: mapcss.Color{255, 0, 102, 0.2},
			ok:       true,
		},
		{
			name:     "rgb",
			value:    "rgb(126, 54, 213)",
			expected: mapcss.Color{126, 54, 213, 1},
			ok:       true,
		},
		{
			name:     "rgba",
			value:    "rgba(126, 54, 213, 0.8)",
			expected: mapcss.Color{126, 54, 213, 0.8},
			ok:       true,
		},
		{
			name:  "invalid rgb",
			value: "rgb(360, 54, 213)",
			ok:    false,
		},
		{
			name:  "invalid rgba",
			value: "rgba(126, 54, 213, 1.2)",
			ok:    false,
		},
	}

	for _, tt := range tests {
		run(t, tt.name, func(t *testing.T, s *mapcss.Stylesheet, err error) {
			if tt.ok {
				require.NoError(t, err)
				require.NotNil(t, s.Canvas.FillColor)
				require.Equal(t, tt.expected, *s.Canvas.FillColor)
			} else {
				require.Error(t, err)
			}
		}, `
			canvas {
				fill-color: %s;
			}
		`, tt.value)
	}
}

func TestParseRuleSelector(t *testing.T) {
	tests := []struct {
		name     string
		rule     string
		expected mapcss.Rule
	}{
		{
			name: "with type",
			rule: `node {}`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "node"},
				},
			},
		},
		{
			name: "multiple types",
			rule: `node, way, relation {}`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "node"},
					{Type: "way"},
					{Type: "relation"},
				},
			},
		},
		{
			name: "zoom range",
			rule: `way|z1-10 {}`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "way", Zoom: &mapcss.Zoom{
						Type: mapcss.ZoomRange, Min: 1, Max: 10,
					}},
				},
			},
		},
		{
			name: "minimum zoom",
			rule: `way|z5- {}`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "way", Zoom: &mapcss.Zoom{
						Type: mapcss.MinZoom, Min: 5,
					}},
				},
			},
		},
		{
			name: "exact zoom",
			rule: `way|z5 {}`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "way", Zoom: &mapcss.Zoom{
						Type: mapcss.ExactZoom, Min: 5, Max: 5,
					}},
				},
			},
		},
		{
			name: "selector with attribute",
			rule: `way[highway]`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "way", Attributes: []mapcss.Attribute{
						{Name: "highway"},
					}},
				},
			},
		},
		{
			name: "selector with negated attribute",
			rule: `way[!highway]`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "way", Attributes: []mapcss.Attribute{
						{Name: "highway", Not: true},
					}},
				},
			},
		},
		{
			name: "multiple attributes",
			rule: `node[place][POI]`,
			expected: mapcss.Rule{
				Selectors: []mapcss.Selector{
					{Type: "node", Attributes: []mapcss.Attribute{
						{Name: "place"},
						{Name: "POI"},
					}},
				},
			},
		},
	}

	for _, tt := range tests {
		run(t, tt.name, func(t *testing.T, s *mapcss.Stylesheet, err error) {
			require.Contains(t, s.Rules, tt.expected)
		}, tt.rule)
	}
}

func run(t *testing.T, name string, f func(*testing.T, *mapcss.Stylesheet, error), format string, args ...interface{}) {
	if name == "" {
		name = "empty"
	}

	t.Run(name, func(t *testing.T) {
		s, err := mapcss.Parse(strings.NewReader(fmt.Sprintf(format, args...)),
			mapcss.WithErrorReporter(func(msg string) {
				t.Log(msg)
			}))
		f(t, s, err)
	})
}
