package testmiddleware

import (
	"github.com/drharryhe/has/common/herrors"
	"github.com/drharryhe/has/common/hlogger"
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/core"
)

func New() core.IAPIMiddleware {
	return new(Middleware)
}

type Middleware struct {
	core.InOutMiddleware
	conf      TestMiddleware
}

func (this *Middleware) Open(gw core.IAPIGateway, ins core.IAPIMiddleware) *herrors.Error {
	_ = this.BaseMiddleware.Open(gw, ins)
	return nil
}

func (this *Middleware) HandleIn(seq uint64, version string, api string, data htypes.Map) (bool, *herrors.Error) {
	hlogger.Info("in: ", seq, version, api, data)
	return false, nil
}

func (this *Middleware) HandleOut(seq uint64, version string, api string, result htypes.Any, e *herrors.Error) (stop bool, err *herrors.Error) {
	hlogger.Info("out: ", seq, version, api, result, e)
	return false, nil
}

func (this *Middleware) Config() core.IEntityConf {
	return &this.conf
}

func (this *Middleware) EntityStub() *core.EntityStub {
	return core.NewEntityStub(
		&core.EntityStubOptions{
			Owner: this,
		})
}