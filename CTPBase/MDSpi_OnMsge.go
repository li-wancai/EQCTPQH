package CTPBase

func (s *MDSpiN) OnMessage(data interface{}, bIsLast bool) {
	log.Infof("收到信息%v", data)
}
