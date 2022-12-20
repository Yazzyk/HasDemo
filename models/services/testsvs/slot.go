package testsvs

import (
	"github.com/drharryhe/has/common/herrors"
	"github.com/drharryhe/has/common/hlogger"
	"github.com/drharryhe/has/common/htypes"
	"github.com/drharryhe/has/connectors/hwebconnector"
	"github.com/drharryhe/has/core"
)

type TestRequest struct {
	core.SlotRequestBase // 注意继承 core.SlotRequestBase
	// 注意类型要带指针,才能使用data验证
	Name   *string `json:"name" param:"require"`   // 名称
	Change *bool   `json:"change" param:"require"` // 是否改变name param:"require"代表此字段必传
}
// 我们建议在项目开发中，开发者为此处的每个字段都写上注释，便于后期维护

func (this *Service) TestSlot(req *TestRequest, res *core.SlotResponse) {
	hlogger.Info("request name:", *req.Name) // HAS中带有log组件可以直接使用
	hlogger.Info("request change:", *req.Change)
	resultName := *req.Name
	if *req.Change {
		resultName = this.TestName
	}
	if resultName == "" {
		// 定义error debug模式下 New里面的内容会直接以log输出 所以不需要单独打log
		res.Error = herrors.ErrSysInternal.New("名称为空").D("此处信息会返回前端")
		return
	}
	// 返回数据
	this.Response(res, map[string]interface{}{
		"change":     *req.Change,
		"resultName": resultName,
	}, nil)
}


type WsTestRequest struct {
	core.SlotWsRequestBase

	Msg *string `json:"msg" param:"require"`
}

func (this *Service) WsTestSlot(req *WsTestRequest, res *core.SlotResponse) {
	// 创建ws连接
	if *req.INITWS {
		hlogger.Info("初始化连接")
		//res.Error = herrors.ErrSysInternal.New("可返回错误，返回错误会直接断开连接")
		return
	}

	if *req.BREAK {
		hlogger.Info("连接断开")
		return
	}
	// 后续的消息处理
	// ...
	// 向客户端发送消息
	//res.Error = herrors.ErrSysInternal.New("可返回错误，返回错误会直接断开连接")
	// 可通过 CloseWsConn 方法关闭连接
	//this.gateway.Connectors()["web"].(*hwebconnector.Connector).CloseWsConn(*req.WsID)
	this.gateway.Connectors()["web"].(*hwebconnector.Connector).SendWsResponse(*req.WsID, htypes.Map{
		"request_msg": *req.Msg,
	}, nil)
}