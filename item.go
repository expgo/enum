package enum

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
)

const (
	ItemName  = "Name"
	ItemValue = "Val"
)

type Item struct {
	enum              *Enum
	idx               int
	Name              string
	Value             any
	DocComment        string
	LineComment       string
	AttributeData     []any
	IsBlankIdentifier bool
}

// GetCodeName return the item name used in code
func (ei *Item) GetCodeName() string {
	if ei.IsBlankIdentifier {
		return BlankIdentifier
	}

	casedName := ei.Name
	if ei.enum.Config.NoCamel {
		casedName = capitalize(ei.Name)
	} else {
		casedName = strcase.ToCamel(ei.Name)
	}

	if ei.enum.Config.NoPrefix {
		return ei.enum.Config.Prefix + casedName
	} else {
		return ei.enum.Config.Prefix + ei.enum.Name + casedName
	}
}

// GetName return the item real name, default equals with the code name, or an attribute named `Name`
func (ei *Item) GetName() string {
	nameAttr := ei.enum.FindAttributeByName(ItemName)
	if nameAttr == nil {
		panic("Name attribute is not set")
	}

	name := ei.AttributeData[nameAttr.idx].(string)

	// fix number _9600, see example/number.go
	if len(name) > 1 && strings.HasPrefix(name, BlankIdentifier) {
		name = name[1:]
	}

	if ei.enum.Config.ForceUpper {
		return strings.ToUpper(name)
	}

	if ei.enum.Config.ForceLower {
		return strings.ToLower(name)
	}

	return name
}

func (ei *Item) GetConstLine() string {
	if ei.Value == nil {
		if ei.idx == 0 {
			return fmt.Sprintf("%s %s = iota", ei.GetCodeName(), ei.enum.Name)
		} else {
			return ei.GetCodeName()
		}
	} else {
		if ei.IsBlankIdentifier {
			return BlankIdentifier
		} else {
			if ei.enum.Type == reflect.String {
				return fmt.Sprintf("%s %s = \"%s\"", ei.GetCodeName(), ei.enum.Name, ei.Value)
			} else {
				return fmt.Sprintf("%s %s = %v", ei.GetCodeName(), ei.enum.Name, ei.Value)
			}
		}
	}
}
