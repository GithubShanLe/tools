package data

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"os"
	"time"
)

var CronStruct Cronself

type TaskFunc func() error

type Cronself struct {
	Hour    int
	Min     int
	Second  int
	Nsecond int
}

func init() {
	err := ReadFileToJson(CronStruct, "../conf/dynamic.conf")
	if err != nil {
		beego.Info("jjjjjjj")
	}
}

func BackUpTask(f TaskFunc, cron ...int) {

	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), CronStruct.Hour, CronStruct.Min, CronStruct.Second, CronStruct.Nsecond, next.Location())
		back_timer := time.NewTicker(next.Sub(now))
		defer func() {
			back_timer.Stop()
			if err := recover(); err != nil {
				beego.Error("error:", err)
			}
		}()

		<-back_timer.C
		//backup()
		f()
	}
}

func ReadFileToJson(i interface{}, path string) error {
	f, err := os.Open(path)
	if err != nil {

		return err
	}
	defer f.Close()
	jsondec := json.NewDecoder(f)
	if err = jsondec.Decode(i); err != nil {

		return err
	}
	return nil
}
