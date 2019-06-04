package main

import (
	"github.com/astaxie/beego"
	"log"
	_ "network/conf"
	"network/data"
	_ "network/models"
	_ "network/routers"
)

func main() {
	//user := models.User{15, "15", "Lee", 25, true, "17621657350@163.com", "*****"}
	//data, _ := json.Marshal(user)
	//body, _ := util.HttpMethod("GET", "http://192.168.3.217:8081/sl/v1/users", nil)
	//fmt.Println(string(body), []byte(body))
	//body1, _ := util.HttpMethod("POST", "http://192.168.3.217:8081/sl/v1/users", &data)
	//fmt.Println(string(body1), []byte(body1))
	data.BackUpTask(17, 0, 0, 0)
	log.Println("JJJJJ")
	beego.Run() //}

}
