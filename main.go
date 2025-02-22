package main

import (
	"fmt"
	"logAgent/conf"
	"logAgent/etcd"
	"logAgent/kafka"
	"logAgent/taillog"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
	wg  sync.WaitGroup
)

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "conf/config.ini")
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		return
	}

	// 1.初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed ,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")

	// 2. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")

	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("get conf from etcd failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	// 3. 收集日志发往Kafka
	taillog.Init(logEntryConf)
	// 因为NewConfChan访问了tskMgr的newConfChan, 这个channel是在taillog.Init(logEntryConf) 执行的初始化
	newConfChan := taillog.NewConfChan()
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan)
	wg.Wait()
}
