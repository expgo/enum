package example

//go:generate ag --dev-plugin=github.com/expgo/enum

// @EnumConfig(prefix="AcmeInc")
// Products of AcmeInc @ENUM{
// Anvil,
// Dynamite,
// Glue
// }
type Product int32
