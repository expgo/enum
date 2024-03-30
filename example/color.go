package example

//go:generate ag --dev-plugin=github.com/expgo/enum --dev-plugin-dir=../

// Color is an enumeration of colors that are allowed.
// @EnumConfig(marshal, noCase, Mustparse, ptr)
/* @ENUM (Name string){
Black(_), White(_), Red(_)
Green(_) = 33 // Green starts with 33
*/
// Blue(_)
// grey(_)=45
// _
// _
// yellow(_)
// blue_green("blue-green")
// red_orange("red-orange")
// yellow_green(_)
// red_orange_blue("red-orange-blue")
// }
type Color int
