package CTPBase

import "github.com/pseudocodes/go2ctp/thost"

// 报单通知
func (s *TDSpiN) OnRtnOrder(pOrder *thost.CThostFtdcOrderField) {
	data := map[string]interface{}{
		"请求编号":         pOrder.RequestID,
		"最小成交量":        pOrder.MinVolume,            // 最小成交量
		"今成交数量":        pOrder.VolumeTraded,         // 今成交数量
		"剩余数量":         pOrder.VolumeTotal,          // 剩余数量
		"前置编号":         pOrder.FrontID,              // 前置编号
		"会话编号":         pOrder.SessionID,            // 会话编号
		"数量":           pOrder.VolumeTotalOriginal,  // 数量
		"安装编号":         pOrder.InstallID,            // 安装编号
		"结算编号":         pOrder.SettlementID,         // 结算编号
		"报单编号":         pOrder.OrderSysID,           // 报单编号
		"报单提示序号":       pOrder.NotifySequence,       // 报单提示序号
		"郑商所成交数量":      pOrder.ZCETotalTradedVolume, // 郑商所成交数量
		"用户强评标志":       pOrder.UserForceClose,       // 用户强评标志
		"互换单标志":        pOrder.IsSwapOrder,          // 互换单标志
		"自动挂起标志":       pOrder.IsAutoSuspend,        // 自动挂起标志
		"经纪公司报单编号":     pOrder.BrokerOrderSeq,       // 经纪公司报单编号
		"GTD日期":        pOrder.GTDDate.String(),     // GTD日期
		"MacAddress":   pOrder.MacAddress.String(),
		"IP地址":         pOrder.IPAddress.String(),
		"InstrumentID": pOrder.InstrumentID.String(),
		"投资者代码":        pOrder.InvestorID.String(),
		"交易所代码":        pOrder.ExchangeID.String(),
		"投资单元代码":       pOrder.InvestUnitID.String(),
		"经纪公司代码":       pOrder.BrokerID.String(),
		"用户代码":         pOrder.UserID.String(),
		"业务单元":         pOrder.BusinessUnit.String(),
		"交易编码":         pOrder.ClientID.String(),
		"报单引用":         pOrder.OrderRef.String(),            // 报单引用
		"报单价格条件":       pOrder.OrderPriceType.String(),      // 报单价格条件
		"买卖方向":         pOrder.Direction.String(),           // 买卖方向
		"组合开平标志":       pOrder.CombOffsetFlag.String(),      // 组合开平标志
		"组合投机套保标志":     pOrder.CombHedgeFlag.String(),       // 组合投机套保标志
		"价格":           pOrder.LimitPrice.String(),          // 价格
		"有效期类型":        pOrder.TimeCondition.String(),       // 有效期类型
		"成交量类型":        pOrder.VolumeCondition.String(),     // 成交量类型
		"触发条件":         pOrder.ContingentCondition.String(), // 触发条件
		"止损价":          pOrder.StopPrice.String(),           // 止损价
		"强平原因":         pOrder.ForceCloseReason.String(),    // 强平原因
		"本地报单编号":       pOrder.OrderLocalID.String(),        // 本地报单编号
		"会员代码":         pOrder.ParticipantID.String(),       // 会员代码
		"客户代码":         pOrder.ClientID.String(),            // 客户代码
		"交易所交易员代码":     pOrder.TraderID.String(),            // 交易所交易员代码
		"报单提交状态":       pOrder.OrderSubmitStatus.String(),   // 报单提交状态
		"用户端产品信息":      pOrder.UserProductInfo.String(),     // 用户端产品信息
		"状态信息":         pOrder.StatusMsg.GBString(),         // 状态信息
		"交易日":          pOrder.TradingDay.String(),          // 交易日
		"报单来源":         pOrder.OrderSource.String(),         // 报单来源
		"报单状态":         pOrder.OrderStatus.String(),         // 报单状态
		"报单类型":         pOrder.OrderType.String(),           // 报单类型
		"报单日期":         pOrder.InsertDate.String(),          // 报单日期
		"委托时间":         pOrder.InsertTime.String(),          // 委托时间
		"激活时间":         pOrder.ActiveTime.String(),          // 激活时间
		"挂起时间":         pOrder.SuspendTime.String(),         // 挂起时间
		"最后修改时间":       pOrder.UpdateTime.String(),          // 最后修改时间
		"撤销时间":         pOrder.CancelTime.String(),          // 撤销时间
		"最后修改交易所交易员代码": pOrder.ActiveTraderID.String(),      // 最后修改交易所交易员代码
		"结算会员编号":       pOrder.ClearingPartID.String(),      // 结算会员编号
		"相关报单":         pOrder.RelativeOrderSysID.String(),  // 相关报单
		"营业部编号":        pOrder.BranchID.String(),            // 营业部编号
		"资金账号":         pOrder.AccountID.String(),           // 资金账号
		"币种代码":         pOrder.CurrencyID.String(),          // 币种代码
		"操作用户代码":       pOrder.ActiveUserID.String(),        // 操作用户代码
	}
	s.OnData(data, true, "OnRtnOrder")
}

