package CTPBase

import (
	"github.com/li-wancai/EQCTPQH/Thost"

	"github.com/li-wancai/GoScripts/Formulae"

	"github.com/pseudocodes/go2ctp/thost"
)

/*
ReqQuoteInsert
报价录入请求，
如果出错，则返回响应OnRspQuoteInsert和OnErrRtnQuoteInsert；正确则推送OnRtnQuote、OnRtnOrder和OnRtnTrade。
单边报价和双边报价，都是用一个接口 ReqQuoteInsert。
在单边报价的时候，只需要另一边的数量填0，交易核心就能区分开。另外，无论是单边还是双边，Ask/BidOrderRef都是要填的。
除上期所的期货合约使用 ReqOrderInsert 接口报价，其他交易所均使用本接口报价。
*/

// 报单录入
func (s *TDSpiN) DoReqQuoteInsert(
	IPAddress /*IP地址*/, MacAddress, /*Mac地址*/
	InstrumentID /*合约代码*/, ExchangeID, /*交易所代码*/
	AskOffsetFlag, BidOffsetFlag,
	AskHedgeFlag, BidHedgeFlag string,
	AskVolume, BidVolume int,
	AskPrice, BidPrice float64) int {
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	s.RequestID++
	req := &thost.CThostFtdcInputQuoteField{}
	copy(req.UserID[:], s.Cfg.UserID)
	copy(req.IPAddress[:], IPAddress)
	copy(req.MacAddress[:], MacAddress)
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.InvestorID[:], s.Cfg.InvestorID)
	copy(req.ExchangeID[:], ExchangeID)     // 交易所代码
	copy(req.InstrumentID[:], InstrumentID) // 合约代码

	req.AskPrice = thost.TThostFtdcPriceType(AskPrice)
	req.BidPrice = thost.TThostFtdcPriceType(BidPrice)
	req.AskVolume = thost.TThostFtdcVolumeType(AskVolume)
	req.BidVolume = thost.TThostFtdcVolumeType(BidVolume)
	req.AskHedgeFlag = Thost.ToHedgeFlag(AskHedgeFlag)    //卖投机套保标志
	req.BidHedgeFlag = Thost.ToHedgeFlag(BidHedgeFlag)    //买投机套保标志
	req.AskOffsetFlag = Thost.ToOffsetFlag(AskOffsetFlag) //卖开平标志
	req.BidOffsetFlag = Thost.ToOffsetFlag(BidOffsetFlag) //买开平标志
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	rst := s.Api.ReqQuoteInsert(req, s.RequestID) // 报单录入请求
	RspRst(rst, "【%s】【报单录入】请求 ReqQuoteInsert 已发送", s.Cfg.UserID)
	return rst
}

