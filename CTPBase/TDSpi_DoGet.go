package CTPBase

import (
	"strings"

	"github.com/li-wancai/GoScripts/Formulae"
)

func (s *TDSpiN) DoGetOrders( //订单查询
	InstrumentID /*合约代码*/ string, OnlyCancelAble /*仅可撤订单*/ bool) map[string]string {
	OKData := make(map[string]string, 0)
	HKeyVS := s.RdSt.HGetall("报单明细")
	if len(HKeyVS) < 1 {
		return OKData
	}
	for key, strinfo := range HKeyVS {
		if InstrumentID != "" {
			if !strings.Contains(key, InstrumentID) {
				continue
			}
		}
		if OnlyCancelAble {
			rst := Formulae.EVALI(strinfo) //格式化处理
			OrderType := rst.(map[string]interface{})["状态信息"].(string)
			if !strings.Contains(OrderType, "未成交") && !strings.Contains(OrderType, "部分成交") {
				continue
			}
		}
		OKData[key] = strinfo
	}
	return OKData
}
func (s *TDSpiN) DoGetTrades(InstrumentID /*合约代码*/ string) map[string]string { //成交查询
	OKData := make(map[string]string, 0)
	HKeyVS := s.RdSt.HGetall("报单明细")
	if len(HKeyVS) < 1 {
		return OKData
	}
	for key, strinfo := range HKeyVS {
		if InstrumentID != "" {
			if !strings.Contains(key, InstrumentID) {
				continue
			}
		}
		rst := Formulae.EVALI(strinfo) //格式化处理
		OrderType := rst.(map[string]interface{})["状态信息"].(string)
		if !strings.Contains(OrderType, "全部成交") && !strings.Contains(OrderType, "部分成交") {
			continue
		}
		OKData[key] = strinfo
	}
	return OKData
}
