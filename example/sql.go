package example

//go:generate ag --dev-plugin=github.com/expgo/enum

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
