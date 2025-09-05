package CTPBase

import (
	"path/filepath"
	"sort"
	"time"

	"github.com/li-wancai/GoScripts/DirsFile"

	"github.com/li-wancai/GoScripts/Formulae"
)

func (s *MDSpiN) OnData(data interface{}, bIsLast bool, v ...interface{}) {
	if len(v) > 0 {
		s.DoSubscribeMarketData(v[0].(string))      // 登录成功后的行情订阅
		s.PutDataToOneCodeChan(data, v[0].(string)) // 将数据发送到通道
	}
}

// 登录成功 ==>> 【行情订阅】
func (s *MDSpiN) DoSubscribeMarketData(vmsg string) {
	//登录成功 ==>> 【行情订阅】
	if vmsg == "登录成功" {
		CodeInfos := s.RdSt.HKeyS("合约信息")
		rst := s.Api.SubscribeMarketData(CodeInfos...)
		RspRst(rst, "【行情订阅】%s请求已发送", s.Cfg.UserID)
	}
}

// 发送市场数据到对应的通道
func (s *MDSpiN) PutDataToOneCodeChan(data interface{}, vmsg string) {
	if vmsg == "行情数据" {
		InstrumentID := data.(map[string]interface{})["InstrumentID"].(string)
		if ch, ok := s.CodeMDchan[InstrumentID]; ok {
			ch <- data
		}
	}
}

// 协程处理市场数据
func (s *MDSpiN) OneCodeDoMDWork(ch chan interface{}, InstrumentID string) {
	Keys, CodeCsv, HeaderWritten := s.NewCodeCsv(InstrumentID)
	defer func() {
		CodeCsv.Flush()
		CodeCsv.Close()
		close(ch) // 确保协程结束时关闭通道
	}()
	for {
		select {
		/////////////////////////////////////////////////////////////////////////
		// 从通道中获取数据
		case data, ok := <-ch:
			/////////////////////////////////////////////////////////////////////
			if !ok {
				log.Errorf("%s通道已经被关闭", InstrumentID)
				return
			}
			/////////////////////////////////////////////////////////////////////
			jsdt := data.(map[string]interface{})  // 格式化数据
			jsdt["TradingDay"] = s.Cfg.TradingDay  // 更新交易日
			jsdt, _ = Formulae.FilterCTPData(jsdt) // 过滤数据
			rst := Formulae.DictToStrJson(jsdt)    // 格式化数据
			s.RdSt.HSet("最新行情", InstrumentID, rst) // 保存到Redis
			s.RdSt.PubMsg(InstrumentID, rst)       // 通过Redis发布实时行情数据
			s.SSEs.PushToRoom(InstrumentID, s.SSEs.ToSSeEvent(rst, "TickData_QH", InstrumentID, 300))
			s.SSEs.PushToRoom("All_TickData_QH", s.SSEs.ToSSeEvent(rst, "TickData_QH", InstrumentID, 300))
			/////////////////////////////////////////////////////////////////////
			if !HeaderWritten { // 检查是否需要写入表头
				Keys = Formulae.GETKEY(jsdt)
				sort.Strings(Keys) // 要排序统一顺序
				CodeCsv.Writer(Keys)
				HeaderWritten = true
			}
			values := Formulae.GetDictKeysValues(jsdt, Keys)
			CodeCsv.Writer(values) // 写入行情数据
			/////////////////////////////////////////////////////////////////////
			if Formulae.NoTradingTime(s.Cfg.NoTradeTimeRange) { // 检查当前时间是否超过休盘时间
				log.Infof("休盘时间已到，停止处理%s市场数据", InstrumentID)
				CodeCsv.Flush()
				return
			}
		/////////////////////////////////////////////////////////////////////////
		case <-time.After(5 * time.Second): // 每5秒检查一次时间，防止长时间阻塞导致错过休盘时间
			CodeCsv.Flush() //主动刷新存储到目录
			if Formulae.NoTradingTime(s.Cfg.NoTradeTimeRange) {
				log.Infof("休盘时间已到，停止处理%s市场数据", InstrumentID)
				return
			}
		}
	}
}

func (s *MDSpiN) NewCodeCsv(InstrumentID string) ([]string, *DirsFile.TodoCsvN, bool) {
	Keys := []string{}
	s.Cfg.TradingDay = s.Api.GetTradingDay()
	FilePath := filepath.Join(s.Cfg.DataPath, s.Cfg.TradingDay)
	CodeCsv, err := DirsFile.TodoCsv(InstrumentID, FilePath, "GB18030", "a+", "", 0)
	if err != nil {
		log.Errorf("创建【%s】本地Csv存储文件失败: %s", InstrumentID, err)
	}
	HeaderWritten := false
	if CodeCsv.GetSize() > 0 {
		log.Infof("【%s】本地Csv存储文件已存在", InstrumentID)
		HeaderWritten = true
		Keys, err = CodeCsv.GetHeader()
		if err != nil {
			log.Errorf("获取表头失败:%s", err)
			HeaderWritten = false
		}
	}
	return Keys, CodeCsv, HeaderWritten
}
