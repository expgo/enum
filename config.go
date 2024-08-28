package enum

import (
	"fmt"
)

type Config struct {
	enum            *Enum
	Prefix          string // default prefix will be the enum type name
	NoPrefix        bool   `value:"false"` // if true, the enum will use the enum item name direct, else will create enum item with enum type prefix
	StringParse     bool   `value:"true"`  // create
	StringParseName string `value:"Name"`  // parse method use which enum parameter
	Flag            bool   `value:"false"` // if true, create code used with flag
	MustParse       bool   `value:"false"` // if true, create muse parse method
	Marshal         bool   `value:"false"` // if true, create Marshal interface method
	MarshalName     string `value:"Name"`  // Marshal interface use which enum parameter
	Sql             bool   `value:"false"` // if true, create code used with sql
	SqlName         string `value:"Val"`   // sql method use which enum parameter
	Names           bool   `value:"false"` // if true, create enum name list
	Values          bool   `value:"false"` // if true, create enum item list
	NoCase          bool   `value:"false"` // case insensitivity with parse method
	NoCamel         bool   `value:"false"` // if true, do nothing with enum name
	NoComments      bool   `value:"false"` // if true, will not create comments
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
		contains := false
		for _, enumType := range enumTypes {
			if enumType == attr.Type {
				contains = true
				break
			}
		}
		if !contains {
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
