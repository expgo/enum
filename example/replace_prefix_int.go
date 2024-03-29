package example

//go:generate go run ../../main.go

// @EnumConfig(marshal, prefix="AcmeInt_", noprefix, nocamel, names)
// Shops @ENUM{
// SOME_PLACE_AWESOME,
// SomewhereElse,
// LocationUnknown
// }
type IntShop int