// DoReqOrderInsert 执行报单录入请求。
// 该函数负责根据提供的参数构建报单请求，并将其发送到交易所。
// 参数包括合约代码、交易所代码、开平标志、投机套保标志、报单价格条件、买卖方向、有效期类型、成交量类型、触发条件、强平原因，
// 交易价格、止损价、交易数量、自动挂起标志和互换单标志。
// 返回值为请求的结果。
func (s *TDSpiN) DoReqOrderInsert(
	IPAddress /*IP地址*/, MacAddress, /*Mac地址*/
	InstrumentID /*合约代码*/, ExchangeID, /*交易所代码*/
	CombOffsetFlag /*开平标志*/, CombHedgeFlag, /*投机套保标志*/
	OrderPriceType /*报单价格条件*/, Direction /*买卖方向*/, TimeCondition, /*有效期类型*/
	VolumeCondition /*成交量类型*/, ContingentCondition /*触发条件*/, ForceCloseReason /*强平原因*/ string,
	LimitPrice /*交易价格*/, StopPrice /*止损价*/ float64, VolumeTotalOriginal /*交易数量*/, IsAutoSuspend /*自动挂起标志*/, IsSwapOrder /*互换单标志*/ int) int {
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	s.RequestID++
	req := &thost.CThostFtdcInputOrderField{}
	copy(req.UserID[:], s.Cfg.UserID)
	copy(req.IPAddress[:], IPAddress)
	copy(req.MacAddress[:], MacAddress)
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.InvestorID[:], s.Cfg.InvestorID)
	copy(req.ExchangeID[:], ExchangeID)         // 交易所代码
	copy(req.InstrumentID[:], InstrumentID)     // 合约代码
	copy(req.CombHedgeFlag[:], CombHedgeFlag)   // 投机套保标志 注：郑商所平仓只能报入“投机”属性 例:投机
	copy(req.CombOffsetFlag[:], CombOffsetFlag) // 开平标志

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	req.Direction = Thost.ToDirection(Direction)                               // 买卖方向 容错默认 买 Buy Sell
	req.StopPrice = thost.TThostFtdcPriceType(StopPrice)                       // 止损价
	req.LimitPrice = thost.TThostFtdcPriceType(Formulae.Round(LimitPrice, 2))  // 交易价格
	req.IsSwapOrder = thost.TThostFtdcBoolType(IsSwapOrder)                    // 互换单标志
	req.TimeCondition = Thost.ToTimeCondition(TimeCondition)                   // 有效期类型 目前只支持 GFD 当日有效 和 IOC 立即完成，否则撤销
	req.IsAutoSuspend = thost.TThostFtdcBoolType(IsAutoSuspend)                // 自动挂起标志
	req.OrderPriceType = Thost.ToOrderPriceType(OrderPriceType)                // 报单价格条件 容错默认 最新价
	req.VolumeCondition = Thost.ToVolumeCondition(VolumeCondition)             // 成交量类型 容错默认CV //"AV"任何数量"MV"最小数量"CV"全部数量
	req.VolumeTotalOriginal = thost.TThostFtdcVolumeType(VolumeTotalOriginal)  // 交易数量
	req.ContingentCondition = Thost.ToContingentCondition(ContingentCondition) // 触发条件 容错默认 Immediately立即
	req.ForceCloseReason = Thost.ToForceCloseReason(ForceCloseReason)          // 强平原因
	req.MinVolume = thost.TThostFtdcVolumeType(1)
	req.RequestID = thost.TThostFtdcRequestIDType(s.RequestID)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	rst := s.Api.ReqOrderInsert(req, s.RequestID) // 报单录入请求
	RspRst(rst, "【%s】【报单录入】请求 ReqOrderInsert 已发送", s.Cfg.UserID)
	return rst
}

/*
ToHedgeFlag
"Speculation":
    THOST_FTDC_HF_Speculation // 投机
"Arbitrage":
    THOST_FTDC_HF_Arbitrage // 套利
"Hedge":
    THOST_FTDC_HF_Hedge // 套保
"MarketMaker":
    THOST_FTDC_HF_MarketMaker // 做市商
"SpecHedge":
    THOST_FTDC_HF_SpecHedge // 第一腿投机第二腿套保
"HedgeSpec":
    THOST_FTDC_HF_HedgeSpec // 第一腿套保第二腿投机
*/

/*
ToOffsetFlag
"Open":
	THOST_FTDC_OF_Open // 开仓
"Close":
	THOST_FTDC_OF_Close // 平仓
"ForceClose":
	THOST_FTDC_OF_ForceClose // 强平
"CloseToday":
	THOST_FTDC_OF_CloseToday // 平今
"CloseYesterday":
	THOST_FTDC_OF_CloseYesterday // 平昨
"ForceOff":
	THOST_FTDC_OF_ForceOff // 强减
"LocalForceClose":
	THOST_FTDC_OF_LocalForceClose // 本地强平
*/

/*
CombOffsetFlag：确定开平标志。
上期所和能源交易所支持平昨和平仓指令，分别对应昨仓和今仓； 如果下平仓指令，则等同于平昨指令。
大商所、广期所、郑商所、中金所只能报入平仓指令，不支持平昨和平今。
*/

