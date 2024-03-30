package example

//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

// @EnumConfig(marshal, prefix="AcmeInt_", noprefix, nocamel, names)
// Shops @ENUM{
// SOME_PLACE_AWESOME,
// SomewhereElse,
// LocationUnknown
// }
type IntShop int
