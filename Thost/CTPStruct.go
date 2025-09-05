package Thost

import (
	"github.com/li-wancai/GoScripts/Formulae"

	"github.com/pseudocodes/go2ctp/thost"
)

type CTPBaseCfg struct {
	FlowPath         string
	DataPath         string
	UsingUDP         bool
	Multicast        bool
	UserID           string
	InvestorID       string
	Password         string
	AppID            string
	BrokerID         string
	AuthCode         string
	ProductInfo      string
	TDAddress        []string
	MDAddress        []string
	TradingDay       string
	NoTradeTimeRange []Formulae.TimeRange
}

// 错误返回代码
type ErrReason int

// 枚举常量
const (
	NetworkReadFailed     ErrReason = 0x1001
	NetworkWriteFailed    ErrReason = 0x1002
	HeartbeatRecvTimeout  ErrReason = 0x2001
	HeartbeatSendFailed   ErrReason = 0x2002
	InvalidPacketReceived ErrReason = 0x2003
)

// 错误原因与描述的映射
var ErrReasonMap = map[ErrReason]string{
	NetworkReadFailed:     "网络读失败",
	NetworkWriteFailed:    "网络写失败",
	HeartbeatRecvTimeout:  "接收心跳超时",
	HeartbeatSendFailed:   "发送心跳失败",
	InvalidPacketReceived: "收到错误报文",
}

// 错误返回代码
type RspRst int

// 枚举常量
const (
	Success                     RspRst = 0
	NetworkConnectionError      RspRst = -1
	UnprocessedRequestsExceeded RspRst = -2
	RequestsPerSecondExceeded   RspRst = -3
)

// 请求结果与描述的映射
var RspRstMap = map[RspRst]string{
	Success:                     "成功",
	NetworkConnectionError:      "网络连接失败",
	UnprocessedRequestsExceeded: "未处理请求超过许可数",
	RequestsPerSecondExceeded:   "每秒发送请求数超过许可数",
}

// ToDirection 将字符串类型的买卖方向转换为 TThostFtdcDirectionType 类型
func ToDirection(direction string) thost.TThostFtdcDirectionType {
	switch direction {
	case "Buy":
		return thost.THOST_FTDC_D_Buy
	case "Sell":
		return thost.THOST_FTDC_D_Sell
	default:
		return thost.THOST_FTDC_D_Buy // 默认返回买
	}
}

// ToActionFlag 将字符串类型的操作标志转换为 TThostFtdcActionFlagType 类型
func ToActionFlag(actionFlag string) thost.TThostFtdcActionFlagType {
	switch actionFlag {
	case "Delete":
		return thost.THOST_FTDC_AF_Delete // 删除
	case "Modify":
		return thost.THOST_FTDC_AF_Modify // 修改
	default:
		return thost.THOST_FTDC_AF_Delete // 容错默认 删除
	}
}

// ToTimeCondition 将字符串类型的有效期类型转换为 TThostFtdcTimeConditionType 类型
func ToTimeCondition(timeCondition string) thost.TThostFtdcTimeConditionType {
	// 确定报单有效期类型，目前只支持 GFD 和 IOC ，其他都不支持 例:立即完成，否则撤销 THOST_FTDC_TC_IOC
	switch timeCondition {
	case "IOC":
		return thost.THOST_FTDC_TC_IOC // 立即完成，否则撤销
	case "GFS":
		return thost.THOST_FTDC_TC_GFS // 本节有效
	case "GFD":
		return thost.THOST_FTDC_TC_GFD // 当日有效
	case "GTD":
		return thost.THOST_FTDC_TC_GTD // 指定日期前有效
	case "GTC":
		return thost.THOST_FTDC_TC_GTC // 撤销前有效
	case "GFA":
		return thost.THOST_FTDC_TC_GFA // 集合竞价有效
	default:
		return thost.THOST_FTDC_TC_GFD // 容错默认 撤销前有效
	}
}

// ToVolumeCondition 将字符串类型的成交量类型转换为 TThostFtdcVolumeConditionType 类型
func ToVolumeCondition(volumeCondition string) thost.TThostFtdcVolumeConditionType {
	switch volumeCondition {
	case "AV":
		return thost.THOST_FTDC_VC_AV // 任何数量
	case "MV":
		return thost.THOST_FTDC_VC_MV // 最小数量
	case "CV":
		return thost.THOST_FTDC_VC_CV // 全部数量
	default:
		return thost.THOST_FTDC_VC_CV // 容错默认 全部数量
	}
}

