package main

import (
	"encoding/json"
	"log"
	"network/util"
	"os"
	"strings"
)

//func main() {
//	timer1 := time.NewTimer(time.Second * 2)
//	t1 := time.Now()
//	fmt.Printf("t1:%v\n", t1)
//
//	t2 := <-timer1.C
//	fmt.Printf("t2:%v\n", t2)
//
//	//如果只是想单纯的等待的话，可以使用 time.Sleep 来实现
//	timer2 := time.NewTimer(time.Second * 2)
//	<-timer2.C
//	fmt.Println("2s后", time.Now())
//
//	time.Sleep(time.Second * 2)
//	fmt.Println("再一次2s后", time.Now())
//
//	<-time.After(time.Second * 2) //time.After函数的返回值是chan Time
//	fmt.Println("再再一次2s后", time.Now())
//
//	timer3 := time.NewTimer(time.Second)
//	go func() {
//		<-timer3.C
//		fmt.Println("Timer 3 expired", time.Now())
//	}()
//
//	//stop := timer3.Stop() //停止定时器
//	//////阻止timer事件发生，当该函数执行后，timer计时器停止，相应的事件不再执行
//	//if stop {
//	//	fmt.Println("Timer 3 stopped", time.Now())
//	//}
//
//	fmt.Println("before")
//	timer4 := time.NewTimer(time.Second * 5) //原来设置5s
//	timer4.Reset(time.Second * 1)            //重新设置时间,即修改NewTimer的时间
//	<-timer4.C
//	fmt.Println("after")
//}
var CronStruct Cronself

type Cronself struct {
	Hour    int
	Min     int
	Second  int
	Nsecond int
}

//func init() {
//	err := ReadFileToJson(CronStruct, "/Users/lee/application/go/src/network/conf/dynamic.conf")
//	if err != nil {
//		beego.Info("jjjjjjj")
//	}
//	fmt.Println("befor11", time.Now(), CronStruct.Hour, CronStruct.Min, CronStruct.Second, CronStruct.Nsecond)
//
//}
//func main() {
//	kk := new(Cronself)
//	kk.Hour = 10
//	err1 := WriteJsonToFile(kk, common.ConfigPath, "\t", true)
//	log.Println(err1, kk)
//	err := ReadFileToJson(&CronStruct, "/Users/lee/application/go/src/network/conf/dynamic.conf")
//	if err != nil {
//		logself.LogCenter.WriteLog(common.LogPath, "kkk.log", err.Error())
//	}
//	fmt.Println("befor11", time.Now(), CronStruct.Hour, CronStruct.Min, CronStruct.Second, CronStruct.Nsecond)
//	go func() {
//		for {
//			now := time.Now()
//			next := now.Add(time.Hour * 24)
//			//nexttime = time.Date(nexttime.Year(), nexttime.Month(), nexttime.Day(), 0, 0, 0, 0, nexttime.Location())
//			next = time.Date(next.Year(), next.Month(), next.Day(), CronStruct.Hour, CronStruct.Min, CronStruct.Second, CronStruct.Nsecond, next.Location())
//			timer4 := time.NewTimer(next.Sub(now))
//
//			<-timer4.C
//			//TASK()
//		}
//	}()
//	for {
//		//sflksdhfjk
//		go func() {}()
//	}
//
//}

//注意 i是局部变量， 它的形参类型应是指针对象，
func ReadFileToJson(i interface{}, path string) error {
	//
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	//defer f.Close()
	jsondec := json.NewDecoder(f)
	if err = jsondec.Decode(i); err != nil {

		return err
	}
	return nil
}

func WriteJsonToFile(i interface{}, path string, indent string, blank bool) error {
	os.Remove(path)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Println("Open file failed:", err)
		return err
	}
	defer f.Close()
	data, _ := json.MarshalIndent(i, "", indent)
	//主要是为了去除MashalIndent后对象和属性之间有空格
	//edg:port:_80  -->  port:80；'_'代表空格
	if !blank {
		str := strings.Replace(string(data), "", "", -1)
		data = []byte(str)
	}
	f.Write(data)
	return nil
}

func main() {

	kk := util.Converter("255.255.240.0")
	log.Println(kk)
}
