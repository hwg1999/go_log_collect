package main

import (
	"logCollect/logTransfer/conf"
	"logCollect/logTransfer/es"
	"logCollect/logTransfer/logger"
	"path"
	"time"

	"logCollect/logTransfer/kafka"

	"gopkg.in/ini.v1"
)

var cfg = new(conf.AppConf)

func main() {
	// 1.加载配置文件
	err := ini.MapTo(cfg, "conf/config.ini")
	if err != nil {
		panic(err)
	}

	// 加载日志配置
	err = logger.Init(path.Join(cfg.LogConf.FilePath, cfg.LogConf.FileName), cfg.LogConf.LogLevel, time.Duration(cfg.LogConf.MaxAge)*time.Hour*24)
	if err != nil {
		logger.Log.Warnf("初始化日志文件失败, err:%v\n", err)
	}

	// 2.初始化es连接
	err = es.Init(cfg.EsConf.Address, cfg.EsConf.MaxChanSize, cfg.EsConf.Nums)
	if err != nil {
		logger.Log.Errorf("init ES client failed,err:%v\n", err)
		return
	}
	logger.Log.Debug("init es success.")

	// 3.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		logger.Log.Errorf("init kafka failed, err:%v\n", err)
		return
	}
	logger.Log.Debug("init kafka success")
}
