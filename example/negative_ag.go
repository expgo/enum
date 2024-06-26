// Code generated by https://github.com/expgo/ag DO NOT EDIT.
// Plugins:
//   - github.com/expgo/enum

package example

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// AllNegativeUnknown is an AllNegative of type Unknown.
	AllNegativeUnknown AllNegative = -5
	// AllNegativeGood is an AllNegative of type Good.
	AllNegativeGood AllNegative = -4
	// AllNegativeBad is an AllNegative of type Bad.
	AllNegativeBad AllNegative = -3
	// AllNegativeUgly is an AllNegative of type Ugly.
	AllNegativeUgly AllNegative = -2
)

const (
	// StatusUnknown is a Status of type Unknown.
	StatusUnknown Status = -1
	// StatusGood is a Status of type Good.
	StatusGood Status = 0
	// StatusBad is a Status of type Bad.
	StatusBad Status = 1
)

var ErrInvalidAllNegative = errors.New("not a valid AllNegative")

var _AllNegativeName = "UnknownGoodBadUgly"

var _AllNegativeMapName = map[AllNegative]string{
	AllNegativeUnknown: _AllNegativeName[0:7],
	AllNegativeGood:    _AllNegativeName[7:11],
	AllNegativeBad:     _AllNegativeName[11:14],
	AllNegativeUgly:    _AllNegativeName[14:18],
}

// Name is the attribute of AllNegative.
func (x AllNegative) Name() string {
	if v, ok := _AllNegativeMapName[x]; ok {
		return v
	}
	return fmt.Sprintf("AllNegative(%d).Name", x)
}

// Val is the attribute of AllNegative.
func (x AllNegative) Val() int {
	return int(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x AllNegative) IsValid() bool {
	_, ok := _AllNegativeMapName[x]
	return ok
}

// String implements the Stringer interface.
func (x AllNegative) String() string {
	return x.Name()
}

var _AllNegativeNameMap = map[string]AllNegative{
	_AllNegativeName[0:7]:                    AllNegativeUnknown,
	strings.ToLower(_AllNegativeName[0:7]):   AllNegativeUnknown,
	_AllNegativeName[7:11]:                   AllNegativeGood,
	strings.ToLower(_AllNegativeName[7:11]):  AllNegativeGood,
	_AllNegativeName[11:14]:                  AllNegativeBad,
	strings.ToLower(_AllNegativeName[11:14]): AllNegativeBad,
	_AllNegativeName[14:18]:                  AllNegativeUgly,
	strings.ToLower(_AllNegativeName[14:18]): AllNegativeUgly,
}

// ParseAllNegative converts a string to an AllNegative.
func ParseAllNegative(value string) (AllNegative, error) {
	if x, ok := _AllNegativeNameMap[value]; ok {
		return x, nil
	}
	if x, ok := _AllNegativeNameMap[strings.ToLower(value)]; ok {
		return x, nil
	}
	return AllNegative(0), fmt.Errorf("%s is %w", value, ErrInvalidAllNegative)
}

var ErrInvalidStatus = errors.New("not a valid Status")

var _StatusName = "UnknownGoodBad"

var _StatusMapName = map[Status]string{
	StatusUnknown: _StatusName[0:7],
	StatusGood:    _StatusName[7:11],
	StatusBad:     _StatusName[11:14],
}

// Name is the attribute of Status.
func (x Status) Name() string {
	if v, ok := _StatusMapName[x]; ok {
		return v
	}
	return fmt.Sprintf("Status(%d).Name", x)
}

// Val is the attribute of Status.
func (x Status) Val() int {
	return int(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Status) IsValid() bool {
	_, ok := _StatusMapName[x]
	return ok
}

// String implements the Stringer interface.
func (x Status) String() string {
	return x.Name()
}

var _StatusNameMap = map[string]Status{
	_StatusName[0:7]:                    StatusUnknown,
	strings.ToLower(_StatusName[0:7]):   StatusUnknown,
	_StatusName[7:11]:                   StatusGood,
	strings.ToLower(_StatusName[7:11]):  StatusGood,
	_StatusName[11:14]:                  StatusBad,
	strings.ToLower(_StatusName[11:14]): StatusBad,
}

// ParseStatus converts a string to a Status.
func ParseStatus(value string) (Status, error) {
	if x, ok := _StatusNameMap[value]; ok {
		return x, nil
	}
	if x, ok := _StatusNameMap[strings.ToLower(value)]; ok {
		return x, nil
	}
	return Status(0), fmt.Errorf("%s is %w", value, ErrInvalidStatus)
}
