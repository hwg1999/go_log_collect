package taillog

import "logAgent/etcd"

var taskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry: logEntryConf, // 把当前的日志收集项配置信息保存起来
	}

	for _, logEntry := range logEntryConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
