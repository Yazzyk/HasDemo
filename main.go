package main

import (
	"HASDemo/hooks"
	"HASDemo/middleware"
	"HASDemo/middleware/testmiddleware"
	"HASDemo/models/objs"
	"HASDemo/models/services/testsvs"
	"HASDemo/plugin/hasdemoplugin"
	"github.com/drharryhe/has/connectors/hwebconnector"
	"github.com/drharryhe/has/core"
	"github.com/drharryhe/has/datapackers/hjsonpacker"
	"github.com/drharryhe/has/middlewares/hpermmw"
	"github.com/drharryhe/has/plugins/hdatabaseplugin"
	"github.com/drharryhe/has/plugins/hmemcacheplugin"
	"github.com/drharryhe/has/routers/hlocalrouter"
	"github.com/drharryhe/has/services/hapauthsvs"
	"github.com/drharryhe/has/services/hdatasvs"
	"github.com/drharryhe/has/services/hellosvs"
	"github.com/drharryhe/has/services/hfilesvs"
	"github.com/drharryhe/has/services/hsessionsvs"
)

func main() {
	// 创建 gateway
	gateway := core.NewAPIGateway(&core.APIGatewayOptions{
		// Server配置
		ServerOptions: core.ServerOptions{
			// 路由配置
			Router: hlocalrouter.New(),
			Plugins: []core.IPlugin{
				hdatabaseplugin.New(),
				hmemcacheplugin.New(),
				hasdemoplugin.New(),
			},
		},
		Connectors: []core.IAPIConnector{
			hwebconnector.New(), // Web服务
		},
		Middlewares: []core.IAPIMiddleware{
			hpermmw.New(middleware.NewPermFuncWrapper()),
			testmiddleware.New(),
		},
		Packers: []core.IAPIDataPacker{
			hjsonpacker.New(),
		},
	})
	// 注册服务
	gateway.Server().RegisterService(&hellosvs.Service{}, nil)
	gateway.Server().RegisterService(&hsessionsvs.Service{}, nil)
	gateway.Server().RegisterService(&testsvs.Service{}, gateway)
	gateway.Server().RegisterService(&hdatasvs.Service{}, &hdatasvs.Options{
		Hooks:        nil,
		Views:        objs.Views(),
		Objects:      objs.Objects(),
		FieldFuncMap: nil,
	})
	gateway.Server().RegisterService(&hapauthsvs.Service{}, &hapauthsvs.Options{
		Hooks: &hooks.ApHooks{},
		PasswordEncoder: func(pwd string) string {
			return "HAS_" + pwd
		},
	})
	gateway.Server().RegisterService(&hfilesvs.Service{}, nil)
	// 启动服务
	gateway.Start()
}
