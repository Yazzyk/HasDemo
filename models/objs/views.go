package objs

import (
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/services/hdatasvs"
)

type ViewTestInfo struct {
	hdatasvs.DataView `data:"key:test_info;from:test;join:info@test.id=info.test_id"`

	ID uint64 `json:"id" data:"key:id;field:test.id"`
	User string `json:"user" data:"key:user;field:info.user"`
	Name string `json:"name" data:"key:name;field:test.name"`
	Phone string `json:"phone" data:"key:phone;field:test.phone"`
}

func Views() []htypes.Any  {
	return []htypes.Any{
		ViewTestInfo{},
	}
}