package example

//go:generate ag --dev-plugin=github.com/expgo/enum

/*
DataType 数据类型

	@EnumConfig(NoCamel)
	@Enum {
		AI = 0x0A   // 表示模拟量
		DI = 0x0D   // 表示数字量
	}
*/
type DataType byte
