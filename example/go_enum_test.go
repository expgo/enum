package example

//go:generate ag --dev-plugin=github.com/expgo/enum

/*
@ENUM(Name string){

	ABCDX("ABCD (x)"),
	EFGHY("EFGH (y)"),

}
*/
type TestOnlyEnum string
