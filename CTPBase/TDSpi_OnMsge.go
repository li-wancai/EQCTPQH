package CTPBase

func (s *TDSpiN) OnMessage(data interface{}, bIsLast bool) {
	log.Infof("收到信息%v", data)
}
