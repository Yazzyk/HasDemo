package objs

import (
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/services/hdatasvs"
)

type MyTest struct {
	hdatasvs.DataObject `data:"db:mysql;key:test"`

	ID    uint64 `json:"id" gorm:"primaryKey" data:"primary;deny:create"`
	Name  string `json:"name" gorm:"size:50" data:"require:create"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Addr  string `json:"addr"`
}

type MyInfo struct {
	hdatasvs.DataObject `data:"db:mysql;key:info"`

	ID     uint64 `json:"id" gorm:"primaryKey" data:"primary;deny:create"`
	TestID uint64 `json:"test_id"`
	User   string `json:"user" gorm:"size:20" data:"require:create"`
	Pwd    string `json:"pwd" gorm:"size:50" data:"require:create"`
}

func Objects() []htypes.Any {
	return []htypes.Any{
		MyTest{}, MyInfo{},
	}
}
