package CTPBase

import (
	"github.com/li-wancai/GoScripts/Formulae"

	"github.com/pseudocodes/go2ctp/thost"
)

// 客户端认证响应
func (s *TDSpiN) OnRspAuthenticate(pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if pRspInfo.ErrorID == 0 {
		log.Infof("【%v】【客户端认证】成功", pRspAuthenticateField.UserID)
		s.OnRspAuthenticateCallback(pRspAuthenticateField, pRspInfo, nRequestID, bIsLast)
	} else {
		log.Errorf(
			"【%v】认证失败，错误码: %v 错误信息: %v",
			pRspAuthenticateField.UserID, pRspInfo.ErrorID, pRspInfo.ErrorMsg.GBString())
	}
}

// 客户端认证响应回调
func (s *TDSpiN) OnRspAuthenticateCallback(pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	s.RequestID++
	req := &thost.CThostFtdcReqUserLoginField{}
	copy(req.UserID[:], s.Cfg.UserID)
	copy(req.Password[:], s.Cfg.Password)
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.UserProductInfo[:], s.Cfg.ProductInfo)
	s.Api.ReqUserLogin(req, s.RequestID)
}

// 登录请求响应
func (s *TDSpiN) OnRspUserLogin(pRspUserLogin *thost.CThostFtdcRspUserLoginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if pRspInfo.ErrorID == 0 { // 检查登录响应是否成功
		log.Infof("【交易登录】交易日: %v 成功", pRspUserLogin.TradingDay) // 登录成功，输出相关信息
		s.OnRspUserLoginCallback(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
	} else {
		log.Errorf("【交易登录】失败 错误代码：%d, 错误信息：%v", pRspInfo.ErrorID, pRspInfo.ErrorMsg.GBString()) // 登录失败，记录错误信息
	}
}

func (s *TDSpiN) OnRspUserLoginCallback(
	pRspUserLogin *thost.CThostFtdcRspUserLoginField,
	pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	////////////////////////////////////////////////////////////////////////////////////////////////////
	if ErrorRspInfo(pRspInfo) || pRspUserLogin == nil {
		s.OnData(map[string]interface{}{}, true, "登录成功")
		return
	}
	UserLoginInfo := map[string]interface{}{
		"前置编号":   pRspUserLogin.FrontID,
		"会话编号":   pRspUserLogin.SessionID,
		"用户代码":   pRspUserLogin.UserID,
		"交易日":    pRspUserLogin.TradingDay,
		"登录成功时间": pRspUserLogin.LoginTime,
		"经纪公司代码": pRspUserLogin.BrokerID,
		"最大报单引用": pRspUserLogin.MaxOrderRef,
		"上期所时间":  pRspUserLogin.SHFETime,
		"大商所时间":  pRspUserLogin.DCETime,
		"郑商所时间":  pRspUserLogin.CZCETime,
		"中金所时间":  pRspUserLogin.FFEXTime,
		"能源中心时间": pRspUserLogin.INETime,
		"广期所时间":  pRspUserLogin.GFEXTime,
		"后台版本信息": pRspUserLogin.SysVersion,
		"交易系统名称": pRspUserLogin.SystemName}
	log.Infof("【交易登录信息】: \n%+v", Formulae.DictToStrJson(UserLoginInfo))
	////////////////////////////////////////////////////////////////////////////////////////////////////
	s.OnData(UserLoginInfo, bIsLast, "登录成功") //丢给OnData处理
	////////////////////////////////////////////////////////////////////////////////////////////////////
}

// 发送投资者结算单确认响应
func (s *TDSpiN) OnRspSettlementInfoConfirm(
	pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField,
	pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pSettlementInfoConfirm == nil {
		s.OnData(map[string]interface{}{}, true, "结算信息")
		return
	}
	SettlementInfo := map[string]interface{}{
		"经纪公司代码": pSettlementInfoConfirm.BrokerID,
		"投资者代码":  pSettlementInfoConfirm.InvestorID,
		"确认日期":   pSettlementInfoConfirm.ConfirmDate,
		"确认时间":   pSettlementInfoConfirm.ConfirmTime,
		"结算编号":   pSettlementInfoConfirm.SettlementID,
		"投资者帐号":  pSettlementInfoConfirm.AccountID,
		"币种代码":   pSettlementInfoConfirm.CurrencyID,
	}
	s.OnData(SettlementInfo, bIsLast, "结算信息")
}

// 登出请求响应
func (s *TDSpiN) OnRspUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserLogoutCallback != nil {
		s.OnRspUserLogoutCallback(pUserLogout, pRspInfo, nRequestID, bIsLast)
	}
}

// 用户口令更新请求响应
func (s *TDSpiN) OnRspUserPasswordUpdate(pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserPasswordUpdateCallback != nil {
		s.OnRspUserPasswordUpdateCallback(pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast)
	}
}

// 资金账户口令更新请求响应
func (s *TDSpiN) OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspTradingAccountPasswordUpdateCallback != nil {
		s.OnRspTradingAccountPasswordUpdateCallback(pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast)
	}
}

