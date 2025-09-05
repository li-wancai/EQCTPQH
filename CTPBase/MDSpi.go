package CTPBase

import (
	"github.com/li-wancai/EQSSEHub"

	"github.com/li-wancai/EQDBTools/DBRedis"

	"github.com/li-wancai/EQCTPQH/Thost"

	"github.com/pseudocodes/go2ctp/ctp"
	"github.com/pseudocodes/go2ctp/thost"
)

type MDSpiN struct {
	RequestID int
	ctp.BaseMdSpi
	Api  thost.MdApi
	RdSt *DBRedis.RdStN
	SSEs *EQSSEHub.PusherN //SSE推送实例
	Cfg  Thost.CTPBaseCfg

	CodeMDchan map[string]chan interface{} // 保存每个合约代码对应的通道

}

// 初始化
func InitMDSpi(config map[string]interface{}, KeyName string, RdStN *DBRedis.RdStN, PusherN *EQSSEHub.PusherN) *MDSpiN {
	if KeyName == "" {
		KeyName = config["EQCTPCfg_Account"].(string)
	}
	CTPCfg := CTPConfig(
		KeyName,
		config["EQCTPCfg_FileName"].(string),
		config["EQCTPCfg_TomlPath"].(string))
	return CreateMDSpi(CTPCfg, RdStN, PusherN)
}
func CreateMDSpi(cfg Thost.CTPBaseCfg, RdStN *DBRedis.RdStN, PusherN *EQSSEHub.PusherN) *MDSpiN {
	MDSpi := &MDSpiN{
		Cfg:        cfg,
		RdSt:       RdStN,
		SSEs:       PusherN,
		RequestID:  1,
		CodeMDchan: make(map[string]chan interface{}),
	}
	return MDSpi
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (s *MDSpiN) OnHeartBeatWarning(nTimeLapse int) {
	log.Warnf("心跳警告: 已经%d毫秒未收到报文", nTimeLapse)
}

// 建立连接
func (s *MDSpiN) OnFrontConnected() {
	s.OnFrontConnectedCallback() //当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *MDSpiN) OnFrontConnectedCallback() {
	//连接成功后 ==>> 登录
	s.RequestID++
	loginR := &thost.CThostFtdcReqUserLoginField{}
	copy(loginR.UserID[:], s.Cfg.UserID)
	copy(loginR.BrokerID[:], s.Cfg.BrokerID)
	copy(loginR.Password[:], s.Cfg.Password)
	copy(loginR.UserProductInfo[:], s.Cfg.ProductInfo)
	rst := s.Api.ReqUserLogin(loginR, s.RequestID)
	RspRst(rst, "【行情登录】%s请求已发送", s.Cfg.UserID)

}

// 断线重连
func (s *MDSpiN) OnFrontDisconnected(nReason int) {
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
	s.OnFrontDisconnectedCallback(nReason) //未连接回调函数,用于断线重连

}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (s *MDSpiN) OnFrontDisconnectedCallback(nReason int) {
}
