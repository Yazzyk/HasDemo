package hasdemoplugin

import (
	"github.com/drharryhe/has/common/herrors"
	"github.com/drharryhe/has/common/hlogger"
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/core"
)

type Plugin struct {
	core.BasePlugin

	conf       HASDemoPlugin // 配置文件
	// ...其它参数可以自定义
}

var plugin = &Plugin{}

// New 调用
func New() *Plugin {
	return plugin
}

// Open 初始化插件
func (this *Plugin) Open(s core.IServer, ins core.IPlugin) *herrors.Error {
	if err := this.BasePlugin.Open(s, ins); err != nil {
		return err
	}
	hlogger.Info("Hello %s Plugin", this.conf.Name)
	return nil
}

// Config 暴露配置文件
func (this *Plugin) Config() core.IEntityConf {
	return &this.conf
}

func (this *Plugin) EntityStub() *core.EntityStub {
	return core.NewEntityStub(
		&core.EntityStubOptions{
			Owner: this,
		})
}

// Capability 暴露其它的参数，自定义
func (this *Plugin) Capability() htypes.Any {
	return map[string]interface{}{}
}

func (this *Plugin) Test() {
	hlogger.Debug("Test Plugin, %s", this.conf.Name)
}