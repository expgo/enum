package example

//@EnumConfig(forceupper)
//go:generate ag --dev-plugin=github.com/expgo/enum

// @EnumConfig(forcelower)
// @ENUM{
// DataSwap,
// BootNode,
// }
type ForceLowerType int
