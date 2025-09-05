package CTPBase

import "github.com/pseudocodes/go2ctp/thost"

// 深度行情通知
func (s *MDSpiN) OnRtnDepthMarketData(pDepthMarketData *thost.CThostFtdcDepthMarketDataField) {
	if pDepthMarketData == nil {
		return
	}
	DepthMarketData := map[string]interface{}{
		"TradingDay":         pDepthMarketData.TradingDay.String(),
		"ActionDay":          pDepthMarketData.ActionDay.String(),
		"LowestPrice":        pDepthMarketData.LowestPrice.String(),
		"Volume":             pDepthMarketData.Volume,
		"BidPrice1":          pDepthMarketData.BidPrice1.String(),
		"BidVolume1":         pDepthMarketData.BidVolume1,
		"AskPrice1":          pDepthMarketData.AskPrice1.String(),
		"AskVolume1":         pDepthMarketData.AskVolume1,
		"TurnOver":           pDepthMarketData.Turnover.String(),
		"OpenInterest":       pDepthMarketData.OpenInterest.String(),
		"OpenPrice":          pDepthMarketData.OpenPrice.String(),
		"HighestPrice":       pDepthMarketData.HighestPrice.String(),
		"UpdateTime":         pDepthMarketData.UpdateTime.String(),
		"UpdateMillisec":     pDepthMarketData.UpdateMillisec,
		"AveragePrice":       pDepthMarketData.AveragePrice.String(),
		"InstrumentID":       pDepthMarketData.InstrumentID.String(),
		"LastPrice":          pDepthMarketData.LastPrice.String(),
		"PreSettlementPrice": pDepthMarketData.PreSettlementPrice.String(),
		"SettlementPrice":    pDepthMarketData.SettlementPrice.String(),
		"PreDelta":           pDepthMarketData.PreDelta.String(),
		"CurrDelta":          pDepthMarketData.CurrDelta.String(),
		"ClosePrice":         pDepthMarketData.ClosePrice.String(),
		"PreClosePrice":      pDepthMarketData.PreClosePrice.String(),
		"PreOpenInterest":    pDepthMarketData.PreOpenInterest.String(),
		"UpperLimitPrice":    pDepthMarketData.UpperLimitPrice.String(),
		"LowerLimitPrice":    pDepthMarketData.LowerLimitPrice.String(),
		"BandingUpperPrice":  pDepthMarketData.BandingUpperPrice.String(),
		"BandingLowerPrice":  pDepthMarketData.BandingLowerPrice.String(),
	}
	s.OnData(DepthMarketData, true, "行情数据")
}

// 询价通知
func (s *MDSpiN) OnRtnForQuoteRsp(pForQuoteRsp *thost.CThostFtdcForQuoteRspField) {
	if s.OnRtnForQuoteRspCallback != nil {
		s.OnRtnForQuoteRspCallback(pForQuoteRsp)
	}
}
