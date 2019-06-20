//go:generate enumer -type=Antialiasing -trimprefix=Antialiasing -transform=snake

package mapcss

type Antialiasing int

const (
	AntialiasingFull Antialiasing = iota
	AntialiasingText
	AntialiasingNone
)
