package example

//go:generate go run ../../main.go

// @EnumConfig(sql, ptr, marshal, nocomments)
// @ENUM{pending, inWork, completed, rejected}
type ProjectStatus int

// @EnumConfig(sql, ptr, marshal, nocomments)
// @ENUM{pending, inWork, completed, rejected}
type ProjectStrStatus string

// @EnumConfig(sql, ptr, marshal, nocomments, sqlName=dbCode)
//
//	@ENUM(dbCode int) {
//		pending(0)
//		inWork(10)
//		completed(20)
//		rejected(30)
//	}
type ProjectStrStatusIntCode string