/*
ToOrderPriceType
"AnyPrice":                THOST_FTDC_OPT_AnyPrice,                // 任意价
"LimitPrice":              THOST_FTDC_OPT_LimitPrice,              // 限价
"BestPrice":               THOST_FTDC_OPT_BestPrice,               // 最优价
"LastPrice":               THOST_FTDC_OPT_LastPrice,               // 最新价
"LastPricePlusOneTicks":   THOST_FTDC_OPT_LastPricePlusOneTicks,   // 最新价浮动上浮1个ticks
"LastPricePlusTwoTicks":   THOST_FTDC_OPT_LastPricePlusTwoTicks,   // 最新价浮动上浮2个ticks
"LastPricePlusThreeTicks": THOST_FTDC_OPT_LastPricePlusThreeTicks, // 最新价浮动上浮3个ticks
"AskPrice1":               THOST_FTDC_OPT_AskPrice1,               // 卖一价
"AskPrice1PlusOneTicks":   THOST_FTDC_OPT_AskPrice1PlusOneTicks,   // 卖一价浮动上浮1个ticks
"AskPrice1PlusTwoTicks":   THOST_FTDC_OPT_AskPrice1PlusTwoTicks,   // 卖一价浮动上浮2个ticks
"AskPrice1PlusThreeTicks": THOST_FTDC_OPT_AskPrice1PlusThreeTicks, // 卖一价浮动上浮3个ticks
"BidPrice1":               THOST_FTDC_OPT_BidPrice1,               // 买一价
"BidPrice1PlusOneTicks":   THOST_FTDC_OPT_BidPrice1PlusOneTicks,   // 买一价浮动上浮1个ticks
"BidPrice1PlusTwoTicks":   THOST_FTDC_OPT_BidPrice1PlusTwoTicks,   // 买一价浮动上浮2个ticks
"BidPrice1PlusThreeTicks": THOST_FTDC_OPT_BidPrice1PlusThreeTicks, // 买一价浮动上浮3个ticks
"FiveLevelPrice":          THOST_FTDC_OPT_FiveLevelPrice,          // 五档价
*/

/*
ToContingentCondition
"Immediately":
	THOST_FTDC_CC_Immediately // 立即
"Touch":
	THOST_FTDC_CC_Touch // 止损
"TouchProfit":
	THOST_FTDC_CC_TouchProfit // 止赢
"ParkedOrder":
	THOST_FTDC_CC_ParkedOrder // 预埋单
"LastPriceGreaterThanStopPrice":
	THOST_FTDC_CC_LastPriceGreaterThanStopPrice // 最新价大于条件价
"LastPriceGreaterEqualStopPrice":
	THOST_FTDC_CC_LastPriceGreaterEqualStopPrice // 最新价大于等于条件价
"LastPriceLesserThanStopPrice":
	THOST_FTDC_CC_LastPriceLesserThanStopPrice // 最新价小于条件价
"LastPriceLesserEqualStopPrice":
	THOST_FTDC_CC_LastPriceLesserEqualStopPrice // 最新价小于等于条件价
"AskPriceGreaterThanStopPrice":
	THOST_FTDC_CC_AskPriceGreaterThanStopPrice // 卖一价大于条件价
"AskPriceGreaterEqualStopPrice":
	THOST_FTDC_CC_AskPriceGreaterEqualStopPrice // 卖一价大于等于条件价
"AskPriceLesserThanStopPrice":
	THOST_FTDC_CC_AskPriceLesserThanStopPrice // 卖一价小于条件价
"AskPriceLesserEqualStopPrice":
	THOST_FTDC_CC_AskPriceLesserEqualStopPrice // 卖一价小于等于条件价
"BidPriceGreaterThanStopPrice":
	THOST_FTDC_CC_BidPriceGreaterThanStopPrice // 买一价大于条件价
"BidPriceGreaterEqualStopPrice":
	THOST_FTDC_CC_BidPriceGreaterEqualStopPrice // 买一价大于等于条件价
"BidPriceLesserThanStopPrice":
	THOST_FTDC_CC_BidPriceLesserThanStopPrice // 买一价小于条件价
"BidPriceLesserEqualStopPrice":
	THOST_FTDC_CC_BidPriceLesserEqualStopPrice // 买一价小于等于条件价

*/

/*
ToForceCloseReason
"NotForceClose":
    THOST_FTDC_FCC_NotForceClose // 非强平
"LackDeposit":
    THOST_FTDC_FCC_LackDeposit // 资金不足
"ClientOverPositionLimit":
    THOST_FTDC_FCC_ClientOverPositionLimit // 客户超仓
"MemberOverPositionLimit":
    THOST_FTDC_FCC_MemberOverPositionLimit // 会员超仓
"NotMultiple":
    THOST_FTDC_FCC_NotMultiple // 持仓非整数倍
"Violation":
    THOST_FTDC_FCC_Violation // 违规
"Other":
    THOST_FTDC_FCC_Other // 其它
"PersonDeliv":
    THOST_FTDC_FCC_PersonDeliv // 自然人临近交割
"Notverifycapital":
    THOST_FTDC_FCC_Notverifycapital // 风控强平不验证资金
*/