// 成交通知
func (s *TDSpiN) OnRtnTrade(pTrade *thost.CThostFtdcTradeField) {
	tradeData := map[string]interface{}{
		"InstrumentID": pTrade.InstrumentID.String(),   // 合约代码
		"投资者代码":        pTrade.InvestorID.String(),     // 投资者代码
		"报单引用":         pTrade.OrderRef.String(),       // 报单引用
		"用户代码":         pTrade.UserID.String(),         // 用户代码
		"交易所代码":        pTrade.ExchangeID.String(),     // 交易所代码
		"成交编号":         pTrade.TradeID.String(),        // 成交编号
		"买卖方向":         pTrade.Direction.String(),      // 买卖方向
		"报单编号":         pTrade.OrderSysID.String(),     // 报单编号
		"会员代码":         pTrade.ParticipantID.String(),  // 会员代码
		"客户代码":         pTrade.ClientID.String(),       // 客户代码
		"交易角色":         pTrade.TradingRole.String(),    // 交易角色
		"开平标志":         pTrade.OffsetFlag.String(),     // 开平标志
		"投机套保标志":       pTrade.HedgeFlag.String(),      // 投机套保标志
		"价格":           pTrade.Price.String(),          // 价格
		"数量":           pTrade.Volume,                  // 数量
		"成交时期":         pTrade.TradeDate.String(),      // 成交时期
		"成交时间":         pTrade.TradeTime.String(),      // 成交时间
		"成交类型":         pTrade.TradeType.String(),      // 成交类型
		"成交价来源":        pTrade.PriceSource.String(),    // 成交价来源
		"交易所交易员代码":     pTrade.TraderID.String(),       // 交易所交易员代码
		"本地报单编号":       pTrade.OrderLocalID.String(),   // 本地报单编号
		"结算会员编号":       pTrade.ClearingPartID.String(), // 结算会员编号
		"业务单元":         pTrade.BusinessUnit.String(),   // 业务单元
		"序号":           pTrade.SequenceNo,              // 序号
		"交易日":          pTrade.TradingDay.String(),     // 交易日
		"结算编号":         pTrade.SettlementID,            // 结算编号
		"经纪公司报单编号":     pTrade.BrokerOrderSeq,          // 经纪公司报单编号
		"成交来源":         pTrade.TradeSource.String(),    // 成交来源
		"投资单元代码":       pTrade.InvestUnitID.String(),   // 投资单元代码
		"合约代码":         pTrade.InstrumentID.String(),   // 合约代码
		"合约在交易所的代码":    pTrade.ExchangeInstID.String(), // 合约在交易所的代码
	}

	s.OnData(tradeData, true, "OnRtnTrade")
}

