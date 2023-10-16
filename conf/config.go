package conf

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	AK      string
	SK     string
)

func init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败，请检查文件路径:", err)
	}
	LoadApp(file)
}
func LoadApp(file *ini.File) {
	AK = file.Section("app").Key("AK").String()
	SK = file.Section("app").Key("SK").String()
}

