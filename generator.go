package enum

import (
	"embed"
	"errors"
	"fmt"
	"github.com/expgo/ag/api"
	"github.com/expgo/structure"
	"go/ast"
	"reflect"
	"sort"
	"strings"
	"text/template"
)

//go:embed enum.tmpl
var enumTmpl embed.FS

type EnumGenerator struct {
	api.BaseGenerator[Enum]
}

func wrapType(inner string, sourceType string, targetType string) string {
	if sourceType == targetType {
		return inner
	}

	return fmt.Sprintf("%s(%s)", targetType, inner)
}

func defaultValueByKind(kind reflect.Kind) string {
	value, err := structure.ConvertToKind(0, kind)
	if err != nil {
		panic(errors.New(fmt.Sprintf("default value by kind err: %v", err)))
	}

	ret, err := structure.ConvertTo[string](value)
	if err != nil {
		panic(errors.New(fmt.Sprintf("default value by kind, convert to string err: %v", err)))
	}

	return ret
}

func NewEnumGenerator(allEnums []*Enum) *EnumGenerator {
	result := &EnumGenerator{}

	tmpl := template.New("enum")

	funcs := template.FuncMap{}
	funcs["IA"] = IndefiniteArticle
	funcs["WT"] = wrapType
	funcs["DV"] = defaultValueByKind
	tmpl.Funcs(funcs)

	result.Tmpl = template.Must(tmpl.ParseFS(enumTmpl, "*.tmpl"))

	result.DataList = append([]*Enum(nil), allEnums...)

	sort.Slice(result.DataList, func(i, j int) bool {
		x := result.DataList[i]
		y := result.DataList[j]
		return strings.Compare(x.Name, y.Name) < 0
	})

	return result
}

func (eg *EnumGenerator) GetImports() []string {
	return []string{"errors", "fmt"}
}

func annotationsToEnumConfig(annotations *api.Annotations, globalConfig *Config) (result *Config, err error) {
	enumConfAnnotation := annotations.FindAnnotationByName(AnnotationEnumConfig.Val())

	if enumConfAnnotation != nil {
		result = structure.Clone(globalConfig)
		err = enumConfAnnotation.To(result)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("pase annotation err: %v", err))
		}
		return result, nil
	}

	if globalConfig != nil {
		return structure.Clone(globalConfig), nil
	} else {
		return nil, nil
	}
}

func AnnotationsToEnum(annotations *api.Annotations, ts *ast.TypeSpec, globalConfig *Config) (*Enum, error) {
	enumAnnotation := annotations.FindAnnotationByName(AnnotationEnum.Val())
	if enumAnnotation == nil {
		return nil, nil
	}

	ec, err := annotationsToEnumConfig(annotations, globalConfig)
	if err != nil {
		return nil, err
	}

	enum := &Enum{
		Config: ec,
	}
	ec.enum = enum

	t, err := getEnumKindByName(fmt.Sprintf("%s", ts.Type))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("enum type err: %v", err))
	}

	enum.Name = ts.Name.Name
	enum.Type = t

	enum.Comment = strings.Join(enumAnnotation.Doc, "\n")
	if len(enum.Comment) == 0 {
		enum.Comment = enumAnnotation.Comment
	}

	err = enum.UpdateAttributes(enumAnnotation)
	if err != nil {
		return nil, err
	}

	err = enum.UpdateItems(enumAnnotation)
	if err != nil {
		return nil, err
	}

	err = enum.CheckValid()
	if err != nil {
		return nil, err
	}

	return enum, nil
}
