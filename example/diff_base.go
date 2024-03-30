package example

//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

/*
@EnumConfig(forcelower)
@ENUM{

	B3  = 03
	B4  = 04
	B5  = 5
	B6  = 0b110
	B7  = 0b111
	B8  = 0x08
	B9  = 0x09
	B10 = 0x0B
	B11 = 0x2B

}
*/
type DiffBase int