// 报单录入错误回报
func (s *TDSpiN) OnErrRtnOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if pInputOrder == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "OnErrRtnOrderInsert")
		return
	}
	data := map[string]interface{}{
		"InstrumentID": pInputOrder.InstrumentID.String(),
		"投资者代码":        pInputOrder.InvestorID.String(),
		"交易所代码":        pInputOrder.ExchangeID.String(),
		"投资单元代码":       pInputOrder.InvestUnitID.String(),
		"经纪公司代码":       pInputOrder.BrokerID.String(),
		"报单引用":         pInputOrder.OrderRef.String(),
		"用户代码":         pInputOrder.UserID.String(),
		"报单价格条件":       pInputOrder.OrderPriceType.String(),
		"买卖方向":         pInputOrder.Direction.String(),
		"组合开平标志":       pInputOrder.CombOffsetFlag.String(),
		"组合投机套保标志":     pInputOrder.CombHedgeFlag.String(),
		"价格":           pInputOrder.LimitPrice.String(),
		"数量":           pInputOrder.VolumeTotalOriginal,
		"有效期类型":        pInputOrder.TimeCondition.String(),
		"GTD日期":        pInputOrder.GTDDate.String(),
		"成交量类型":        pInputOrder.VolumeCondition.String(),
		"最小成交量":        pInputOrder.MinVolume,
		"触发条件":         pInputOrder.ContingentCondition.String(),
		"止损价":          pInputOrder.StopPrice.String(),
		"强平原因":         pInputOrder.ForceCloseReason.String(),
		"自动挂起标志":       pInputOrder.IsAutoSuspend,
		"业务单元":         pInputOrder.BusinessUnit.String(),
		"请求编号":         pInputOrder.RequestID,
		"用户强评标志":       pInputOrder.UserForceClose,
		"互换单标志":        pInputOrder.IsSwapOrder,
		"资金账号":         pInputOrder.AccountID.String(),
		"币种代码":         pInputOrder.CurrencyID.String(),
		"交易编码":         pInputOrder.ClientID.String(),
		"Mac地址":        pInputOrder.MacAddress.String(),
		"IP地址":         pInputOrder.IPAddress.String(),
	}
	s.OnData(data, true, "OnErrRtnOrderInsert")
	ErrorRspInfo(pRspInfo)
}

// 报单操作错误回报
func (s *TDSpiN) OnErrRtnOrderAction(pOrderAction *thost.CThostFtdcOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if pOrderAction == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "OnErrRtnOrderAction")
		return
	}
	data := map[string]interface{}{
		"经纪公司代码":       pOrderAction.BrokerID.String(),
		"投资者代码":        pOrderAction.InvestorID.String(),
		"报单操作引用":       pOrderAction.OrderActionRef,
		"报单引用":         pOrderAction.OrderRef.String(),
		"请求编号":         pOrderAction.RequestID,
		"前置编号":         pOrderAction.FrontID,
		"会话编号":         pOrderAction.SessionID,
		"交易所代码":        pOrderAction.ExchangeID.String(),
		"报单编号":         pOrderAction.OrderSysID.String(),
		"操作标志":         pOrderAction.ActionFlag.String(),
		"价格":           pOrderAction.LimitPrice.String(),
		"数量变化":         pOrderAction.VolumeChange,
		"操作日期":         pOrderAction.ActionDate.String(),
		"操作时间":         pOrderAction.ActionTime.String(),
		"交易所交易员代码":     pOrderAction.TraderID.String(),
		"安装编号":         pOrderAction.InstallID,
		"本地报单编号":       pOrderAction.OrderLocalID.String(),
		"操作本地编号":       pOrderAction.ActionLocalID.String(),
		"会员代码":         pOrderAction.ParticipantID.String(),
		"客户代码":         pOrderAction.ClientID.String(),
		"业务单元":         pOrderAction.BusinessUnit.String(),
		"报单操作状态":       pOrderAction.OrderActionStatus.String(),
		"用户代码":         pOrderAction.UserID.String(),
		"状态信息":         pOrderAction.StatusMsg.String(),
		"营业部编号":        pOrderAction.BranchID.String(),
		"投资单元代码":       pOrderAction.InvestUnitID.String(),
		"Mac地址":        pOrderAction.MacAddress.String(),
		"InstrumentID": pOrderAction.InstrumentID.String(),
		"IP地址":         pOrderAction.IPAddress.String(),
	}
	s.OnData(data, true, "OnErrRtnOrderAction")
	ErrorRspInfo(pRspInfo)
}

// 合约交易状态通知
func (s *TDSpiN) OnRtnInstrumentStatus(pInstrumentStatus *thost.CThostFtdcInstrumentStatusField) {
	if s.OnRtnInstrumentStatusCallback != nil {
		s.OnRtnInstrumentStatusCallback(pInstrumentStatus)
	}
}

// 交易所公告通知
func (s *TDSpiN) OnRtnBulletin(pBulletin *thost.CThostFtdcBulletinField) {
	if s.OnRtnBulletinCallback != nil {
		s.OnRtnBulletinCallback(pBulletin)
	}
}

