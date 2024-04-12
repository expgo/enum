package example

//go:generate ag --dev-plugin=github.com/expgo/enum

// @EnumConfig(marshal, prefix="AcmeInc_", noprefix, nocamel, names)
// Shops @ENUM{
// SOME_PLACE_AWESOME,
// SomewhereElse,
// LocationUnknown
// }
type Shop string
