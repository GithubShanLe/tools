package util

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

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

//@@ i interface{} 结构体的数据,用来写入文件的源数据
//@@ path 文件路径
//@@ indent 锁紧格式 "\t" | "\n" | "" | "    "| ....
//@@ 是否保留{port: 80} 中间的空格 true :保留 false:不保留 ，golang 的配置文件一般会存在空格
//@@function 写入配置文件
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
		str := strings.Replace(string(data), " ", "", -1)
		data = []byte(str)
	}
	f.Write(data)
	return nil
}