// 交易通知
func (s *TDSpiN) OnRtnTradingNotice(pTradingNoticeInfo *thost.CThostFtdcTradingNoticeInfoField) {
	if s.OnRtnTradingNoticeCallback != nil {
		s.OnRtnTradingNoticeCallback(pTradingNoticeInfo)
	}
}

// 提示条件单校验错误
func (s *TDSpiN) OnRtnErrorConditionalOrder(pErrorConditionalOrder *thost.CThostFtdcErrorConditionalOrderField) {
	if s.OnRtnErrorConditionalOrderCallback != nil {
		s.OnRtnErrorConditionalOrderCallback(pErrorConditionalOrder)
	}
}

// 执行宣告通知
func (s *TDSpiN) OnRtnExecOrder(pExecOrder *thost.CThostFtdcExecOrderField) {
	if s.OnRtnExecOrderCallback != nil {
		s.OnRtnExecOrderCallback(pExecOrder)
	}
}

// 执行宣告录入错误回报
func (s *TDSpiN) OnErrRtnExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnExecOrderInsertCallback != nil {
		s.OnErrRtnExecOrderInsertCallback(pInputExecOrder, pRspInfo)
	}
}

// 执行宣告操作错误回报
func (s *TDSpiN) OnErrRtnExecOrderAction(pExecOrderAction *thost.CThostFtdcExecOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnExecOrderActionCallback != nil {
		s.OnErrRtnExecOrderActionCallback(pExecOrderAction, pRspInfo)
	}
}

// 询价录入错误回报
func (s *TDSpiN) OnErrRtnForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnForQuoteInsertCallback != nil {
		s.OnErrRtnForQuoteInsertCallback(pInputForQuote, pRspInfo)
	}
}

// 报价通知
func (s *TDSpiN) OnRtnQuote(pQuote *thost.CThostFtdcQuoteField) {
	quoteData := map[string]interface{}{
		"InstrumentID": pQuote.InstrumentID.String(),      // 合约代码
		"投资者代码":        pQuote.InvestorID.String(),        // 投资者代码
		"报价引用":         pQuote.QuoteRef.String(),          // 报价引用
		"用户代码":         pQuote.UserID.String(),            // 用户代码
		"卖价格":          pQuote.AskPrice.String(),          // 卖价格
		"买价格":          pQuote.BidPrice.String(),          // 买价格
		"卖数量":          pQuote.AskVolume,                  // 卖数量
		"买数量":          pQuote.BidVolume,                  // 买数量
		"请求编号":         pQuote.RequestID,                  // 请求编号
		"业务单元":         pQuote.BusinessUnit.String(),      // 业务单元
		"卖开平标志":        pQuote.AskOffsetFlag.String(),     // 卖开平标志
		"买开平标志":        pQuote.BidOffsetFlag.String(),     // 买开平标志
		"卖投机套保标志":      pQuote.AskHedgeFlag.String(),      // 卖投机套保标志
		"买投机套保标志":      pQuote.BidHedgeFlag.String(),      // 买投机套保标志
		"本地报价编号":       pQuote.QuoteLocalID.String(),      // 本地报价编号
		"交易所代码":        pQuote.ExchangeID.String(),        // 交易所代码
		"会员代码":         pQuote.ParticipantID.String(),     // 会员代码
		"客户代码":         pQuote.ClientID.String(),          // 客户代码
		"交易所交易员代码":     pQuote.TraderID.String(),          // 交易所交易员代码
		"安装编号":         pQuote.InstallID,                  // 安装编号
		"报价提示序号":       pQuote.NotifySequence,             // 报价提示序号
		"报价提交状态":       pQuote.OrderSubmitStatus.String(), // 报价提交状态
		"交易日":          pQuote.TradingDay.String(),        // 交易日
		"结算编号":         pQuote.SettlementID,               // 结算编号
		"报价编号":         pQuote.QuoteSysID,                 // 报价编号
		"报单日期":         pQuote.InsertDate.String(),        // 报单日期
		"插入时间":         pQuote.InsertTime.String(),        // 插入时间
		"撤销时间":         pQuote.CancelTime.String(),        // 撤销时间
		"报价状态":         pQuote.QuoteStatus.String(),       // 报价状态
		"结算会员编号":       pQuote.ClearingPartID.String(),    // 结算会员编号
		"序号":           pQuote.SequenceNo,                 // 序号
		"卖方报单编号":       pQuote.AskOrderSysID,              // 卖方报单编号
		"买方报单编号":       pQuote.BidOrderSysID,              // 买方报单编号
		"前置编号":         pQuote.FrontID,                    // 前置编号
		"会话编号":         pQuote.SessionID,                  // 会话编号
		"用户端产品信息":      pQuote.UserProductInfo.String(),   // 用户端产品信息
		"状态信息":         pQuote.StatusMsg.GBString(),       // 状态信息
		"操作用户代码":       pQuote.ActiveUserID.String(),      // 操作用户代码
		"经纪公司报价编号":     pQuote.BrokerQuoteSeq,             // 经纪公司报价编号
		"衍生卖报单引用":      pQuote.AskOrderRef.String(),       // 衍生卖报单引用
		"衍生买报单引用":      pQuote.BidOrderRef.String(),       // 衍生买报单引用
		"应价编号":         pQuote.ForQuoteSysID,              // 应价编号
		"营业部编号":        pQuote.BranchID.String(),          // 营业部编号
		"投资单元代码":       pQuote.InvestUnitID.String(),      // 投资单元代码
		"资金账号":         pQuote.AccountID.String(),         // 资金账号
		"币种代码":         pQuote.CurrencyID.String(),        // 币种代码
		"Mac地址":        pQuote.MacAddress.String(),        // Mac地址
		"合约代码":         pQuote.InstrumentID.String(),      // 合约代码
		"合约在交易所的代码":    pQuote.ExchangeInstID.String(),    // 合约在交易所的代码
		"IP地址":         pQuote.IPAddress.String(),         // IP地址
		"被顶单编号":        pQuote.ReplaceSysID.String(),      // 被顶单编号
	}
	s.OnData(quoteData, true, "OnRtnQuote")
}

