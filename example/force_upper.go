package example

//@EnumConfig(forcelower)
//go:generate ag --dev-plugin=github.com/expgo/enum

// @EnumConfig(forceupper)
// @ENUM{
// DataSwap,
// BootNode,
// }
type ForceUpperType int
