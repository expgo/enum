// Code generated by https://github.com/expgo/ag DO NOT EDIT.
// Plugins:
//   - github.com/expgo/enum

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

const (
	ProjectStatusPending ProjectStatus = iota
	ProjectStatusInWork
	ProjectStatusCompleted
	ProjectStatusRejected
)

const (
	ProjectStrStatusPending   ProjectStrStatus = "pending"
	ProjectStrStatusInWork    ProjectStrStatus = "inWork"
	ProjectStrStatusCompleted ProjectStrStatus = "completed"
	ProjectStrStatusRejected  ProjectStrStatus = "rejected"
)

const (
	ProjectStrStatusIntCodePending   ProjectStrStatusIntCode = "pending"
	ProjectStrStatusIntCodeInWork    ProjectStrStatusIntCode = "inWork"
	ProjectStrStatusIntCodeCompleted ProjectStrStatusIntCode = "completed"
	ProjectStrStatusIntCodeRejected  ProjectStrStatusIntCode = "rejected"
)

var ErrInvalidProjectStatus = errors.New("not a valid ProjectStatus")

var _ProjectStatusName = "pendinginWorkcompletedrejected"

var _ProjectStatusMapName = map[ProjectStatus]string{
	ProjectStatusPending:   _ProjectStatusName[0:7],
	ProjectStatusInWork:    _ProjectStatusName[7:13],
	ProjectStatusCompleted: _ProjectStatusName[13:22],
	ProjectStatusRejected:  _ProjectStatusName[22:30],
}

// Name is the attribute of ProjectStatus.
func (x ProjectStatus) Name() string {
	if v, ok := _ProjectStatusMapName[x]; ok {
		return v
	}
	return fmt.Sprintf("ProjectStatus(%d).Name", x)
}

// Val is the attribute of ProjectStatus.
func (x ProjectStatus) Val() int {
	return int(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ProjectStatus) IsValid() bool {
	_, ok := _ProjectStatusMapName[x]
	return ok
}

// String implements the Stringer interface.
func (x ProjectStatus) String() string {
	return x.Name()
}

var _ProjectStatusNameMap = map[string]ProjectStatus{
	_ProjectStatusName[0:7]:   ProjectStatusPending,
	_ProjectStatusName[7:13]:  ProjectStatusInWork,
	_ProjectStatusName[13:22]: ProjectStatusCompleted,
	_ProjectStatusName[22:30]: ProjectStatusRejected,
}

// ParseProjectStatus converts a string to a ProjectStatus.
func ParseProjectStatus(value string) (ProjectStatus, error) {
	if x, ok := _ProjectStatusNameMap[value]; ok {
		return x, nil
	}
	return ProjectStatus(0), fmt.Errorf("%s is %w", value, ErrInvalidProjectStatus)
}

func (x ProjectStatus) Ptr() *ProjectStatus {
	return &x
}

// MarshalText implements the text marshaller method.
func (x ProjectStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ProjectStatus) UnmarshalText(text []byte) error {
	val, err := ParseProjectStatus(string(text))
	if err != nil {
		return err
	}
	*x = val
	return nil
}

var ErrProjectStatusNilPtr = errors.New("value pointer is nil")

// Scan implements the Scanner interface.
func (x *ProjectStatus) Scan(value any) (err error) {
	if value == nil {
		*x = ProjectStatus(0)
		return
	}

	switch v := value.(type) {
	case int:
		*x = ProjectStatus(v)
	case int64:
		*x = ProjectStatus(v)
	case uint:
		*x = ProjectStatus(v)
	case uint64:
		*x = ProjectStatus(v)
	case float64:
		*x = ProjectStatus(v)
	case *int:
		if v == nil {
			return ErrProjectStatusNilPtr
		}
		*x = ProjectStatus(*v)
	case *int64:
		if v == nil {
			return ErrProjectStatusNilPtr
		}
		*x = ProjectStatus(*v)
	case *uint:
		if v == nil {
			return ErrProjectStatusNilPtr
		}
		*x = ProjectStatus(*v)
	case *uint64:
		if v == nil {
			return ErrProjectStatusNilPtr
		}
		*x = ProjectStatus(*v)
	case *float64:
		if v == nil {
			return ErrProjectStatusNilPtr
		}
		*x = ProjectStatus(*v)
	case ProjectStatus:
		*x = v
	case *ProjectStatus:
		if v == nil {
			return ErrProjectStatusNilPtr
		}
		*x = *v
	}

	if !x.IsValid() {
		return ErrInvalidProjectStatus
	}
	return
}

// Value implements the driver Valuer interface.
func (x ProjectStatus) Value() (driver.Value, error) {
	return x.Val(), nil
}

var ErrInvalidProjectStrStatus = errors.New("not a valid ProjectStrStatus")

var _ProjectStrStatusNameMap = map[string]ProjectStrStatus{
	"pending":   ProjectStrStatusPending,
	"inWork":    ProjectStrStatusInWork,
	"completed": ProjectStrStatusCompleted,
	"rejected":  ProjectStrStatusRejected,
}

// Name is the attribute of ProjectStrStatus.
func (x ProjectStrStatus) Name() string {
	if v, ok := _ProjectStrStatusNameMap[string(x)]; ok {
		return string(v)
	}
	return fmt.Sprintf("ProjectStrStatus(%s).Name", string(x))
}

// Val is the attribute of ProjectStrStatus.
func (x ProjectStrStatus) Val() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ProjectStrStatus) IsValid() bool {
	_, ok := _ProjectStrStatusNameMap[string(x)]
	return ok
}

