package CTPBase

import (
	"fmt"

	"github.com/li-wancai/EQCTPQH/Thost"

	"github.com/li-wancai/GoScripts/DirsFile"

	"github.com/li-wancai/GoScripts/Formulae"

	"github.com/pseudocodes/go2ctp/ctp"
	"github.com/pseudocodes/go2ctp/thost"
)

// 错误请求处理
func ErrorRspInfo(pRspInfo *thost.CThostFtdcRspInfoField) bool {
	if pRspInfo == nil {
		return false
	}
	bResult := (pRspInfo.ErrorID != 0)
	if bResult {
		log.Warnf("错误代码: %d, 错误信息: %s",
			pRspInfo.ErrorID, pRspInfo.ErrorMsg.GBString())
	}
	return bResult
}
func ErrReason(nReason int) {
	reason := Thost.ErrReason(nReason) // 将int转换为自定义枚举类型
	reasonDesc, exists := Thost.ErrReasonMap[reason]
	if exists { // 查找错误原因对应的中文描述
		log.Errorf("【错误】: %s (代码: %d)", reasonDesc, reason)
	} else {
		log.Errorf("【未知错误】: (代码: %d)", reason)
	}
}

func RspRst(nReason int, msg string, v ...interface{}) {
	msg = fmt.Sprintf(msg, v...)
	reason := Thost.RspRst(nReason) // 将int转换为自定义枚举类型
	reasonDesc, exists := Thost.RspRstMap[reason]
	if exists { // 查找错误原因对应的中文描述
		if reason == Thost.Success {
			log.Infof("%s【完成】: %s (代码: %d)", msg, reasonDesc, reason)
		} else {
			log.Errorf("%s【错误】: %s (代码: %d)", msg, reasonDesc, reason)
		}

	} else {
		log.Errorf("%s【未知】: (代码: %d)", msg, reason)
	}
}

// 设置CTPConfig配置
func CTPConfig(keyname string, filename string, path string) Thost.CTPBaseCfg {
	// 如果未提供文件名，则使用默认值"MySqlDB"
	if filename == "" {
		filename = "CTPBase"
	}
	// 如果未提供路径，则使用默认路径"./Datas/etc/"
	if path == "" {
		path = "./Datas/etc/"
	}
	// 通过ReadToml函数从指定路径的配置文件中读取配置信息
	cfgall, err := DirsFile.ReadToml(filename, path)
	if err != nil {
		return Thost.CTPBaseCfg{}
	}
	cfgget, ok := cfgall[keyname].(map[string]interface{})
	if !ok {
		return Thost.CTPBaseCfg{}
	}
	config := Thost.CTPBaseCfg{
		FlowPath:         cfgget["flowPath"].(string),
		DataPath:         cfgget["dataPath"].(string),
		UsingUDP:         cfgget["usingUDP"].(bool),
		Multicast:        cfgget["multicast"].(bool),
		UserID:           cfgget["userid"].(string),
		InvestorID:       cfgget["investorId"].(string),
		Password:         cfgget["password"].(string),
		AppID:            cfgget["appid"].(string),
		BrokerID:         cfgget["brokerid"].(string),
		AuthCode:         cfgget["auth_code"].(string),
		ProductInfo:      cfgget["product_info"].(string),
		TDAddress:        Formulae.ToStringList(cfgget["td_address"].([]interface{})),
		MDAddress:        Formulae.ToStringList(cfgget["md_address"].([]interface{})),
		NoTradeTimeRange: Formulae.TradeTimeRanges(Formulae.ListStrList(cfgget["NoTradeTimeRange"].([]interface{}))),
	}
	return config
}

type MDApiN struct {
	Api       thost.MdApi
	Cfg       Thost.CTPBaseCfg
	ApiPtr    uintptr
	RequestID int
}

// 初始化
func InitMDApi(config map[string]interface{}, KeyName string) *MDApiN {
	if KeyName == "" {
		KeyName = config["EQCTPCfg_Account"].(string)
	}
	CTPCfg := CTPConfig(
		KeyName,
		config["EQCTPCfg_FileName"].(string),
		config["EQCTPCfg_TomlPath"].(string))
	return CreateMDApi(CTPCfg)
}

// 创建MdApi
func CreateMDApi(cfg Thost.CTPBaseCfg) *MDApiN {
	MDapi := &MDApiN{
		Cfg:       cfg,
		RequestID: 1,
	}
	MDapi.Api = ctp.CreateMdApi(ctp.MdFlowPath(cfg.FlowPath), ctp.MdUsingUDP(cfg.UsingUDP), ctp.MdMultiCast(cfg.Multicast))
	return MDapi
}

// 获取配置信息
func (s *MDApiN) GetCfg() Thost.CTPBaseCfg {
	return s.Cfg
}

// 盘中死循环运行
func CTPBreakRun(Cfg Thost.CTPBaseCfg, Spantime int /*主循环等待的时间秒*/, EndSpantime int /*非交易日时间段的主循环等待退出前等待的时间秒*/) {
	for {
		Formulae.Sleep(Spantime)
		if Formulae.NoTradingTime(Cfg.NoTradeTimeRange) {
			Formulae.Sleep(EndSpantime)
			return
		}
	}
}
