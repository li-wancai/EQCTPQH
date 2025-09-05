package CTPBase

import (
	"fmt"

	"github.com/li-wancai/EQUseApi"

	"github.com/pseudocodes/go2ctp/thost"
)

// 登录请求响应
func (s *MDSpiN) OnRspUserLogin(
	pRspUserLogin *thost.CThostFtdcRspUserLoginField,
	pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if pRspInfo.ErrorID == 0 { // 检查登录响应是否成功
		msg := fmt.Sprintf("【行情登录】交易日: %v 成功", pRspUserLogin.TradingDay)
		EQUseApi.SendLogTxT(msg, EQUseApi.SendToGroupList, "INFO")
		s.OnData(map[string]interface{}{}, true, "登录成功")
	} else {
		msg := fmt.Sprintf("【行情登录】失败 错误代码：%d, 错误信息：%s",
			pRspInfo.ErrorID, pRspInfo.ErrorMsg.GBString())
		EQUseApi.SendLogTxT(msg, EQUseApi.SendToGroupList, "ERROR")
	}
}

// 登出请求响应
func (s *MDSpiN) OnRspUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUserLogoutCallback != nil {
		s.OnRspUserLogoutCallback(pUserLogout, pRspInfo, nRequestID, bIsLast)
	}
}

// 请求查询组播合约响应
func (s *MDSpiN) OnRspQryMulticastInstrument(pMulticastInstrument *thost.CThostFtdcMulticastInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspQryMulticastInstrumentCallback != nil {
		s.OnRspQryMulticastInstrumentCallback(pMulticastInstrument, pRspInfo, nRequestID, bIsLast)
	}
}

// 错误应答
func (s *MDSpiN) OnRspError(pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspErrorCallback != nil {
		s.OnRspErrorCallback(pRspInfo, nRequestID, bIsLast)
	}
}

// 订阅行情应答
func (s *MDSpiN) OnRspSubMarketData(
	pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField,
	pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if ErrorRspInfo(pRspInfo) || pSpecificInstrument == nil {
		return
	}
	InstrumentID := pSpecificInstrument.InstrumentID.String()
	if _, ok := s.CodeMDchan[InstrumentID]; !ok {
		s.CodeMDchan[InstrumentID] = make(chan interface{}, 10)        // 创建一个新的通道
		go s.OneCodeDoMDWork(s.CodeMDchan[InstrumentID], InstrumentID) // 启动协程处理数据
		log.Infof("【%s】【后台任务】: go run OneCodeDoMDWork", InstrumentID)
	}

}

// 取消订阅行情应答
func (s *MDSpiN) OnRspUnSubMarketData(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUnSubMarketDataCallback != nil {
		s.OnRspUnSubMarketDataCallback(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
	}
}

// 订阅询价应答
func (s *MDSpiN) OnRspSubForQuoteRsp(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspSubForQuoteRspCallback != nil {
		s.OnRspSubForQuoteRspCallback(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
	}
}

// 取消订阅询价应答
func (s *MDSpiN) OnRspUnSubForQuoteRsp(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	if s.OnRspUnSubForQuoteRspCallback != nil {
		s.OnRspUnSubForQuoteRspCallback(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
	}
}