// String implements the Stringer interface.
func (x ProjectStrStatus) String() string {
	return x.Name()
}

// ParseProjectStrStatus converts a string to a ProjectStrStatus.
func ParseProjectStrStatus(value string) (ProjectStrStatus, error) {
	if x, ok := _ProjectStrStatusNameMap[value]; ok {
		return x, nil
	}
	return "", fmt.Errorf("%s is %w", value, ErrInvalidProjectStrStatus)
}

func (x ProjectStrStatus) Ptr() *ProjectStrStatus {
	return &x
}

// MarshalText implements the text marshaller method.
func (x ProjectStrStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ProjectStrStatus) UnmarshalText(text []byte) error {
	val, err := ParseProjectStrStatus(string(text))
	if err != nil {
		return err
	}
	*x = val
	return nil
}

var ErrProjectStrStatusNilPtr = errors.New("value pointer is nil")

// Scan implements the Scanner interface.
func (x *ProjectStrStatus) Scan(value any) (err error) {
	if value == nil {
		*x = ""
		return
	}

	switch v := value.(type) {
	case string:
		*x = ProjectStrStatus(v)
	case []byte:
		*x = ProjectStrStatus(string(v))
	case *string:
		if v == nil {
			return ErrProjectStrStatusNilPtr
		}
		*x = ProjectStrStatus(*v)
	case ProjectStrStatus:
		*x = v
	case *ProjectStrStatus:
		if v == nil {
			return ErrProjectStrStatusNilPtr
		}
		*x = *v
	}

	if !x.IsValid() {
		return ErrInvalidProjectStrStatus
	}
	return
}

// Value implements the driver Valuer interface.
func (x ProjectStrStatus) Value() (driver.Value, error) {
	return x.Val(), nil
}

var ErrInvalidProjectStrStatusIntCode = errors.New("not a valid ProjectStrStatusIntCode")

var _ProjectStrStatusIntCodeNameMap = map[string]ProjectStrStatusIntCode{
	"pending":   ProjectStrStatusIntCodePending,
	"inWork":    ProjectStrStatusIntCodeInWork,
	"completed": ProjectStrStatusIntCodeCompleted,
	"rejected":  ProjectStrStatusIntCodeRejected,
}

