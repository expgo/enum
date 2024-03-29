package main

import (
	"fmt"
	"github.com/expgo/generic/stream"
	"reflect"
)

type Config struct {
	enum            *Enum
	Prefix          string
	NoPrefix        bool   `value:"false"` // 所有生成的枚举不携带类型名称前缀
	StringParse     bool   `value:"true"`
	StringParseName string `value:"Name"`
	Flag            bool   `value:"false"`
	MustParse       bool   `value:"false"`
	Marshal         bool   `value:"false"`
	MarshalName     string `value:"Name"`
	Sql             bool   `value:"false"`
	SqlName         string `value:"Val"`
	Names           bool   `value:"false"` // enum name list
	Values          bool   `value:"false"` // enum item list
	NoCase          bool   `value:"false"` // case insensitivity
	NoCamel         bool   `value:"false"`
	NoComments      bool   `value:"false"`
	Ptr             bool   `value:"false"`
	ForceUpper      bool   `value:"false"`
	ForceLower      bool   `value:"false"`
	PanicIfInvalid  bool   `value:"false"`
}

func (ec *Config) SetStringParse(stringParse bool) {
	// if stringParse set to false, flag must be set to false
	if !stringParse {
		ec.Flag = false
	}
	ec.StringParse = stringParse
}

func (ec *Config) SetFlag(flag bool) {
	// if set flag true, the stringParse must be set to true
	if flag {
		ec.StringParse = true
	}
	ec.Flag = flag
}

func (ec *Config) SetForceLower(lower bool) {
	if lower {
		if ec.ForceUpper {
			ec.ForceUpper = false
		}
	}
	ec.ForceLower = lower
}

func (ec *Config) SetForceUpper(upper bool) {
	if upper {
		if ec.ForceLower {
			ec.ForceLower = false
		}
	}
	ec.ForceUpper = upper
}

func (ec *Config) checkConfigAttributeName(paramName, errName string) error {
	if attr := ec.enum.FindAttributeByName(paramName); attr == nil {
		return fmt.Errorf("enum config %s must exist in enum attributes", errName)
	} else {
		// all enumTypes is number or string, bool and float not suitable as a map key
		if !stream.Must(stream.Of(enumTypes).Contains(attr.Type, func(x, y reflect.Kind) (bool, error) { return x == y, nil })) {
			return fmt.Errorf("%s 's type muse be number or string", errName)
		}
	}

	return nil
}

func (ec *Config) CheckValid() (err error) {
	if err = ec.checkConfigAttributeName(ec.StringParseName, "StringParseName"); err != nil {
		return err
	}

	if err = ec.checkConfigAttributeName(ec.MarshalName, "MarshalName"); err != nil {
		return err
	}

	if err = ec.checkConfigAttributeName(ec.SqlName, "SqlName"); err != nil {
		return err
	}

	return nil
}
