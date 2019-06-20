//go:generate enumer -type=Antialiasing -trimprefix=Antialiasing -transform=snake

package property

type Antialiasing int

const (
	AntialiasingFull Antialiasing = iota
	AntialiasingText
	AntialiasingNone
)
