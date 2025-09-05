package CTPBase

import (
	"github.com/li-wancai/EQCTPQH/Thost"

	"github.com/pseudocodes/go2ctp/thost"
)

func (s *TDSpiN) DoReqQryInvestorPosition( //持仓查询
	InstrumentID /*合约代码*/, ExchangeID /*交易所代码*/ string) int {
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	HKeys := s.RdSt.HKeyS("持仓明细")
	if len(HKeys) > 0 {
		s.RdSt.HDel("持仓明细", HKeys)
	}
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	s.ReqSleep()
	s.RequestID++
	req := &thost.CThostFtdcQryInvestorPositionField{}
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.InvestorID[:], s.Cfg.InvestorID)
	copy(req.ExchangeID[:], ExchangeID)
	copy(req.InstrumentID[:], InstrumentID)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	rst := s.Api.ReqQryInvestorPosition(req, s.RequestID) // 报单录入请求
	RspRst(rst, "【%s】【持仓查询】请求 ReqQryInvestorPosition 已发送", s.Cfg.UserID)
	return rst
}

func (s *TDSpiN) DoReqOrderAction( //撤单
	IPAddress /*IP地址*/, MacAddress, /*Mac地址*/
	InstrumentID /*合约代码*/, ExchangeID, /*交易所代码*/
	OrderRef /*报单引用*/, OrderSysID /*报单编号注意原值有空格*/ string,
	FrontID /*前置编号*/, SessionID /*会话编号*/ int) int {
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	s.ReqSleep()
	s.RequestID++
	req := &thost.CThostFtdcInputOrderActionField{}
	copy(req.UserID[:], s.Cfg.UserID)
	copy(req.IPAddress[:], IPAddress)
	copy(req.MacAddress[:], MacAddress)
	copy(req.ExchangeID[:], ExchangeID) //  交易所代码
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.InstrumentID[:], InstrumentID) //  合约代码
	copy(req.InvestorID[:], s.Cfg.InvestorID)
	req.RequestID = thost.TThostFtdcRequestIDType(s.RequestID)           //  请求编号
	req.OrderActionRef = thost.TThostFtdcOrderActionRefType(s.RequestID) //  报单操作引用
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	copy(req.OrderRef[:], OrderRef)               //  报单引用 ★★ 使用FrontID+SessionID+OrderRef撤单
	copy(req.OrderSysID[:], OrderSysID)           //  报单编号 ★★ 使用OrderSysID撤单（推荐使用）注意原值有空格如："      252833"
	req.ActionFlag = thost.THOST_FTDC_AF_Delete   //  操作标志 只支持删除
	req.ActionFlag = Thost.ToActionFlag("Delete") //  操作标志 只支持删除
	if OrderRef != "" {
		req.FrontID = thost.TThostFtdcFrontIDType(FrontID)       //  前置编号
		req.SessionID = thost.TThostFtdcSessionIDType(SessionID) //  会话编号
	}
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	rst := s.Api.ReqOrderAction(req, s.RequestID) // 报单录入请求
	RspRst(rst, "【%s】【撤单】请求 ReqOrderAction 已发送", s.Cfg.UserID)
	return rst
}
