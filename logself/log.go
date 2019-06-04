package logself

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"sync"
)

var LogCenter LogQueue

type LogQueue struct {
	LoggerQueue map[string]*logs.BeeLogger
	Lock        *sync.RWMutex
}

type LogConf struct {
	Maxdays  int
	Filename string
}

func init() {
	LogCenter.LoggerQueue = make(map[string]*logs.BeeLogger)
	LogCenter.Lock = new(sync.RWMutex)
}

func initLog(path string) *logs.BeeLogger {
	logger := logs.NewLogger()
	log_config := LogConf{7, path}
	logconfig_json, _ := json.Marshal(log_config)
	logger.SetLogger(logs.AdapterFile, string(logconfig_json))
	logger.Async()
	return logger
}

func (l *LogQueue) WriteLog(path string, filename string, info string) {
	if path[len(path)-1] != '/' {
		path = path + "/"
	}
	l.Lock.RLock()
	logger, ok := l.LoggerQueue[path+filename]
	l.Lock.RUnlock()
	if !ok {
		logger = initLog(path)
		l.Lock.Lock()
		l.LoggerQueue[path] = logger
		l.Lock.Unlock()
	}
	//go beego.Info(info)
	go logger.Debug(info)
}
