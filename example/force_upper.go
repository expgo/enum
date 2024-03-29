package example

//@EnumConfig(forcelower)
//go:generate go run ../../main.go

// @EnumConfig(forceupper)
// @ENUM{
// DataSwap,
// BootNode,
// }
type ForceUpperType int
