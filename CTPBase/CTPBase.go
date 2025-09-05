package CTPBase

import "github.com/li-wancai/logger"

var log *logger.LogN

func SetLogger(l *logger.LogN) {
	log = l //配置log信息
}
