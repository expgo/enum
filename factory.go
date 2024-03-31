package enum

import (
	"github.com/expgo/ag/api"
	"github.com/expgo/factory"
	"go/ast"
)

// @Singleton
type Factory struct {
}

func (f *Factory) Annotations() map[string][]api.AnnotationType {
	return map[string][]api.AnnotationType{
		AnnotationEnum.Val():       {api.AnnotationTypeType},
		AnnotationEnumConfig.Val(): {api.AnnotationTypeGlobal, api.AnnotationTypeType},
	}
}

func (f *Factory) New(typedAnnotations []*api.TypedAnnotation) (api.Generator, error) {
	var ec = factory.New[Config]()

	for _, typedAnnotation := range typedAnnotations {
		if typedAnnotation.Type == api.AnnotationTypeGlobal && typedAnnotation.Annotations != nil {
			if an := typedAnnotation.Annotations.FindAnnotationByName(AnnotationEnumConfig.Val()); an != nil {
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
