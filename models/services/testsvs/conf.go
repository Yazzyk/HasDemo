package testsvs

import "github.com/drharryhe/has/core"

type TestService struct {
	core.ServiceConf

	TestBool bool // 从配置文件定义一个bool值
	TestName string // 从配置文件定义一个字符串
}
// 我们建议在项目开发中，开发者为此处的每个字段都写上注释，便于后期维护