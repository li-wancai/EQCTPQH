package CTPBase

import (
	"EarnQuant/GoScripts/EQUseApi"
	"fmt"
	"strings"

	"github.com/li-wancai/GoScripts/Formulae"

	"github.com/pseudocodes/go2ctp/thost"
)

func (s *TDSpiN) OnData(data interface{}, bIsLast bool, v ...interface{}) {
	if len(v) > 0 {
		s.DoReqSettlementInfoConfirm(v[0].(string))         // 登录成功后的结算确认请求
		s.SaveTradingAccount(data, v[0].(string))           // 保存账户信息到redis
		s.SaveInstrumentInfo(data, v[0].(string))           // 保存合约信息到redis
		s.SaveCommissionRate(data, v[0].(string))           // 保存手续费率到redis
		s.SaveMarginRate(data, v[0].(string))               // 保存保证金率到redis
		s.SaveOnRtnOrder(data, v[0].(string))               // 保存委托通知到redis
		s.SaveOnRspQryInvestorPosition(data, v[0].(string)) // 保存持仓通知到redis
		if strings.Contains("登录成功,合约信息,资金账户,手续费率,保证金率", v[0].(string)) {
			s.InitBase(bIsLast) //查询账户信息 => 查询合约信息 => 查询保证金信息 => 查询手续费信息
		}
	}
}

// 登录成功后 初始化任务：查询账户信息 => 查询合约信息 => 查询保证金信息 => 查询手续费信息
func (s *TDSpiN) InitBase(bIsLast bool) {
	if !bIsLast || s.DoList.End {
		return
	}
	if s.DoList.Item != "手续费率" && s.DoList.Item != "保证金率" {
		s.DoList.Next()
	}
	if s.DoList.Item == "合约信息" {
		s.DoReqQryInstrument("", "", "", "")
	}
	if s.DoList.Item == "资金账户" {
		CodeInfos := s.RdSt.HKeyS("合约信息")
		s.FeeList = Formulae.IterList(CodeInfos) //手续费查询列表
		s.MrgList = Formulae.IterList(CodeInfos) //保证金查询列表
		s.DoReqQryTradingAccount()               //交易账户信息查询
	}
	if s.DoList.Item == "手续费率" {
		s.FeeList.Next()
		if s.FeeList.End {
			s.DoList.Next()
			msg := fmt.Sprintf("【%s】获取所有 【手续费率】 信息任务完成", s.Cfg.UserID)
			EQUseApi.SendLogTxT(msg, EQUseApi.SendToGroupList, "INFO")
		} else {
			s.DoReqQryInstrumentCommissionRate(s.FeeList.Item)
		}
	}
	if s.DoList.Item == "保证金率" {
		s.MrgList.Next()
		if s.MrgList.End {
			s.DoList.Next()
			msg := fmt.Sprintf("【%s】获取所有 【保证金率】 信息任务完成", s.Cfg.UserID)
			EQUseApi.SendLogTxT(msg, EQUseApi.SendToGroupList, "INFO")

		} else {
			s.DoReqQryInstrumentMarginRate(s.MrgList.Item)
		}
	}
}
func (s *TDSpiN) DoReqQryTradingAccount() {
	s.ReqSleep()
	s.RequestID++
	req := &thost.CThostFtdcQryTradingAccountField{}
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.InvestorID[:], s.Cfg.InvestorID)
	rst := s.Api.ReqQryTradingAccount(req, s.RequestID)
	RspRst(rst, "【%s】【资金账户】请求已发送", s.Cfg.UserID)
}

// 请求所有的合约信息
func (s *TDSpiN) DoReqQryInstrument( //请求合约信息
	ExchangeInstID string, InstrumentID string, ExchangeID string, ProductID string) {
	s.ReqSleep()
	s.RequestID++
	req := &thost.CThostFtdcQryInstrumentField{}
	copy(req.ExchangeInstID[:], ExchangeInstID) //合约在交易所的代码
	copy(req.InstrumentID[:], InstrumentID)     //合约代码
	copy(req.ExchangeID[:], ExchangeID)         //交易所代码
	copy(req.ProductID[:], ProductID)           //产品代码
	rst := s.Api.ReqQryInstrument(req, s.RequestID)
	RspRst(rst, "【%s】【合约查询】请求已发送", s.Cfg.UserID)
}
func (s *TDSpiN) SaveInstrumentInfo(data interface{}, vmsg string) {
	if vmsg == "合约信息" {
		Info := data.(map[string]interface{})
		if Info["产品类型"] != "Futures" { //不是期货
			return
		}
		InstrumentID := Info["InstrumentID"].(string)
		rst := map[string]interface{}{InstrumentID: Formulae.DictToStrJson(data.(map[string]interface{}))}
		err := s.RdSt.HmSet("合约信息", rst)
		if err != nil {
			log.Errorf("将%s 合约信息 数据存入redis失败: %v", InstrumentID, err)
		}
	}
}
func (s *TDSpiN) SaveMarginRate(data interface{}, vmsg string) {
	if vmsg == "保证金率" {
		Info := data.(map[string]interface{})
		InstrumentID := Info["InstrumentID"].(string)
		rst := map[string]interface{}{InstrumentID: Formulae.DictToStrJson(Info)}
		err := s.RdSt.HmSet("保证金率", rst)
		if err != nil {
			log.Errorf("将%s 保证金率 数据存入redis失败: %v", InstrumentID, err)
		}
	}
}

