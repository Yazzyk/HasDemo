package objs

import (
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/services/hdatasvs"
)

type ViewTestInfo struct {
	hdatasvs.DataView `data:"key:test_info;from:test;join:info@test.id=info.test_id"`

	ID    uint64 `json:"id" data:"key:id;field:test.id"`
	User  string `json:"user" data:"key:user;field:info.user"`
	Name  string `json:"name" data:"key:name;field:test.name"`
	Phone string `json:"phone" data:"key:phone;field:test.phone"`
}

type ViewAll struct {
	hdatasvs.DataView `data:"key:all;from:info;join:test@test.id=info.test_id,test2@test2.id=info.test2_id"`

	InfoID uint64 `json:"info_id" data:"key:info_id;field:info.id"`
	User   string `json:"user" data:"key:user;field:info.user"`
	Name   string `json:"name" data:"key:name;field:test2.name"`
	Addr   string `json:"addr" data:"key:addr;field:test.addr"`
}

func Views() []htypes.Any {
	return []htypes.Any{
		ViewTestInfo{}, ViewAll{},
	}
}
