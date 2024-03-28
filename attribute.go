package enum

import (
	"bytes"
	"fmt"
	"github.com/expgo/structure"
	"reflect"
	"strings"
)

type Attribute struct {
	enum                   *Enum
	idx                    int
	isValue                bool
	enum2AttributeRendered bool
	attribute2EnumRendered bool
	Name                   string
	Type                   reflect.Kind
	Comment                string
}

func (ea *Attribute) Enum() *Enum {
	return ea.enum
}

func (ea *Attribute) IsValue() bool {
	return ea.isValue
}

func (ea *Attribute) Enum2AttributeVarName() string {
	if ea.isValue {
		return ""
	}
	return fmt.Sprintf("_%sMap%s", ea.enum.Name, ea.Name)
}

func (ea *Attribute) Enum2AttributeMap() string {
	if ea.enum2AttributeRendered == true {
		return ""
	}

	if ea.isValue {
		return ""
	}

	if ea.Name == ItemName && ea.enum.Type == reflect.String {
		return ""
	}

	buf := bytes.NewBuffer([]byte{})

	buf.WriteString(fmt.Sprintf("var %s = map[%s]%s{\n", ea.Enum2AttributeVarName(), ea.enum.Name, ea.Type.String()))
	if ea.idx == 0 {
		index := 0
		for _, item := range ea.enum.GetItems() {
			nextIndex := index + len(item.GetName())
			buf.WriteString(fmt.Sprintf("	%s: _%sName[%d:%d],\n", item.GetCodeName(), ea.enum.Name, index, nextIndex))
			index = nextIndex
		}
	} else {
		for _, item := range ea.enum.GetItems() {
			switch ea.Type {
			case reflect.String:
				buf.WriteString(fmt.Sprintf("	%s: \"%s\",\n", item.GetCodeName(), structure.MustConvertTo[string](item.AttributeData[ea.idx])))
			default:
				buf.WriteString(fmt.Sprintf("	%s: %v,\n", item.GetCodeName(), structure.MustConvertToKind(item.AttributeData[ea.idx], ea.Type)))
			}
		}
	}
	buf.WriteString("}\n")

	ea.enum2AttributeRendered = true

	return buf.String()
}

func (ea *Attribute) Attribute2EnumVarName() string {
	if ea.isValue {
		return ""
	}
	return fmt.Sprintf("_%s%sMap", ea.enum.Name, ea.Name)
}

func (ea *Attribute) Attribute2EnumMap() string {
	if ea.attribute2EnumRendered == true {
		return ""
	}

	if ea.isValue {
		return ""
	}

	buf := bytes.NewBuffer([]byte{})

	buf.WriteString(fmt.Sprintf("var %s = map[%s]%s{\n", ea.Attribute2EnumVarName(), ea.Type.String(), ea.enum.Name))

	index := 0
	for _, item := range ea.enum.GetItems() {
		var itemValue = ""
		nextIndex := index + len(item.GetName())

		if ea.Type == reflect.String {
			if ea.Name == ItemName && ea.enum.Type != reflect.String {
				itemValue = item.GetName()
				buf.WriteString(fmt.Sprintf("	_%sName[%d:%d]: %s,\n", ea.enum.Name, index, nextIndex, item.GetCodeName()))
			} else {
				itemValue = structure.MustConvertTo[string](item.AttributeData[ea.idx])
				buf.WriteString(fmt.Sprintf("	\"%s\": %s,\n", itemValue, item.GetCodeName()))
			}
		} else {
			buf.WriteString(fmt.Sprintf("	%v: %s,\n", structure.MustConvertToKind(item.AttributeData[ea.idx], ea.Type), item.GetCodeName()))
		}
		if ea.enum.Config.NoCase && ea.Type == reflect.String && (itemValue != strings.ToLower(itemValue)) {
			if ea.Name == ItemName && ea.enum.Type != reflect.String {
				buf.WriteString(fmt.Sprintf("	strings.ToLower(_%sName[%d:%d]): %s,\n", ea.enum.Name, index, nextIndex, item.GetCodeName()))
			} else {
				buf.WriteString(fmt.Sprintf("	\"%s\": %s,\n", strings.ToLower(itemValue), item.GetCodeName()))
			}
		}

		index = nextIndex
	}
	buf.WriteString("}\n")

	ea.attribute2EnumRendered = true

	return buf.String()
}

func (ea *Attribute) FirstValueBits() int {
	if ea.Type == reflect.String {
		return 0
	}
	return reflect.TypeOf(ea.enum.Items[0].AttributeData[ea.idx]).Bits()
}

func (ea *Attribute) ParseNumberFuncString() string {
	if ea.Type.String()[0] == 'u' {
		return fmt.Sprintf("strconv.ParseUint(value, 0, %d)", ea.FirstValueBits())
	} else {
		return fmt.Sprintf("strconv.ParseInt(value, 0, %d)", ea.FirstValueBits())
	}
}
