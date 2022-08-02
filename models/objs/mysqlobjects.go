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

func Objects() []htypes.Any {
	return []htypes.Any{
		MyTest{},
	}
}