// 报价录入错误回报
func (s *TDSpiN) OnErrRtnQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if ErrorRspInfo(pRspInfo) || pInputQuote == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "OnErrRtnQuoteInsert")
		return
	}
	errData := map[string]interface{}{ // 构建错误信息映射
		"InstrumentID": pInputQuote.InstrumentID.String(),  // 合约代码
		"投资者代码":        pInputQuote.InvestorID.String(),    // 投资者代码
		"报价引用":         pInputQuote.QuoteRef.String(),      // 报价引用
		"用户代码":         pInputQuote.UserID.String(),        // 用户代码
		"卖价格":          pInputQuote.AskPrice.String(),      // 卖价格
		"买价格":          pInputQuote.BidPrice.String(),      // 买价格
		"卖数量":          pInputQuote.AskVolume,              // 卖数量
		"买数量":          pInputQuote.BidVolume,              // 买数量
		"请求编号":         pInputQuote.RequestID,              // 请求编号
		"业务单元":         pInputQuote.BusinessUnit.String(),  // 业务单元
		"卖开平标志":        pInputQuote.AskOffsetFlag.String(), // 卖开平标志
		"买开平标志":        pInputQuote.BidOffsetFlag.String(), // 买开平标志
		"卖投机套保标志":      pInputQuote.AskHedgeFlag.String(),  // 卖投机套保标志
		"买投机套保标志":      pInputQuote.BidHedgeFlag.String(),  // 买投机套保标志
		"衍生卖报单引用":      pInputQuote.AskOrderRef.String(),   // 衍生卖报单引用
		"衍生买报单引用":      pInputQuote.BidOrderRef.String(),   // 衍生买报单引用
		"应价编号":         pInputQuote.ForQuoteSysID,          // 应价编号
		"交易所代码":        pInputQuote.ExchangeID.String(),    // 交易所代码
		"投资单元代码":       pInputQuote.InvestUnitID.String(),  // 投资单元代码
		"交易编码":         pInputQuote.ClientID.String(),      // 交易编码
		"Mac地址":        pInputQuote.MacAddress.String(),    // Mac地址
		"IP地址":         pInputQuote.IPAddress.String(),     // IP地址
		"被顶单编号":        pInputQuote.ReplaceSysID.String(),  // 被顶单编号
		"错误代码":         pRspInfo.ErrorID,                   // 错误代码
		"错误信息":         pRspInfo.ErrorMsg.String(),         // 错误信息
	}
	s.OnData(errData, true, "OnErrRtnQuoteInsert") // 处理错误信息
}

