package main

import (
	"HASDemo/models/objs"
	"HASDemo/models/services/testsvs"
	"github.com/drharryhe/has/connectors/hwebconnector"
	"github.com/drharryhe/has/core"
	"github.com/drharryhe/has/datapackers/hjsonpacker"
	"github.com/drharryhe/has/plugins/hdatabaseplugin"
	"github.com/drharryhe/has/routers/hlocalrouter"
	"github.com/drharryhe/has/services/hdatasvs"
	"github.com/drharryhe/has/services/hellosvs"
)

func main() {
	// 创建gateway
	gateway := core.NewAPIGateway(&core.APIGatewayOptions{
		// Server配置
		ServerOptions: core.ServerOptions{
			// 路由配置
			Router: hlocalrouter.New(),
			Plugins: []core.IPlugin{
				hdatabaseplugin.New(),
			},
		},
		Connectors: []core.IAPIConnector{
			hwebconnector.New(nil), // Web服务
		},
		Packers: []core.IAPIDataPacker{
			hjsonpacker.New(),
		},
	})
	// 注册服务
	gateway.Server().RegisterService(&hellosvs.Service{}, nil)
	gateway.Server().RegisterService(&testsvs.Service{}, nil)
	gateway.Server().RegisterService(&hdatasvs.Service{}, &hdatasvs.Options{
		Hooks:        nil,
		Views:        nil,
		Objects:      objs.Objects(),
		FieldFuncMap: nil,
	})
	// 启动服务
	gateway.Start()
}
