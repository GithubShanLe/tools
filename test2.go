package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type ListenFlow struct { //网卡接收和发送的流量总流量
	recv   float64
	sends  float64
	status string
}

var TempHardStat SystemStatus
var CurrentInfo, CurrentInfoBefor [10]ListenFlow //暂时使用的数据为10个为上限,不设为数组也可。

type NetCardStatus struct { //网卡信息
	Name         string
	Status       string
	BitrateRecev float64
	BitrateSends float64
}

type SystemStatus struct {
	NetCard []NetCardStatus //网卡信息
}

/*获取网卡的流量信息*/
func ListenFlowAndState(name string) ListenFlow {
	var temp ListenFlow
	f, err := os.Open("/proc/net/dev")
	if err != nil {
		return ListenFlow{}
	}
	defer f.Close()
	log.Println("}}}}}}}}}}}}}}}}1")
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return ListenFlow{}
		}
		if !strings.Contains(line, ":") {
			continue
		}
		ethname := strings.Split(line, ":")
		Name := strings.Fields(ethname[0])
		if Name[0] == "lo" {
			continue
		}
		if Name[0] == name {
			//Trim(&ethname[0])
			fields := strings.Fields(ethname[1])
			temp.recv, _ = strconv.ParseFloat(fields[0], 64)
			temp.sends, _ = strconv.ParseFloat(fields[8], 64)
			return temp
		}
	}
	return ListenFlow{}
}

// func Trim(str *string) {
// 	*str = strings.TrimSpace(*str)
// 	*str = strings.Replace(*str, "\n", "", -1)
// }
/*获取所有网卡的名称*/
func GetNetCardName() []net.Interface {

	var NCS NetCardStatus
	ncs, err := net.Interfaces()
	if err != nil {
		// return  nil,err
	}
	for k, v := range ncs {
		if v.Name == "lo" {
			ncs = append(ncs[:k], ncs[k+1:]...) //获取网卡名称，除了"lo"，表示切片删除
			break
		}

	}
	for i := 0; i < len(ncs); i++ {
		TempHardStat.NetCard = append(TempHardStat.NetCard, NCS) //追加一个空的NetCardStatus,之后赋值
	}
	return ncs //返回网卡名称
}

func main() {

	for {
		t := time.NewTimer(time.Second * 1)
		<-t.C

		TempHardStat.NetCard = append(TempHardStat.NetCard[:0], TempHardStat.NetCard[len(TempHardStat.NetCard):]...)
		NetCardName := GetNetCardName()
		log.Println("++++++", len(NetCardName), NetCardName)
		for k, netcard := range NetCardName {
			CurrentInfo[k] = ListenFlowAndState(netcard.Name) //接收收到网卡的基本信息
			log.Println(CurrentInfo[k].recv, CurrentInfo[k].sends)
			TempHardStat.NetCard[k].Name = netcard.Name
			TempHardStat.NetCard[k].BitrateRecev = (CurrentInfo[k].recv - CurrentInfoBefor[k].recv) * 8 / 1000
			TempHardStat.NetCard[k].BitrateSends = (CurrentInfo[k].sends - CurrentInfoBefor[k].sends) * 8 / 1000
			Flag := strings.Split(netcard.Flags.String(), "|")
			TempHardStat.NetCard[k].Status = Flag[0]
			fmt.Printf("网卡状态：%s\n", Flag[0])
			fmt.Printf("%s发送流量速度：%d B/s\n", netcard.Name, CurrentInfo[k].sends-CurrentInfoBefor[k].sends)
			fmt.Printf("%s接收流量速度：%d B/s\n\n", netcard.Name, CurrentInfo[k].recv-CurrentInfoBefor[k].recv)
			CurrentInfoBefor[k] = CurrentInfo[k] //作为上一次的数据保存

		}
	}

}