// 报价操作错误回报
func (s *TDSpiN) OnErrRtnQuoteAction(pQuoteAction *thost.CThostFtdcQuoteActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnQuoteActionCallback != nil {
		s.OnErrRtnQuoteActionCallback(pQuoteAction, pRspInfo)
	}
}

// 询价通知
func (s *TDSpiN) OnRtnForQuoteRsp(pForQuoteRsp *thost.CThostFtdcForQuoteRspField) {
	if s.OnRtnForQuoteRspCallback != nil {
		s.OnRtnForQuoteRspCallback(pForQuoteRsp)
	}
}

// 保证金监控中心用户令牌
func (s *TDSpiN) OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *thost.CThostFtdcCFMMCTradingAccountTokenField) {
	if s.OnRtnCFMMCTradingAccountTokenCallback != nil {
		s.OnRtnCFMMCTradingAccountTokenCallback(pCFMMCTradingAccountToken)
	}
}

// 批量报单操作错误回报
func (s *TDSpiN) OnErrRtnBatchOrderAction(pBatchOrderAction *thost.CThostFtdcBatchOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnBatchOrderActionCallback != nil {
		s.OnErrRtnBatchOrderActionCallback(pBatchOrderAction, pRspInfo)
	}
}

// 期权自对冲通知
func (s *TDSpiN) OnRtnOptionSelfClose(pOptionSelfClose *thost.CThostFtdcOptionSelfCloseField) {
	if s.OnRtnOptionSelfCloseCallback != nil {
		s.OnRtnOptionSelfCloseCallback(pOptionSelfClose)
	}
}

// 期权自对冲录入错误回报
func (s *TDSpiN) OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnOptionSelfCloseInsertCallback != nil {
		s.OnErrRtnOptionSelfCloseInsertCallback(pInputOptionSelfClose, pRspInfo)
	}
}

// 期权自对冲操作错误回报
func (s *TDSpiN) OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *thost.CThostFtdcOptionSelfCloseActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnOptionSelfCloseActionCallback != nil {
		s.OnErrRtnOptionSelfCloseActionCallback(pOptionSelfCloseAction, pRspInfo)
	}
}

// 申请组合通知
func (s *TDSpiN) OnRtnCombAction(pCombAction *thost.CThostFtdcCombActionField) {
	if s.OnRtnCombActionCallback != nil {
		s.OnRtnCombActionCallback(pCombAction)
	}
}

// 申请组合录入错误回报
func (s *TDSpiN) OnErrRtnCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnCombActionInsertCallback != nil {
		s.OnErrRtnCombActionInsertCallback(pInputCombAction, pRspInfo)
	}
}

// 银行发起银行资金转期货通知
func (s *TDSpiN) OnRtnFromBankToFutureByBank(pRspTransfer *thost.CThostFtdcRspTransferField) {
	if s.OnRtnFromBankToFutureByBankCallback != nil {
		s.OnRtnFromBankToFutureByBankCallback(pRspTransfer)
	}
}

// 银行发起期货资金转银行通知
func (s *TDSpiN) OnRtnFromFutureToBankByBank(pRspTransfer *thost.CThostFtdcRspTransferField) {
	if s.OnRtnFromFutureToBankByBankCallback != nil {
		s.OnRtnFromFutureToBankByBankCallback(pRspTransfer)
	}
}

// 银行发起冲正银行转期货通知
func (s *TDSpiN) OnRtnRepealFromBankToFutureByBank(pRspRepeal *thost.CThostFtdcRspRepealField) {
	if s.OnRtnRepealFromBankToFutureByBankCallback != nil {
		s.OnRtnRepealFromBankToFutureByBankCallback(pRspRepeal)
	}
}

// 银行发起冲正期货转银行通知
func (s *TDSpiN) OnRtnRepealFromFutureToBankByBank(pRspRepeal *thost.CThostFtdcRspRepealField) {
	if s.OnRtnRepealFromFutureToBankByBankCallback != nil {
		s.OnRtnRepealFromFutureToBankByBankCallback(pRspRepeal)
	}
}

// 期货发起银行资金转期货通知
func (s *TDSpiN) OnRtnFromBankToFutureByFuture(pRspTransfer *thost.CThostFtdcRspTransferField) {
	if s.OnRtnFromBankToFutureByFutureCallback != nil {
		s.OnRtnFromBankToFutureByFutureCallback(pRspTransfer)
	}
}

