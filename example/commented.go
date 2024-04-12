package example

//@EnumConfig(marshal, nocase)
//go:generate ag --dev-plugin=github.com/expgo/enum

// Commented is an enumeration of commented values
/*
@ENUM{
value1 // Commented value 1
value2
value3 // Commented value 3
}
*/
type Commented int

// ComplexCommented has some extra complicated parsing rules.
/*
@ENUM{
	_, // Placeholder with a ','  in it. (for harder testing)
value1 // Commented value 1
value2,
_
_
value3 // Commented value 3
}
*/
type ComplexCommented int
