package main

import (
	"log"
	"time"
)

type TaskFunc func() error

func BackUpTask(f TaskFunc, cron ...int) {
	back_timer := time.NewTicker(time.Second)
	for {
		//now := time.Now()
		//next := now.Add(time.Hour * 24)
		////next = time.Date(next.Year(), next.Month(), next.Day(), cron[0], cron[1], cron[2], cron[3], next.Location())
		//next = time.Date(next.Year(), next.Month(), next.Day(), CronStruct.Hour, CronStruct.Min, CronStruct.Second, CronStruct.Nsecond, next.Location())

		<-back_timer.C
		//backup()
		f()
	}
}
func main() {

	BackUpTask(func() error {
		log.Println("opopop")
		return nil
	}, 1, 2)
}