// 期货发起期货资金转银行通知
func (s *TDSpiN) OnRtnFromFutureToBankByFuture(pRspTransfer *thost.CThostFtdcRspTransferField) {
	if s.OnRtnFromFutureToBankByFutureCallback != nil {
		s.OnRtnFromFutureToBankByFutureCallback(pRspTransfer)
	}
}

// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (s *TDSpiN) OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *thost.CThostFtdcRspRepealField) {
	if s.OnRtnRepealFromBankToFutureByFutureManualCallback != nil {
		s.OnRtnRepealFromBankToFutureByFutureManualCallback(pRspRepeal)
	}
}

// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (s *TDSpiN) OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *thost.CThostFtdcRspRepealField) {
	if s.OnRtnRepealFromFutureToBankByFutureManualCallback != nil {
		s.OnRtnRepealFromFutureToBankByFutureManualCallback(pRspRepeal)
	}
}

// 期货发起查询银行余额通知
func (s *TDSpiN) OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *thost.CThostFtdcNotifyQueryAccountField) {
	if s.OnRtnQueryBankBalanceByFutureCallback != nil {
		s.OnRtnQueryBankBalanceByFutureCallback(pNotifyQueryAccount)
	}
}

// 期货发起银行资金转期货错误回报
func (s *TDSpiN) OnErrRtnBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnBankToFutureByFutureCallback != nil {
		s.OnErrRtnBankToFutureByFutureCallback(pReqTransfer, pRspInfo)
	}
}

// 期货发起期货资金转银行错误回报
func (s *TDSpiN) OnErrRtnFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnFutureToBankByFutureCallback != nil {
		s.OnErrRtnFutureToBankByFutureCallback(pReqTransfer, pRspInfo)
	}
}

// 系统运行时期货端手工发起冲正银行转期货错误回报
func (s *TDSpiN) OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal *thost.CThostFtdcReqRepealField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnRepealBankToFutureByFutureManualCallback != nil {
		s.OnErrRtnRepealBankToFutureByFutureManualCallback(pReqRepeal, pRspInfo)
	}
}

// 系统运行时期货端手工发起冲正期货转银行错误回报
func (s *TDSpiN) OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal *thost.CThostFtdcReqRepealField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnRepealFutureToBankByFutureManualCallback != nil {
		s.OnErrRtnRepealFutureToBankByFutureManualCallback(pReqRepeal, pRspInfo)
	}
}

// 期货发起查询银行余额错误回报
func (s *TDSpiN) OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, pRspInfo *thost.CThostFtdcRspInfoField) {
	if s.OnErrRtnQueryBankBalanceByFutureCallback != nil {
		s.OnErrRtnQueryBankBalanceByFutureCallback(pReqQueryAccount, pRspInfo)
	}
}

// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (s *TDSpiN) OnRtnRepealFromBankToFutureByFuture(pRspRepeal *thost.CThostFtdcRspRepealField) {
	if s.OnRtnRepealFromBankToFutureByFutureCallback != nil {
		s.OnRtnRepealFromBankToFutureByFutureCallback(pRspRepeal)
	}
}

// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (s *TDSpiN) OnRtnRepealFromFutureToBankByFuture(pRspRepeal *thost.CThostFtdcRspRepealField) {
	if s.OnRtnRepealFromFutureToBankByFutureCallback != nil {
		s.OnRtnRepealFromFutureToBankByFutureCallback(pRspRepeal)
	}
}

// 银行发起银期开户通知
func (s *TDSpiN) OnRtnOpenAccountByBank(pOpenAccount *thost.CThostFtdcOpenAccountField) {
	if s.OnRtnOpenAccountByBankCallback != nil {
		s.OnRtnOpenAccountByBankCallback(pOpenAccount)
	}
}

// 银行发起银期销户通知
func (s *TDSpiN) OnRtnCancelAccountByBank(pCancelAccount *thost.CThostFtdcCancelAccountField) {
	if s.OnRtnCancelAccountByBankCallback != nil {
		s.OnRtnCancelAccountByBankCallback(pCancelAccount)
	}
}

// 银行发起变更银行账号通知
func (s *TDSpiN) OnRtnChangeAccountByBank(pChangeAccount *thost.CThostFtdcChangeAccountField) {
	if s.OnRtnChangeAccountByBankCallback != nil {
		s.OnRtnChangeAccountByBankCallback(pChangeAccount)
	}
}
