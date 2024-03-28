package enum

import (
	"github.com/expgo/ag"
	"github.com/expgo/factory"
	"go/ast"
)

type Factory struct {
}

func init() {
	factory.Singleton[Factory]()
}
func (f *Factory) Annotations() map[string][]ag.AnnotationType {
	return map[string][]ag.AnnotationType{
		AnnotationEnum.Name():       {ag.AnnotationTypeGlobal, ag.AnnotationTypeType},
		AnnotationEnumConfig.Name(): {ag.AnnotationTypeType},
	}
}

func (f *Factory) New(typedAnnotations []*ag.TypedAnnotation) (ag.Generator, error) {
	var ec = factory.New[Config]()

	for _, typedAnnotation := range typedAnnotations {
		if typedAnnotation.Type == ag.AnnotationTypeGlobal && typedAnnotation.Annotations != nil {
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
		if typedAnnotation.Type == ag.AnnotationTypeType && typedAnnotation.Annotations != nil {
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
