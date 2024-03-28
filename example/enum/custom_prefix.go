package enum

//go:generate go run ../../main.go

// @EnumConfig(prefix="AcmeInc")
// Products of AcmeInc @ENUM{
// Anvil,
// Dynamite,
// Glue
// }
type Product int32
