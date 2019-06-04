package main

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"log"
	"time"
)

var (
	channel1, channel2 chan int
)

func init() {
	channel1 = make(chan int)
	channel2 = make(chan int)
}

func printinfo(s string) {
	for _, v := range s {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 300)
	}
}

func hh() {

	a := <-channel2
	switch {
	case a == 1:
		log.Println("OPPOP")
	case a == 2:
		log.Println("klklkl")
	}
}

func df() error {

	log.Println("df")
	return nil
}

func test1() {

	//printinfo("hello")
	//channel1 <- 8
	channel2 <- 1
}

func test2() {
	//<-channel1
	//printinfo("world")
	channel2 <- 3
}
func main() {

	//for {
	//	go test1()
	//	go test2()
	//	go hh()
	//}
	go func() {
		tk1 := toolbox.NewTask("mytask", "0/2 * * * * *", func() error {
			log.Println("JJJJJJJ")
			return nil
		})
		toolbox.AddTask("mytask", tk1)
		toolbox.StartTask() //真真切切定时执行
	}()

	for {

	}
}
