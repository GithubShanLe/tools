package util

import (
	"bufio"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type ListenFlow struct { //网卡接收和发送的流量总流量
	recv   float64
	sends  float64
	status string
}

var TempHardStat SystemStatus

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
		if ethname[0] == name {
			// Trim(&ethname[0])
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
