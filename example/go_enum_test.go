package example

//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

/*
@ENUM(Name string){

	ABCDX("ABCD (x)"),
	EFGHY("EFGH (y)"),

}
*/
type TestOnlyEnum string
