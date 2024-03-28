package enum

//go:generate go run ../../main.go

/*
@ENUM(Name string){

	ABCDX("ABCD (x)"),
	EFGHY("EFGH (y)"),

}
*/
type TestOnlyEnum string