// 定义一个映射表，用于将字符串类型的 OrderPriceType 转换为 TThostFtdcOrderPriceTypeType 类型
var mpOrderPriceType = map[string]thost.TThostFtdcOrderPriceTypeType{
	"AnyPrice":                thost.THOST_FTDC_OPT_AnyPrice,                // 任意价
	"LimitPrice":              thost.THOST_FTDC_OPT_LimitPrice,              // 限价
	"BestPrice":               thost.THOST_FTDC_OPT_BestPrice,               // 最优价
	"LastPrice":               thost.THOST_FTDC_OPT_LastPrice,               // 最新价
	"LastPricePlusOneTicks":   thost.THOST_FTDC_OPT_LastPricePlusOneTicks,   // 最新价浮动上浮1个ticks
	"LastPricePlusTwoTicks":   thost.THOST_FTDC_OPT_LastPricePlusTwoTicks,   // 最新价浮动上浮2个ticks
	"LastPricePlusThreeTicks": thost.THOST_FTDC_OPT_LastPricePlusThreeTicks, // 最新价浮动上浮3个ticks
	"AskPrice1":               thost.THOST_FTDC_OPT_AskPrice1,               // 卖一价
	"AskPrice1PlusOneTicks":   thost.THOST_FTDC_OPT_AskPrice1PlusOneTicks,   // 卖一价浮动上浮1个ticks
	"AskPrice1PlusTwoTicks":   thost.THOST_FTDC_OPT_AskPrice1PlusTwoTicks,   // 卖一价浮动上浮2个ticks
	"AskPrice1PlusThreeTicks": thost.THOST_FTDC_OPT_AskPrice1PlusThreeTicks, // 卖一价浮动上浮3个ticks
	"BidPrice1":               thost.THOST_FTDC_OPT_BidPrice1,               // 买一价
	"BidPrice1PlusOneTicks":   thost.THOST_FTDC_OPT_BidPrice1PlusOneTicks,   // 买一价浮动上浮1个ticks
	"BidPrice1PlusTwoTicks":   thost.THOST_FTDC_OPT_BidPrice1PlusTwoTicks,   // 买一价浮动上浮2个ticks
	"BidPrice1PlusThreeTicks": thost.THOST_FTDC_OPT_BidPrice1PlusThreeTicks, // 买一价浮动上浮3个ticks
	"FiveLevelPrice":          thost.THOST_FTDC_OPT_FiveLevelPrice,          // 五档价
}

func ToOrderPriceType(orderPriceType string) thost.TThostFtdcOrderPriceTypeType {
	if val, ok := mpOrderPriceType[orderPriceType]; ok {
		return val
	}
	return thost.THOST_FTDC_OPT_LastPrice // 容错:默认返回最新价
}

// ToForceCloseReason 将字符串类型的强平原因类型转换为 TThostFtdcForceCloseReasonType 类型
func ToForceCloseReason(forceCloseReason string) thost.TThostFtdcForceCloseReasonType {
	switch forceCloseReason {
	case "NotForceClose":
		return thost.THOST_FTDC_FCC_NotForceClose // 非强平
	case "LackDeposit":
		return thost.THOST_FTDC_FCC_LackDeposit // 资金不足
	case "ClientOverPositionLimit":
		return thost.THOST_FTDC_FCC_ClientOverPositionLimit // 客户超仓
	case "MemberOverPositionLimit":
		return thost.THOST_FTDC_FCC_MemberOverPositionLimit // 会员超仓
	case "NotMultiple":
		return thost.THOST_FTDC_FCC_NotMultiple // 持仓非整数倍
	case "Violation":
		return thost.THOST_FTDC_FCC_Violation // 违规
	case "Other":
		return thost.THOST_FTDC_FCC_Other // 其它
	case "PersonDeliv":
		return thost.THOST_FTDC_FCC_PersonDeliv // 自然人临近交割
	case "Notverifycapital":
		return thost.THOST_FTDC_FCC_Notverifycapital // 风控强平不验证资金
	default:
		return thost.THOST_FTDC_FCC_NotForceClose // 容错默认 非强平
	}
}

