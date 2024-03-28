package enum

//go:generate go run ../../main.go

// @EnumConfig(marshal, prefix="AcmeInc_", noprefix, nocamel, names)
// Shops @ENUM{
// SOME_PLACE_AWESOME,
// SomewhereElse,
// LocationUnknown
// }
type Shop string