// Name is the attribute of ProjectStrStatusIntCode.
func (x ProjectStrStatusIntCode) Name() string {
	if v, ok := _ProjectStrStatusIntCodeNameMap[string(x)]; ok {
		return string(v)
	}
	return fmt.Sprintf("ProjectStrStatusIntCode(%s).Name", string(x))
}

var _ProjectStrStatusIntCodeMapDbCode = map[ProjectStrStatusIntCode]int{
	ProjectStrStatusIntCodePending:   0,
	ProjectStrStatusIntCodeInWork:    10,
	ProjectStrStatusIntCodeCompleted: 20,
	ProjectStrStatusIntCodeRejected:  30,
}

// DbCode is the attribute of ProjectStrStatusIntCode.
func (x ProjectStrStatusIntCode) DbCode() int {
	if v, ok := _ProjectStrStatusIntCodeMapDbCode[x]; ok {
		return v
	}
	return 0
}

// Val is the attribute of ProjectStrStatusIntCode.
func (x ProjectStrStatusIntCode) Val() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ProjectStrStatusIntCode) IsValid() bool {
	_, ok := _ProjectStrStatusIntCodeNameMap[string(x)]
	return ok
}

// String implements the Stringer interface.
func (x ProjectStrStatusIntCode) String() string {
	return x.Name()
}

// ParseProjectStrStatusIntCode converts a string to a ProjectStrStatusIntCode.
func ParseProjectStrStatusIntCode(value string) (ProjectStrStatusIntCode, error) {
	if x, ok := _ProjectStrStatusIntCodeNameMap[value]; ok {
		return x, nil
	}
	return "", fmt.Errorf("%s is %w", value, ErrInvalidProjectStrStatusIntCode)
}

func (x ProjectStrStatusIntCode) Ptr() *ProjectStrStatusIntCode {
	return &x
}

// MarshalText implements the text marshaller method.
func (x ProjectStrStatusIntCode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ProjectStrStatusIntCode) UnmarshalText(text []byte) error {
	val, err := ParseProjectStrStatusIntCode(string(text))
	if err != nil {
		return err
	}
	*x = val
	return nil
}

var ErrProjectStrStatusIntCodeNilPtr = errors.New("value pointer is nil")

var _ProjectStrStatusIntCodeDbCodeMap = map[int]ProjectStrStatusIntCode{
	0:  ProjectStrStatusIntCodePending,
	10: ProjectStrStatusIntCodeInWork,
	20: ProjectStrStatusIntCodeCompleted,
	30: ProjectStrStatusIntCodeRejected,
}

// Scan implements the Scanner interface.
func (x *ProjectStrStatusIntCode) Scan(value any) (err error) {
	if value == nil {
		*x = ""
		return
	}

	var ok bool
	switch v := value.(type) {
	case int:
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[v]
	case int64:
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(v)]
	case uint:
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(v)]
	case uint64:
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(v)]
	case float64:
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(v)]
	case *int:
		if v == nil {
			return ErrProjectStrStatusIntCodeNilPtr
		}
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[*v]
	case *int64:
		if v == nil {
			return ErrProjectStrStatusIntCodeNilPtr
		}
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(*v)]
	case *uint:
		if v == nil {
			return ErrProjectStrStatusIntCodeNilPtr
		}
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(*v)]
	case *uint64:
		if v == nil {
			return ErrProjectStrStatusIntCodeNilPtr
		}
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(*v)]
	case *float64:
		if v == nil {
			return ErrProjectStrStatusIntCodeNilPtr
		}
		*x, ok = _ProjectStrStatusIntCodeDbCodeMap[int(*v)]
	case ProjectStrStatusIntCode:
		*x = v
		ok = x.IsValid()
	case *ProjectStrStatusIntCode:
		if v == nil {
			return ErrProjectStrStatusIntCodeNilPtr
		}
		*x = *v
		ok = x.IsValid()
	}

	if !ok {
		return ErrInvalidProjectStrStatusIntCode
	}
	return
}

// Value implements the driver Valuer interface.
func (x ProjectStrStatusIntCode) Value() (driver.Value, error) {
	return x.DbCode(), nil
}
