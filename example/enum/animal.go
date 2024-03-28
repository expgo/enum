package enum

//go:generate go run ../../main.go

// Animal x @ENUM(Name string){
// Cat(_),
// Dog(Dog),
// Fish("Fish")
// FishPlusPlus("Fish++")
// FishSharp("Fish#")
// }.
type Animal int32
