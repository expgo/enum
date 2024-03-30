package example

//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

// Animal x @ENUM(Name string){
// Cat(_),
// Dog(Dog),
// Fish("Fish")
// FishPlusPlus("Fish++")
// FishSharp("Fish#")
// }.
type Animal int32
