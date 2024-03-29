// Code generated by "enumer -type=Antialiasing -trimprefix=Antialiasing -transform=snake"; DO NOT EDIT.

//
package mapcss

import (
	"fmt"
)

const _AntialiasingName = "fulltextnone"

var _AntialiasingIndex = [...]uint8{0, 4, 8, 12}

func (i Antialiasing) String() string {
	if i < 0 || i >= Antialiasing(len(_AntialiasingIndex)-1) {
		return fmt.Sprintf("Antialiasing(%d)", i)
	}
	return _AntialiasingName[_AntialiasingIndex[i]:_AntialiasingIndex[i+1]]
}

var _AntialiasingValues = []Antialiasing{0, 1, 2}

var _AntialiasingNameToValueMap = map[string]Antialiasing{
	_AntialiasingName[0:4]:  0,
	_AntialiasingName[4:8]:  1,
	_AntialiasingName[8:12]: 2,
}

// AntialiasingString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AntialiasingString(s string) (Antialiasing, error) {
	if val, ok := _AntialiasingNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Antialiasing values", s)
}

// AntialiasingValues returns all values of the enum
func AntialiasingValues() []Antialiasing {
	return _AntialiasingValues
}

// IsAAntialiasing returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Antialiasing) IsAAntialiasing() bool {
	for _, v := range _AntialiasingValues {
		if i == v {
			return true
		}
	}
	return false
}