// 查询用户当前支持的认证模式的回复
func (s *TDSpiN) OnRspUserAuthMethod(pRspUserAuthMethod *thost.CThostFtdcRspUserAuthMethodField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserAuthMethodCallback != nil {
		s.OnRspUserAuthMethodCallback(pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast)
	}
}

// 获取图形验证码请求的回复
func (s *TDSpiN) OnRspGenUserCaptcha(pRspGenUserCaptcha *thost.CThostFtdcRspGenUserCaptchaField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspGenUserCaptchaCallback != nil {
		s.OnRspGenUserCaptchaCallback(pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast)
	}
}

// 获取短信验证码请求的回复
func (s *TDSpiN) OnRspGenUserText(pRspGenUserText *thost.CThostFtdcRspGenUserTextField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspGenUserTextCallback != nil {
		s.OnRspGenUserTextCallback(pRspGenUserText, pRspInfo, nRequestID, bIsLast)
	}
}

// 报单录入请求响应
func (s *TDSpiN) OnRspOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if pInputOrder == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "OnRspOrderInsert")
		return
	}
	data := map[string]interface{}{
		"经纪公司代码":       pInputOrder.BrokerID.String(),
		"投资者代码":        pInputOrder.InvestorID.String(),
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
		"交易所代码":        pInputOrder.ExchangeID.String(),
		"投资单元代码":       pInputOrder.InvestUnitID.String(),
		"资金账号":         pInputOrder.AccountID.String(),
		"币种代码":         pInputOrder.CurrencyID.String(),
		"交易编码":         pInputOrder.ClientID.String(),
		"Mac地址":        pInputOrder.MacAddress.String(),
		"InstrumentID": pInputOrder.InstrumentID.String(),
		"IP地址":         pInputOrder.IPAddress.String(),
	}
	s.OnData(data, bIsLast, "OnRspOrderInsert")
	ErrorRspInfo(pRspInfo)
}

