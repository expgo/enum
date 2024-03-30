package example

//@EnumConfig(forcelower)
//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

// @EnumConfig(forceupper)
// @ENUM{
// DataSwap,
// BootNode,
// }
type ForceUpperType int
