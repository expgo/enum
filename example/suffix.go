package example

//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../ --file-suffix .enum.gen

// Suffix @ENUM{gen}
type Suffix string
