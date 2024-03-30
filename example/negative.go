package example

//@EnumConfig(nocase)
//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

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