// ToContingentCondition 将字符串类型的触发条件类型转换为 TThostFtdcContingentConditionType 类型
func ToContingentCondition(contingentCondition string) thost.TThostFtdcContingentConditionType {
	switch contingentCondition {
	case "Immediately":
		return thost.THOST_FTDC_CC_Immediately // 立即
	case "Touch":
		return thost.THOST_FTDC_CC_Touch // 止损
	case "TouchProfit":
		return thost.THOST_FTDC_CC_TouchProfit // 止赢
	case "ParkedOrder":
		return thost.THOST_FTDC_CC_ParkedOrder // 预埋单
	case "LastPriceGreaterThanStopPrice":
		return thost.THOST_FTDC_CC_LastPriceGreaterThanStopPrice // 最新价大于条件价
	case "LastPriceGreaterEqualStopPrice":
		return thost.THOST_FTDC_CC_LastPriceGreaterEqualStopPrice // 最新价大于等于条件价
	case "LastPriceLesserThanStopPrice":
		return thost.THOST_FTDC_CC_LastPriceLesserThanStopPrice // 最新价小于条件价
	case "LastPriceLesserEqualStopPrice":
		return thost.THOST_FTDC_CC_LastPriceLesserEqualStopPrice // 最新价小于等于条件价
	case "AskPriceGreaterThanStopPrice":
		return thost.THOST_FTDC_CC_AskPriceGreaterThanStopPrice // 卖一价大于条件价
	case "AskPriceGreaterEqualStopPrice":
		return thost.THOST_FTDC_CC_AskPriceGreaterEqualStopPrice // 卖一价大于等于条件价
	case "AskPriceLesserThanStopPrice":
		return thost.THOST_FTDC_CC_AskPriceLesserThanStopPrice // 卖一价小于条件价
	case "AskPriceLesserEqualStopPrice":
		return thost.THOST_FTDC_CC_AskPriceLesserEqualStopPrice // 卖一价小于等于条件价
	case "BidPriceGreaterThanStopPrice":
		return thost.THOST_FTDC_CC_BidPriceGreaterThanStopPrice // 买一价大于条件价
	case "BidPriceGreaterEqualStopPrice":
		return thost.THOST_FTDC_CC_BidPriceGreaterEqualStopPrice // 买一价大于等于条件价
	case "BidPriceLesserThanStopPrice":
		return thost.THOST_FTDC_CC_BidPriceLesserThanStopPrice // 买一价小于条件价
	case "BidPriceLesserEqualStopPrice":
		return thost.THOST_FTDC_CC_BidPriceLesserEqualStopPrice // 买一价小于等于条件价
	default:
		return thost.THOST_FTDC_CC_Immediately // 容错默认 立即
	}
}

// ToOffsetFlag 将字符串类型的偏移标志类型转换为 TThostFtdcOffsetFlagType 类型
func ToOffsetFlag(offsetFlag string) thost.TThostFtdcOffsetFlagType {
	switch offsetFlag {
	case "Open":
		return thost.THOST_FTDC_OF_Open // 开仓
	case "Close":
		return thost.THOST_FTDC_OF_Close // 平仓
	case "ForceClose":
		return thost.THOST_FTDC_OF_ForceClose // 强平
	case "CloseToday":
		return thost.THOST_FTDC_OF_CloseToday // 平今
	case "CloseYesterday":
		return thost.THOST_FTDC_OF_CloseYesterday // 平昨
	case "ForceOff":
		return thost.THOST_FTDC_OF_ForceOff // 强减
	case "LocalForceClose":
		return thost.THOST_FTDC_OF_LocalForceClose // 本地强平
	default:
		return thost.THOST_FTDC_OF_Close // 容错默认 开仓
	}
}

// ToAskHedgeFlag 将字符串类型的对冲标志类型转换为 TThostFtdcHedgeFlagType 类型
func ToHedgeFlag(hedgeFlag string) thost.TThostFtdcHedgeFlagType {
	switch hedgeFlag {
	case "Speculation":
		return thost.THOST_FTDC_HF_Speculation // 投机
	case "Arbitrage":
		return thost.THOST_FTDC_HF_Arbitrage // 套利
	case "Hedge":
		return thost.THOST_FTDC_HF_Hedge // 套保
	case "MarketMaker":
		return thost.THOST_FTDC_HF_MarketMaker // 做市商
	case "SpecHedge":
		return thost.THOST_FTDC_HF_SpecHedge // 第一腿投机第二腿套保
	case "HedgeSpec":
		return thost.THOST_FTDC_HF_HedgeSpec // 第一腿套保第二腿投机
	default:
		return thost.THOST_FTDC_HF_Speculation // 容错默认 投机
	}
}
