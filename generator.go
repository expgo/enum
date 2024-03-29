package enum

import (
	"embed"
	"errors"
	"fmt"
	"github.com/expgo/ag/api"
	"github.com/expgo/generic/stream"
	"github.com/expgo/structure"
	"go/ast"
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

func NewEnumGenerator(allEnums []*Enum) *EnumGenerator {
	result := &EnumGenerator{}

	tmpl := template.New("enum")

	funcs := template.FuncMap{}
	funcs["IA"] = IndefiniteArticle
	funcs["WT"] = wrapType
	tmpl.Funcs(funcs)

	result.Tmpl = template.Must(tmpl.ParseFS(enumTmpl, "*.tmpl"))

	result.DataList = stream.Must(stream.Of(allEnums).Sort(func(x, y *Enum) int { return strings.Compare(x.Name, y.Name) }).ToSlice())

	return result
}

func (eg *EnumGenerator) GetImports() []string {
	return []string{"errors", "fmt"}
}

func annotationsToEnumConfig(annotations *api.Annotations, globalConfig *Config) (result *Config, err error) {
	enumConfAnnotation := annotations.FindAnnotationByName(AnnotationEnumConfig.Name())

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
	enumAnnotation := annotations.FindAnnotationByName(AnnotationEnum.Name())
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
