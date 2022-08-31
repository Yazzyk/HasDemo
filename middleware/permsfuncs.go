package middleware

import (
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/core"
)

type PermFuncWrapper struct {
	server    core.IServer
	functions htypes.Map
}

func NewPermFuncWrapper() *PermFuncWrapper {

	return new(PermFuncWrapper)
}

func (this *PermFuncWrapper) SetServer(s core.IServer) {
	this.server = s
	this.functions = make(htypes.Map)
	this.functions["CheckTest"] = this.CheckTest
}

func (this *PermFuncWrapper) Functions() htypes.Map {
	return this.functions
}

func (this *PermFuncWrapper) CheckTest(name string, change bool) bool {
	// 根据需求，可以对传参进行逻辑处理，返回false则无法通过,true则通过
	return false
}