func (s *TDSpiN) SaveCommissionRate(data interface{}, vmsg string) {
	if vmsg == "手续费率" {
		Info := data.(map[string]interface{})
		InstrumentID := Info["InstrumentID"].(string)
		rst := map[string]interface{}{InstrumentID: Formulae.DictToStrJson(Info)}
		err := s.RdSt.HmSet("手续费率", rst)
		if err != nil {
			log.Errorf("将%s 手续费率 数据存入redis失败: %v", InstrumentID, err)
		}
	}
}
func (s *TDSpiN) SaveOnRtnOrder(data interface{}, vmsg string) {
	if vmsg == "OnRtnOrder" {
		Info := data.(map[string]interface{})
		LocalID := strings.ReplaceAll(Info["本地报单编号"].(string), " ", "")
		OrderID := strings.ReplaceAll(fmt.Sprintf("%s", Info["报单编号"]), " ", "0")
		MKey := fmt.Sprintf("%s_%s_%s", Info["InstrumentID"].(string), LocalID, OrderID)
		rst := map[string]interface{}{MKey: Formulae.DictToStrJson(Info)}
		err := s.RdSt.HmSet("报单明细", rst)
		if err != nil {
			log.Errorf("将%s 报单明细 数据存入redis失败: %v", MKey, err)
		}
	}
}
func (s *TDSpiN) SaveOnRspQryInvestorPosition(data interface{}, vmsg string) {
	if vmsg == "OnRspQryInvestorPosition" {
		Info := data.(map[string]interface{})
		if Info["持仓成本"] == 0 {
			return
		}
		InstrumentID := Info["InstrumentID"].(string)
		rst := map[string]interface{}{InstrumentID: Formulae.DictToStrJson(Info)}
		err := s.RdSt.HmSet("持仓明细", rst)
		if err != nil {
			log.Errorf("将%s 持仓明细 数据存入redis失败: %v", InstrumentID, err)
		}
	}
}
func (s *TDSpiN) SaveTradingAccount(data interface{}, vmsg string) {
	if vmsg == "资金账户" {
		Info := data.(map[string]interface{})
		AccountID := Info["投资者帐号"].(string)
		Accountmsg := Formulae.DictToStrJson(Info)
		EQUseApi.SendLogTxT(Accountmsg, EQUseApi.SendToGroupList, "INFO")
		rst := map[string]interface{}{AccountID: Accountmsg}
		err := s.RdSt.HmSet("资金账户", rst)
		if err != nil {
			log.Errorf("将%s 资金账户 数据存入redis失败: %v", AccountID, err)
		}
	}
}
func (s *TDSpiN) DoReqSettlementInfoConfirm(vmsg string) {
	//登录成功 ==>> 投资者结算结果确认
	if vmsg == "登录成功" {
		s.ReqSleep()
		s.RequestID++ //请求序列加1
		req := &thost.CThostFtdcSettlementInfoConfirmField{}
		copy(req.BrokerID[:], s.Cfg.BrokerID)
		copy(req.InvestorID[:], s.Cfg.InvestorID)
		s.Api.ReqSettlementInfoConfirm(req, s.RequestID) //投资者结算结果确认
	}
}

// 保证金率
func (s *TDSpiN) DoReqQryInstrumentMarginRate(Code string) {
	s.ReqSleep()
	s.RequestID++
	req := &thost.CThostFtdcQryInstrumentMarginRateField{
		HedgeFlag: thost.THOST_FTDC_HF_Speculation}
	copy(req.InstrumentID[:], Code)
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.InvestorID[:], s.Cfg.InvestorID)
	rst := s.Api.ReqQryInstrumentMarginRate(req, s.RequestID)
	RspRst(rst, "【%s】【合约%s保证金率查询】请求已发送", s.Cfg.UserID, Code)
}

// 手续费率
func (s *TDSpiN) DoReqQryInstrumentCommissionRate(Code string) {
	s.ReqSleep()
	s.RequestID++
	req := &thost.CThostFtdcQryInstrumentCommissionRateField{}
	copy(req.InstrumentID[:], Code)
	copy(req.BrokerID[:], s.Cfg.BrokerID)
	copy(req.InvestorID[:], s.Cfg.InvestorID)
	rst := s.Api.ReqQryInstrumentCommissionRate(req, s.RequestID)
	RspRst(rst, "【%s】【合约%s手续费率查询】请求已发送", s.Cfg.UserID, Code)
}
