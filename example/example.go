package example

// @EnumConfig(marshal, nocase, flag, values, names)
//go:generate ag --dev-plugin=github.com/expgo/enum

// X is doc'ed
type X struct{}

// Make x @ENUM{Toyota,_,Chevy,_,Ford,_,Tesla,_,Hyundai,_,Nissan,_,Jaguar,_,Audi,_,BMW,_,Mercedes_Benz,_,Volkswagon}
type Make int32

// Make x @ENUM{start=20,middle,end,ps,pps,ppps}
type NoZeros int32
