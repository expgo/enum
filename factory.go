package main

import (
	"github.com/expgo/ag/api"
	"github.com/expgo/factory"
	"go/ast"
)

type Factory struct {
}

func init() {
	factory.Singleton[Factory]()
}

func (f *Factory) Annotations() map[string][]api.AnnotationType {
	return map[string][]api.AnnotationType{
		AnnotationEnum.Name():       {api.AnnotationTypeGlobal, api.AnnotationTypeType},
		AnnotationEnumConfig.Name(): {api.AnnotationTypeType},
	}
}

func (f *Factory) New(typedAnnotations []*api.TypedAnnotation) (api.Generator, error) {
	var ec = factory.New[Config]()

	for _, typedAnnotation := range typedAnnotations {
		if typedAnnotation.Type == api.AnnotationTypeGlobal && typedAnnotation.Annotations != nil {
			if an := typedAnnotation.Annotations.FindAnnotationByName(AnnotationEnumConfig.Name()); an != nil {
				if err := an.To(ec); err != nil {
					return nil, err
				}
			}
		}
	}

	// get all enums
	var allEnums []*Enum
	for _, typedAnnotation := range typedAnnotations {
		if typedAnnotation.Type == api.AnnotationTypeType && typedAnnotation.Annotations != nil {
			if en, err := AnnotationsToEnum(typedAnnotation.Annotations, typedAnnotation.Node.(*ast.TypeSpec), ec); err != nil {
				return nil, err
			} else {
				allEnums = append(allEnums, en)
			}
		}
	}

	if len(allEnums) > 0 {
		// create enums
		return NewEnumGenerator(allEnums), nil
	}

	return nil, nil
}
