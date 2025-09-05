package CTPBase

import (
	"time"

	"github.com/li-wancai/EQDBLinks/DBRedis"

	"github.com/li-wancai/EQCTPQH/Thost"

	"github.com/li-wancai/GoScripts/Formulae"

	"github.com/pseudocodes/go2ctp/ctp"
	"github.com/pseudocodes/go2ctp/thost"
)

type TDSpiN struct {
	RequestID int
	ReqTime   time.Time
	ctp.BaseTraderSpi
	Api     thost.TraderApi
	Cfg     Thost.CTPBaseCfg
	RdSt    *DBRedis.RdStN              //存储实例redis
	DoList  *Formulae.IterListN[string] //初始化任务列表
	FeeList *Formulae.IterListN[string] //手续费查询列表
	MrgList *Formulae.IterListN[string] //保证金查询列表

}

// 初始化
func InitTDSpi(config map[string]interface{}, KeyName string, RdStN *DBRedis.RdStN) *TDSpiN {
	if KeyName == "" {
		KeyName = config["EQCTPCfg_Account"].(string)
	}
	CTPCfg := CTPConfig(
		KeyName,
		config["EQCTPCfg_FileName"].(string),
		config["EQCTPCfg_TomlPath"].(string))
	return CreateTDSpi(CTPCfg, RdStN)
}
func CreateTDSpi(cfg Thost.CTPBaseCfg, RdStN *DBRedis.RdStN) *TDSpiN {
	TDSpi := &TDSpiN{
		RequestID: 1,
		Cfg:       cfg,
		RdSt:      RdStN,
		ReqTime:   time.Now(),
		DoList:    Formulae.IterList([]string{"合约信息", "资金账户", "手续费率", "保证金率"}), //执行任务：###查询合约信息==>>查询资金账户==>>查询手续费==>>查询保证金==>>等待交易

	}
	return TDSpi
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *TDSpiN) OnFrontConnected() {
	s.OnFrontConnectedCallback()
}
func (s *TDSpiN) ReqSleep(n ...time.Duration) {
	interval := 1 * time.Second
	if len(n) > 0 {
		interval = n[0]
	}
	if time.Since(s.ReqTime) <= interval {
		time.Sleep(interval - time.Since(s.ReqTime))
	}
	s.ReqTime = time.Now()
}

func (s *TDSpiN) OnFrontConnectedCallback() {
	//连接成功后==>> 登录
	s.RequestID++
	loginR := &thost.CThostFtdcReqAuthenticateField{}
	copy(loginR.AppID[:], s.Cfg.AppID)
	copy(loginR.UserID[:], s.Cfg.UserID)
	copy(loginR.BrokerID[:], s.Cfg.BrokerID)
	copy(loginR.AuthCode[:], s.Cfg.AuthCode)
	copy(loginR.UserProductInfo[:], s.Cfg.ProductInfo)
	rst := s.Api.ReqAuthenticate(loginR, s.RequestID)
	RspRst(rst, "【交易登录】%s请求已发送", s.Cfg.UserID)
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (s *TDSpiN) OnFrontDisconnected(nReason int) {
	/*
		当客户端与交易后台通信连接断开时，该方法被调用。
		当发生这个情况后，API会自动重新连接，客户端可不做处理。
		@param nReason 错误原因
				0x1001 网络读失败
				0x1002 网络写失败
				0x2001 接收心跳超时
				0x2002 发送心跳失败
				0x2003 收到错误报文
	*/
	ErrReason(nReason)
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
// /@param nTimeLapse 距离上次接收报文的时间
func (s *TDSpiN) OnHeartBeatWarning(nTimeLapse int) {
	log.Warnf("心跳警告: 已经%d毫秒未收到报文", nTimeLapse)
}