// 预埋单录入请求响应
func (s *TDSpiN) OnRspParkedOrderInsert(pParkedOrder *thost.CThostFtdcParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspParkedOrderInsertCallback != nil {
		s.OnRspParkedOrderInsertCallback(pParkedOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// 预埋撤单录入请求响应
func (s *TDSpiN) OnRspParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspParkedOrderActionCallback != nil {
		s.OnRspParkedOrderActionCallback(pParkedOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 报单操作请求响应
func (s *TDSpiN) OnRspOrderAction(pInputOrderAction *thost.CThostFtdcInputOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if pInputOrderAction == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "OnRspOrderAction")
		return
	}
	data := map[string]interface{}{
		"经纪公司代码":       pInputOrderAction.BrokerID.String(),
		"投资者代码":        pInputOrderAction.InvestorID.String(),
		"报单操作引用":       pInputOrderAction.OrderActionRef,
		"报单引用":         pInputOrderAction.OrderRef.String(),
		"请求编号":         pInputOrderAction.RequestID,
		"前置编号":         pInputOrderAction.FrontID,
		"会话编号":         pInputOrderAction.SessionID,
		"交易所代码":        pInputOrderAction.ExchangeID.String(),
		"报单编号":         pInputOrderAction.OrderSysID.String(),
		"操作标志":         pInputOrderAction.ActionFlag.String(),
		"价格":           pInputOrderAction.LimitPrice.String(),
		"数量变化":         pInputOrderAction.VolumeChange,
		"用户代码":         pInputOrderAction.UserID.String(),
		"投资单元代码":       pInputOrderAction.InvestUnitID.String(),
		"Mac地址":        pInputOrderAction.MacAddress.String(),
		"InstrumentID": pInputOrderAction.InstrumentID.String(),
		"IP地址":         pInputOrderAction.IPAddress.String(),
	}
	s.OnData(data, bIsLast, "OnRspOrderAction")
	ErrorRspInfo(pRspInfo)
}

// 查询最大报单数量响应
func (s *TDSpiN) OnRspQryMaxOrderVolume(pQryMaxOrderVolume *thost.CThostFtdcQryMaxOrderVolumeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryMaxOrderVolumeCallback != nil {
		s.OnRspQryMaxOrderVolumeCallback(pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast)
	}
}

// // 请求查询资金账户响应
func (s *TDSpiN) OnRspQryTradingAccount(pTradingAccount *thost.CThostFtdcTradingAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pTradingAccount == nil {
		s.OnData(map[string]interface{}{"投资者帐号": s.Cfg.UserID}, true, "资金账户")
		return
	}
	TradingAccount := map[string]interface{}{
		"业务类型":       pTradingAccount.BizType,
		"延时换汇冻结金额":   pTradingAccount.FrozenSwap,
		"剩余换汇额度":     pTradingAccount.RemainSwap,
		"经纪公司代码":     pTradingAccount.BrokerID,
		"投资者帐号":      pTradingAccount.AccountID.GBString(),
		"上次质押金额":     pTradingAccount.PreMortgage,
		"上次信用额度":     pTradingAccount.PreCredit,
		"上次存款额":      pTradingAccount.PreDeposit,
		"上次结算准备金":    pTradingAccount.PreBalance,
		"上次占用的保证金":   pTradingAccount.PreMargin,
		"利息基数":       pTradingAccount.InterestBase,
		"利息收入":       pTradingAccount.Interest,
		"入金金额":       pTradingAccount.Deposit,
		"出金金额":       pTradingAccount.Withdraw,
		"冻结的保证金":     pTradingAccount.FrozenMargin,
		"冻结的资金":      pTradingAccount.FrozenCash,
		"冻结的手续费":     pTradingAccount.FrozenCommission,
		"当前保证金总额":    pTradingAccount.CurrMargin,
		"资金差额":       pTradingAccount.CashIn,
		"手续费":        pTradingAccount.Commission,
		"平仓盈亏":       pTradingAccount.CloseProfit,
		"持仓盈亏":       pTradingAccount.PositionProfit,
		"期货结算准备金":    pTradingAccount.Balance,
		"信用额度":       pTradingAccount.Credit,
		"质押金额":       pTradingAccount.Mortgage,
		"基本准备金":      pTradingAccount.Reserve,
		"可用资金":       pTradingAccount.Available,
		"交易日":        pTradingAccount.TradingDay,
		"结算编号":       pTradingAccount.SettlementID,
		"可取资金":       pTradingAccount.WithdrawQuota,
		"交易所保证金":     pTradingAccount.ExchangeMargin,
		"投资者交割保证金":   pTradingAccount.DeliveryMargin,
		"币种代码":       pTradingAccount.CurrencyID,
		"货币质入金额":     pTradingAccount.FundMortgageIn,
		"货币质出金额":     pTradingAccount.FundMortgageOut,
		"可质押货币金额":    pTradingAccount.MortgageableFund,
		"保底期货结算准备金":  pTradingAccount.ReserveBalance,
		"上次货币质入金额":   pTradingAccount.PreFundMortgageIn,
		"特殊产品占用保证金":  pTradingAccount.SpecProductMargin,
		"上次货币质出金额":   pTradingAccount.PreFundMortgageOut,
		"货币质押余额":     pTradingAccount.FundMortgageAvailable,
		"特殊产品手续费":    pTradingAccount.SpecProductCommission,
		"交易所交割保证金":   pTradingAccount.ExchangeDeliveryMargin,
		"特殊产品平仓盈亏":   pTradingAccount.SpecProductCloseProfit,
		"特殊产品冻结保证金":  pTradingAccount.SpecProductFrozenMargin,
		"特殊产品持仓盈亏":   pTradingAccount.SpecProductPositionProfit,
		"特殊产品交易所保证金": pTradingAccount.SpecProductExchangeMargin,
		"特殊产品冻结手续费":  pTradingAccount.SpecProductFrozenCommission,
		"根据持仓盈亏算法计算的特殊产品持仓盈亏": pTradingAccount.SpecProductPositionProfitByAlg,
	}
	s.OnData(TradingAccount, bIsLast, "资金账户")
}

// 请求查询合约手续费率响应
func (s *TDSpiN) OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *thost.CThostFtdcInstrumentCommissionRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pInstrumentCommissionRate == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "手续费率")
		return
	}
	CommissionRate := map[string]interface{}{
		"投资者代码":        pInstrumentCommissionRate.InvestorID.String(),
		"开仓手续费率":       pInstrumentCommissionRate.OpenRatioByMoney.String(),
		"开仓手续费":        pInstrumentCommissionRate.OpenRatioByVolume.String(),
		"平仓手续费率":       pInstrumentCommissionRate.CloseRatioByMoney.String(),
		"平仓手续费":        pInstrumentCommissionRate.CloseRatioByVolume.String(),
		"平今手续费率":       pInstrumentCommissionRate.CloseTodayRatioByMoney.String(),
		"平今手续费":        pInstrumentCommissionRate.CloseTodayRatioByVolume.String(),
		"交易所代码":        pInstrumentCommissionRate.ExchangeID.String(),
		"投资单元代码":       pInstrumentCommissionRate.InvestUnitID.String(),
		"InstrumentID": pInstrumentCommissionRate.InstrumentID.String(),
		"投资者范围":        pInstrumentCommissionRate.InvestorRange.String(),
		"经纪公司代码":       pInstrumentCommissionRate.BrokerID.String(),
		"业务类型":         pInstrumentCommissionRate.BizType.String(),
	}
	s.OnData(CommissionRate, bIsLast, "手续费率")
}

