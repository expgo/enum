package example

//@EnumConfig(nocase)
//go:generate go run ../../main.go

/*
@ENUM{
Unknown = -1,
Good,
Bad
}.
*/
type Status int

/*
@ENUM{
Unknown = -5,
Good,
Bad,
Ugly
}.
*/
type AllNegative int
