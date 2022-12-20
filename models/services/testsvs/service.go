package testsvs

import (
	"HASDemo/plugin/hasdemoplugin"
	"github.com/drharryhe/has/common/herrors"
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/core"
	"github.com/drharryhe/has/plugins/hdatabaseplugin"
	"gorm.io/gorm"
)

type Service struct {
	core.Service
	conf     TestService
	TestName string   // testName，你可以自定义这些字段，用于存储必要信息
	db       *gorm.DB // mysql 数据库
	gateway  *core.APIGateWayImplement
}

// 我们建议在项目开发中，开发者为此处自定义的字段都写上注释，便于后期维护

func (this *Service) Open(s core.IServer, instance core.IService, options htypes.Any) *herrors.Error {
	if err := this.Service.Open(s, instance, options); err != nil {
		return err
	}

	this.TestName = "World"
	// 使用配置文件中的值
	if this.conf.TestBool {
		this.TestName = this.conf.TestName
	}
	// 从插件使用mysql数据库
	this.db = this.UsePlugin("DatabasePlugin").(*hdatabaseplugin.Plugin).Capability().(map[string]*gorm.DB)["mysql"]
	this.UsePlugin("HASDemoPlugin").(*hasdemoplugin.Plugin).Test()
	this.gateway = options.(*core.APIGateWayImplement)
	return nil
}

func (this *Service) EntityStub() *core.EntityStub {
	return core.NewEntityStub(&core.EntityStubOptions{Owner: this})
}

func (this *Service) Config() core.IEntityConf {
	return &this.conf
}