// 删除预埋单响应
func (s *TDSpiN) OnRspRemoveParkedOrder(pRemoveParkedOrder *thost.CThostFtdcRemoveParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspRemoveParkedOrderCallback != nil {
		s.OnRspRemoveParkedOrderCallback(pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// 删除预埋撤单响应
func (s *TDSpiN) OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction *thost.CThostFtdcRemoveParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspRemoveParkedOrderActionCallback != nil {
		s.OnRspRemoveParkedOrderActionCallback(pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 执行宣告录入请求响应
func (s *TDSpiN) OnRspExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspExecOrderInsertCallback != nil {
		s.OnRspExecOrderInsertCallback(pInputExecOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// 执行宣告操作请求响应
func (s *TDSpiN) OnRspExecOrderAction(pInputExecOrderAction *thost.CThostFtdcInputExecOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspExecOrderActionCallback != nil {
		s.OnRspExecOrderActionCallback(pInputExecOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 询价录入请求响应
func (s *TDSpiN) OnRspForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspForQuoteInsertCallback != nil {
		s.OnRspForQuoteInsertCallback(pInputForQuote, pRspInfo, nRequestID, bIsLast)
	}
}

// 报价录入请求响应
func (s *TDSpiN) OnRspQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pInputQuote == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "OnRspQuoteInsert")
		return
	}
	data := map[string]interface{}{
		"InstrumentID": pInputQuote.InstrumentID.String(),
		"投资者代码":        pInputQuote.InvestorID.String(),
		"交易所代码":        pInputQuote.ExchangeID.String(),
		"投资单元代码":       pInputQuote.InvestUnitID.String(),
		"经纪公司代码":       pInputQuote.BrokerID.String(),
		"报价引用":         pInputQuote.QuoteRef.String(),
		"用户代码":         pInputQuote.UserID.String(),
		"卖价格":          pInputQuote.AskPrice.String(),
		"买价格":          pInputQuote.BidPrice.String(),
		"卖数量":          pInputQuote.AskVolume,
		"买数量":          pInputQuote.BidVolume,
		"请求编号":         pInputQuote.RequestID,
		"业务单元":         pInputQuote.BusinessUnit.String(),
		"卖开平标志":        pInputQuote.AskOffsetFlag.String(),
		"买开平标志":        pInputQuote.BidOffsetFlag.String(),
		"卖投机套保标志":      pInputQuote.AskHedgeFlag.String(),
		"买投机套保标志":      pInputQuote.BidHedgeFlag.String(),
		"衍生卖报单引用":      pInputQuote.AskOrderRef.String(),
		"衍生买报单引用":      pInputQuote.BidOrderRef.String(),
		"应价编号":         pInputQuote.ForQuoteSysID.String(),
		"交易编码":         pInputQuote.ClientID.String(),
		"MacAddress":   pInputQuote.MacAddress.String(),
		"IP地址":         pInputQuote.IPAddress.String(),
		"被顶单编号":        pInputQuote.ReplaceSysID.String(),
	}
	s.OnData(data, bIsLast, "OnRspQuoteInsert")
}

// 报价操作请求响应
func (s *TDSpiN) OnRspQuoteAction(pInputQuoteAction *thost.CThostFtdcInputQuoteActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQuoteActionCallback != nil {
		s.OnRspQuoteActionCallback(pInputQuoteAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 批量报单操作请求响应
func (s *TDSpiN) OnRspBatchOrderAction(pInputBatchOrderAction *thost.CThostFtdcInputBatchOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspBatchOrderActionCallback != nil {
		s.OnRspBatchOrderActionCallback(pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 期权自对冲录入请求响应
func (s *TDSpiN) OnRspOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOptionSelfCloseInsertCallback != nil {
		s.OnRspOptionSelfCloseInsertCallback(pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast)
	}
}

// 期权自对冲操作请求响应
func (s *TDSpiN) OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction *thost.CThostFtdcInputOptionSelfCloseActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspOptionSelfCloseActionCallback != nil {
		s.OnRspOptionSelfCloseActionCallback(pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 申请组合录入请求响应
func (s *TDSpiN) OnRspCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspCombActionInsertCallback != nil {
		s.OnRspCombActionInsertCallback(pInputCombAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询报单响应
func (s *TDSpiN) OnRspQryOrder(pOrder *thost.CThostFtdcOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOrderCallback != nil {
		s.OnRspQryOrderCallback(pOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询成交响应
func (s *TDSpiN) OnRspQryTrade(pTrade *thost.CThostFtdcTradeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTradeCallback != nil {
		s.OnRspQryTradeCallback(pTrade, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询投资者持仓响应
func (s *TDSpiN) OnRspQryInvestorPosition(pInvestorPosition *thost.CThostFtdcInvestorPositionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pInvestorPosition == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "OnRspQryInvestorPosition")
		return
	}
	data := map[string]interface{}{
		"InstrumentID": pInvestorPosition.InstrumentID.String(),
		"经纪公司代码":       pInvestorPosition.BrokerID.String(),
		"投资者代码":        pInvestorPosition.InvestorID.String(),
		"持仓多空方向":       pInvestorPosition.PosiDirection.String(),
		"投机套保标志":       pInvestorPosition.HedgeFlag.String(),
		"持仓日期":         pInvestorPosition.PositionDate.String(),
		"上日持仓":         pInvestorPosition.YdPosition,
		"今日持仓":         pInvestorPosition.Position,
		"多头冻结":         pInvestorPosition.LongFrozen,
		"空头冻结":         pInvestorPosition.ShortFrozen,
		"开仓冻结金额":       pInvestorPosition.LongFrozenAmount,
		"平仓冻结金额":       pInvestorPosition.ShortFrozenAmount,
		"开仓量":          pInvestorPosition.OpenVolume,
		"平仓量":          pInvestorPosition.CloseVolume,
		"开仓金额":         pInvestorPosition.OpenAmount,
		"平仓金额":         pInvestorPosition.CloseAmount,
		"持仓成本":         pInvestorPosition.PositionCost,
		"上次占用的保证金":     pInvestorPosition.PreMargin,
		"占用的保证金":       pInvestorPosition.UseMargin,
		"冻结的保证金":       pInvestorPosition.FrozenMargin,
		"冻结的资金":        pInvestorPosition.FrozenCash,
		"冻结的手续费":       pInvestorPosition.FrozenCommission,
		"资金差额":         pInvestorPosition.CashIn,
		"手续费":          pInvestorPosition.Commission,
		"平仓盈亏":         pInvestorPosition.CloseProfit,
		"持仓盈亏":         pInvestorPosition.PositionProfit,
		"上次结算价":        pInvestorPosition.PreSettlementPrice,
		"本次结算价":        pInvestorPosition.SettlementPrice,
		"交易日":          pInvestorPosition.TradingDay,
		"结算编号":         pInvestorPosition.SettlementID,
		"开仓成本":         pInvestorPosition.OpenCost,
		"交易所保证金":       pInvestorPosition.ExchangeMargin,
		"组合成交形成的持仓":    pInvestorPosition.CombPosition,
		"组合多头冻结":       pInvestorPosition.CombLongFrozen,
		"组合空头冻结":       pInvestorPosition.CombShortFrozen,
		"逐日盯市平仓盈亏":     pInvestorPosition.CloseProfitByDate,
		"逐笔对冲平仓盈亏":     pInvestorPosition.CloseProfitByTrade,
		"保证金率":         pInvestorPosition.MarginRateByMoney,
		"保证金率(按手数)":    pInvestorPosition.MarginRateByVolume,
		"执行冻结":         pInvestorPosition.StrikeFrozen,
		"执行冻结金额":       pInvestorPosition.StrikeFrozenAmount,
		"放弃执行冻结":       pInvestorPosition.AbandonFrozen,
		"交易所代码":        pInvestorPosition.ExchangeID.String(),
		"执行冻结的昨仓":      pInvestorPosition.YdStrikeFrozen,
		"投资单元代码":       pInvestorPosition.InvestUnitID.String(),
		"持仓成本差值":       pInvestorPosition.PositionCostOffset,
		"tas持仓手数":      pInvestorPosition.TasPosition,
		"tas持仓成本":      pInvestorPosition.TasPositionCost,
	}
	s.OnData(data, bIsLast, "OnRspQryInvestorPosition")
}

// 请求查询投资者响应
func (s *TDSpiN) OnRspQryInvestor(pInvestor *thost.CThostFtdcInvestorField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryInvestorCallback != nil {
		s.OnRspQryInvestorCallback(pInvestor, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询交易编码响应
func (s *TDSpiN) OnRspQryTradingCode(pTradingCode *thost.CThostFtdcTradingCodeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTradingCodeCallback != nil {
		s.OnRspQryTradingCodeCallback(pTradingCode, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询合约保证金率响应
func (s *TDSpiN) OnRspQryInstrumentMarginRate(pInstrumentMarginRate *thost.CThostFtdcInstrumentMarginRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pInstrumentMarginRate == nil {
		s.OnData(map[string]interface{}{"InstrumentID": ""}, true, "保证金率")
		return
	}
	MarginRate := map[string]interface{}{
		"投资者代码":        pInstrumentMarginRate.InvestorID.String(),
		"多头保证金率":       pInstrumentMarginRate.LongMarginRatioByMoney.String(),
		"多头保证金费":       pInstrumentMarginRate.LongMarginRatioByVolume.String(),
		"空头保证金率":       pInstrumentMarginRate.ShortMarginRatioByMoney.String(),
		"空头保证金费":       pInstrumentMarginRate.ShortMarginRatioByVolume.String(),
		"交易所代码":        pInstrumentMarginRate.ExchangeID.String(),
		"投资单元代码":       pInstrumentMarginRate.InvestUnitID.String(),
		"InstrumentID": pInstrumentMarginRate.InstrumentID.String(),
		"经纪公司代码":       pInstrumentMarginRate.BrokerID.String(),
		"投机套保标志":       pInstrumentMarginRate.HedgeFlag.String(),
		"是否相对交易所收取":    pInstrumentMarginRate.IsRelative,
		"投资者范围":        pInstrumentMarginRate.InvestorRange.String()}
	s.OnData(MarginRate, bIsLast, "保证金率")
}

// 请求查询交易所响应
func (s *TDSpiN) OnRspQryExchange(pExchange *thost.CThostFtdcExchangeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExchangeCallback != nil {
		s.OnRspQryExchangeCallback(pExchange, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询产品响应
func (s *TDSpiN) OnRspQryProduct(pProduct *thost.CThostFtdcProductField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryProductCallback != nil {
		s.OnRspQryProductCallback(pProduct, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询合约响应
func (s *TDSpiN) OnRspQryInstrument(pInstrument *thost.CThostFtdcInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pInstrument == nil {
		s.OnData(map[string]interface{}{"产品类型": "", "InstrumentID": ""}, true, "合约信息")
		return
	}
	InstrumentInfo := map[string]interface{}{
		"当前是否交易": pInstrument.IsTrading,
		"最小变动价位": pInstrument.PriceTick,
		"交易所代码":  pInstrument.ExchangeID,
		"交割年份":   pInstrument.DeliveryYear,
		"交割月":    pInstrument.DeliveryMonth,
		"合约数量乘数": pInstrument.VolumeMultiple,

		"限价单最大下单量": pInstrument.MaxLimitOrderVolume,
		"限价单最小下单量": pInstrument.MinLimitOrderVolume,
		"市价单最大下单量": pInstrument.MaxMarketOrderVolume,
		"市价单最小下单量": pInstrument.MinMarketOrderVolume,

		"产品代码":         pInstrument.ProductID.String(),
		"上市日":          pInstrument.OpenDate.GBString(),
		"执行价":          pInstrument.StrikePrice.String(),
		"期权类型":         pInstrument.OptionsType.String(),
		"到期日":          pInstrument.ExpireDate.GBString(),
		"创建日":          pInstrument.CreateDate.GBString(),
		"产品类型":         pInstrument.ProductClass.String(),
		"InstrumentID": pInstrument.InstrumentID.String(),
		"持仓类型":         pInstrument.PositionType.String(),

		"组合类型":          pInstrument.CombinationType.String(),
		"结束交割日":         pInstrument.EndDelivDate.GBString(),
		"合约名称":          pInstrument.InstrumentName.GBString(),
		"合约生命周期状态":      pInstrument.InstLifePhase.String(),
		"开始交割日":         pInstrument.StartDelivDate.GBString(),
		"多头保证金率":        pInstrument.LongMarginRatio.String(),
		"持仓日期类型":        pInstrument.PositionDateType.String(),
		"空头保证金率":        pInstrument.ShortMarginRatio.String(),
		"合约在交易所的代码":     pInstrument.ExchangeInstID.String(),
		"基础商品代码":        pInstrument.UnderlyingInstrID.String(),
		"合约基础商品乘数":      pInstrument.UnderlyingMultiple.String(),
		"是否使用大额单边保证金算法": pInstrument.MaxMarginSideAlgorithm.String(),
	}
	s.OnData(InstrumentInfo, bIsLast, "合约信息")
}

// 请求查询行情响应
func (s *TDSpiN) OnRspQryDepthMarketData(pDepthMarketData *thost.CThostFtdcDepthMarketDataField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryDepthMarketDataCallback != nil {
		s.OnRspQryDepthMarketDataCallback(pDepthMarketData, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询投资者结算结果响应
func (s *TDSpiN) OnRspQrySettlementInfo(pSettlementInfo *thost.CThostFtdcSettlementInfoField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySettlementInfoCallback != nil {
		s.OnRspQrySettlementInfoCallback(pSettlementInfo, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询转帐银行响应
func (s *TDSpiN) OnRspQryTransferBank(pTransferBank *thost.CThostFtdcTransferBankField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTransferBankCallback != nil {
		s.OnRspQryTransferBankCallback(pTransferBank, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询投资者持仓明细响应
func (s *TDSpiN) OnRspQryInvestorPositionDetail(pInvestorPositionDetail *thost.CThostFtdcInvestorPositionDetailField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryInvestorPositionDetailCallback != nil {
		s.OnRspQryInvestorPositionDetailCallback(pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询客户通知响应
func (s *TDSpiN) OnRspQryNotice(pNotice *thost.CThostFtdcNoticeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryNoticeCallback != nil {
		s.OnRspQryNoticeCallback(pNotice, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询结算信息确认响应
func (s *TDSpiN) OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySettlementInfoConfirmCallback != nil {
		s.OnRspQrySettlementInfoConfirmCallback(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询投资者持仓明细响应
func (s *TDSpiN) OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *thost.CThostFtdcInvestorPositionCombineDetailField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryInvestorPositionCombineDetailCallback != nil {
		s.OnRspQryInvestorPositionCombineDetailCallback(pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast)
	}
}

// 查询保证金监管系统经纪公司资金账户密钥响应
func (s *TDSpiN) OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *thost.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCFMMCTradingAccountKeyCallback != nil {
		s.OnRspQryCFMMCTradingAccountKeyCallback(pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询仓单折抵信息响应
func (s *TDSpiN) OnRspQryEWarrantOffset(pEWarrantOffset *thost.CThostFtdcEWarrantOffsetField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryEWarrantOffsetCallback != nil {
		s.OnRspQryEWarrantOffsetCallback(pEWarrantOffset, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询投资者品种/跨品种保证金响应
func (s *TDSpiN) OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *thost.CThostFtdcInvestorProductGroupMarginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryInvestorProductGroupMarginCallback != nil {
		s.OnRspQryInvestorProductGroupMarginCallback(pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询交易所保证金率响应
func (s *TDSpiN) OnRspQryExchangeMarginRate(pExchangeMarginRate *thost.CThostFtdcExchangeMarginRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExchangeMarginRateCallback != nil {
		s.OnRspQryExchangeMarginRateCallback(pExchangeMarginRate, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询交易所调整保证金率响应
func (s *TDSpiN) OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *thost.CThostFtdcExchangeMarginRateAdjustField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExchangeMarginRateAdjustCallback != nil {
		s.OnRspQryExchangeMarginRateAdjustCallback(pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询汇率响应
func (s *TDSpiN) OnRspQryExchangeRate(pExchangeRate *thost.CThostFtdcExchangeRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExchangeRateCallback != nil {
		s.OnRspQryExchangeRateCallback(pExchangeRate, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询二级代理操作员银期权限响应
func (s *TDSpiN) OnRspQrySecAgentACIDMap(pSecAgentACIDMap *thost.CThostFtdcSecAgentACIDMapField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySecAgentACIDMapCallback != nil {
		s.OnRspQrySecAgentACIDMapCallback(pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询产品报价汇率
func (s *TDSpiN) OnRspQryProductExchRate(pProductExchRate *thost.CThostFtdcProductExchRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryProductExchRateCallback != nil {
		s.OnRspQryProductExchRateCallback(pProductExchRate, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询产品组
func (s *TDSpiN) OnRspQryProductGroup(pProductGroup *thost.CThostFtdcProductGroupField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryProductGroupCallback != nil {
		s.OnRspQryProductGroupCallback(pProductGroup, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询做市商合约手续费率响应
func (s *TDSpiN) OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *thost.CThostFtdcMMInstrumentCommissionRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryMMInstrumentCommissionRateCallback != nil {
		s.OnRspQryMMInstrumentCommissionRateCallback(pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询做市商期权合约手续费响应
func (s *TDSpiN) OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *thost.CThostFtdcMMOptionInstrCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryMMOptionInstrCommRateCallback != nil {
		s.OnRspQryMMOptionInstrCommRateCallback(pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询报单手续费响应
func (s *TDSpiN) OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *thost.CThostFtdcInstrumentOrderCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryInstrumentOrderCommRateCallback != nil {
		s.OnRspQryInstrumentOrderCommRateCallback(pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询资金账户响应
func (s *TDSpiN) OnRspQrySecAgentTradingAccount(pTradingAccount *thost.CThostFtdcTradingAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySecAgentTradingAccountCallback != nil {
		s.OnRspQrySecAgentTradingAccountCallback(pTradingAccount, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询二级代理商资金校验模式响应
func (s *TDSpiN) OnRspQrySecAgentCheckMode(pSecAgentCheckMode *thost.CThostFtdcSecAgentCheckModeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySecAgentCheckModeCallback != nil {
		s.OnRspQrySecAgentCheckModeCallback(pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询二级代理商信息响应
func (s *TDSpiN) OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo *thost.CThostFtdcSecAgentTradeInfoField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQrySecAgentTradeInfoCallback != nil {
		s.OnRspQrySecAgentTradeInfoCallback(pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询期权交易成本响应
func (s *TDSpiN) OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost *thost.CThostFtdcOptionInstrTradeCostField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptionInstrTradeCostCallback != nil {
		s.OnRspQryOptionInstrTradeCostCallback(pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询期权合约手续费响应
func (s *TDSpiN) OnRspQryOptionInstrCommRate(pOptionInstrCommRate *thost.CThostFtdcOptionInstrCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptionInstrCommRateCallback != nil {
		s.OnRspQryOptionInstrCommRateCallback(pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询执行宣告响应
func (s *TDSpiN) OnRspQryExecOrder(pExecOrder *thost.CThostFtdcExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryExecOrderCallback != nil {
		s.OnRspQryExecOrderCallback(pExecOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询询价响应
func (s *TDSpiN) OnRspQryForQuote(pForQuote *thost.CThostFtdcForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryForQuoteCallback != nil {
		s.OnRspQryForQuoteCallback(pForQuote, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询报价响应
func (s *TDSpiN) OnRspQryQuote(pQuote *thost.CThostFtdcQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryQuoteCallback != nil {
		s.OnRspQryQuoteCallback(pQuote, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询期权自对冲响应
func (s *TDSpiN) OnRspQryOptionSelfClose(pOptionSelfClose *thost.CThostFtdcOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryOptionSelfCloseCallback != nil {
		s.OnRspQryOptionSelfCloseCallback(pOptionSelfClose, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询投资单元响应
func (s *TDSpiN) OnRspQryInvestUnit(pInvestUnit *thost.CThostFtdcInvestUnitField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryInvestUnitCallback != nil {
		s.OnRspQryInvestUnitCallback(pInvestUnit, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询组合合约安全系数响应
func (s *TDSpiN) OnRspQryCombInstrumentGuard(pCombInstrumentGuard *thost.CThostFtdcCombInstrumentGuardField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCombInstrumentGuardCallback != nil {
		s.OnRspQryCombInstrumentGuardCallback(pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询申请组合响应
func (s *TDSpiN) OnRspQryCombAction(pCombAction *thost.CThostFtdcCombActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCombActionCallback != nil {
		s.OnRspQryCombActionCallback(pCombAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询转帐流水响应
func (s *TDSpiN) OnRspQryTransferSerial(pTransferSerial *thost.CThostFtdcTransferSerialField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTransferSerialCallback != nil {
		s.OnRspQryTransferSerialCallback(pTransferSerial, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询银期签约关系响应
func (s *TDSpiN) OnRspQryAccountregister(pAccountregister *thost.CThostFtdcAccountregisterField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryAccountregisterCallback != nil {
		s.OnRspQryAccountregisterCallback(pAccountregister, pRspInfo, nRequestID, bIsLast)
	}
}

// 错误应答
func (s *TDSpiN) OnRspError(pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorCallback != nil {
		s.OnRspErrorCallback(pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询签约银行响应
func (s *TDSpiN) OnRspQryContractBank(pContractBank *thost.CThostFtdcContractBankField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryContractBankCallback != nil {
		s.OnRspQryContractBankCallback(pContractBank, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询预埋单响应
func (s *TDSpiN) OnRspQryParkedOrder(pParkedOrder *thost.CThostFtdcParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryParkedOrderCallback != nil {
		s.OnRspQryParkedOrderCallback(pParkedOrder, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询预埋撤单响应
func (s *TDSpiN) OnRspQryParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryParkedOrderActionCallback != nil {
		s.OnRspQryParkedOrderActionCallback(pParkedOrderAction, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询交易通知响应
func (s *TDSpiN) OnRspQryTradingNotice(pTradingNotice *thost.CThostFtdcTradingNoticeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryTradingNoticeCallback != nil {
		s.OnRspQryTradingNoticeCallback(pTradingNotice, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询经纪公司交易参数响应
func (s *TDSpiN) OnRspQryBrokerTradingParams(pBrokerTradingParams *thost.CThostFtdcBrokerTradingParamsField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryBrokerTradingParamsCallback != nil {
		s.OnRspQryBrokerTradingParamsCallback(pBrokerTradingParams, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询经纪公司交易算法响应
func (s *TDSpiN) OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos *thost.CThostFtdcBrokerTradingAlgosField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryBrokerTradingAlgosCallback != nil {
		s.OnRspQryBrokerTradingAlgosCallback(pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询监控中心用户令牌
func (s *TDSpiN) OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *thost.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryCFMMCTradingAccountTokenCallback != nil {
		s.OnRspQueryCFMMCTradingAccountTokenCallback(pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast)
	}
}

// 期货发起银行资金转期货应答
func (s *TDSpiN) OnRspFromBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspFromBankToFutureByFutureCallback != nil {
		s.OnRspFromBankToFutureByFutureCallback(pReqTransfer, pRspInfo, nRequestID, bIsLast)
	}
}

// 期货发起期货资金转银行应答
func (s *TDSpiN) OnRspFromFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspFromFutureToBankByFutureCallback != nil {
		s.OnRspFromFutureToBankByFutureCallback(pReqTransfer, pRspInfo, nRequestID, bIsLast)
	}
}

// 期货发起查询银行余额应答
func (s *TDSpiN) OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQueryBankAccountMoneyByFutureCallback != nil {
		s.OnRspQueryBankAccountMoneyByFutureCallback(pReqQueryAccount, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询分类合约响应
func (s *TDSpiN) OnRspQryClassifiedInstrument(pInstrument *thost.CThostFtdcInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryClassifiedInstrumentCallback != nil {
		s.OnRspQryClassifiedInstrumentCallback(pInstrument, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求组合优惠比例响应
func (s *TDSpiN) OnRspQryCombPromotionParam(pCombPromotionParam *thost.CThostFtdcCombPromotionParamField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryCombPromotionParamCallback != nil {
		s.OnRspQryCombPromotionParamCallback(pCombPromotionParam, pRspInfo, nRequestID, bIsLast)
	}
}

// 投资者风险结算持仓查询响应
func (s *TDSpiN) OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *thost.CThostFtdcRiskSettleInvstPositionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryRiskSettleInvstPositionCallback != nil {
		s.OnRspQryRiskSettleInvstPositionCallback(pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast)
	}
}

// 风险结算产品查询响应
func (s *TDSpiN) OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus *thost.CThostFtdcRiskSettleProductStatusField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryRiskSettleProductStatusCallback != nil {
		s.OnRspQryRiskSettleProductStatusCallback(pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast)
	}
}
