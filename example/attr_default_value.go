package example

//go:generate ag --dev-plugin=github.com/expgo/enum

/*
Device : protocol deveice

	@Enum(code int, isBit bool, str string, a int16) {
		SM(0x91, true, "123", 1)		// 特殊继电器, bit, 10
		SD(0xA9, false, "456", 2)		// 特殊寄存器, word, 10
	}
*/
type Device int
