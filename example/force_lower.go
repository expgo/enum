package example

//@EnumConfig(forceupper)
//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

// @EnumConfig(forcelower)
// @ENUM{
// DataSwap,
// BootNode,
// }
type ForceLowerType int
