package hooks

import (
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/core"
	"github.com/drharryhe/has/services/hapauthsvs"
)

type ApHooks struct {
}

func (this *ApHooks) LoginHook(service core.IService, req *hapauthsvs.LoginRequest, res *core.CallerResponse){

	res.Data = htypes.Map{
		"hello": "HAS",
	}
}