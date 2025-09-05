package CTPBase

import "C"
import (
	"time"

	"github.com/li-wancai/EQCTPQH/Thost"

	"github.com/pseudocodes/go2ctp/ctp"
	"github.com/pseudocodes/go2ctp/thost"
)

type TDApiN struct {
	apiPtr     uintptr
	Api        thost.TraderApi
	Cfg        Thost.CTPBaseCfg
	length     int
	RequestID  int
	ReqTime    time.Time
	systemInfo [273]byte
}

// 初始化
func InitTDApi(config map[string]interface{}, KeyName string /*数据key前缀*/) *TDApiN {
	if KeyName == "" {
		KeyName = config["EQCTPCfg_Account"].(string)
	}
	CTPCfg := CTPConfig(
		KeyName,
		config["EQCTPCfg_FileName"].(string),
		config["EQCTPCfg_TomlPath"].(string))
	return CreateTDApi(CTPCfg)
}

// 创建TDApi
func CreateTDApi(cfg Thost.CTPBaseCfg) *TDApiN {
	TDApi := &TDApiN{
		Cfg:       cfg,
		RequestID: 1,
	}
	TDApi.Api = ctp.CreateTraderApi(ctp.TraderFlowPath(cfg.FlowPath))
	return TDApi
}

// 获取配置信息
func (s *TDApiN) GetCfg() Thost.CTPBaseCfg {
	return s.Cfg
}
