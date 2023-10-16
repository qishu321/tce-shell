package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"tce-shell/tce"
)

var (
	redisBackup string
	cdnUpdate   string
    ClbAcl   string
)

func main() {
	flag.StringVar(&redisBackup, "redis_backup", "redis", "redis_backup list 实例1 实例2 实例3……")
	flag.StringVar(&cdnUpdate, "cdn_update", "cdn", "cdn_update list CdnUrl1 CdnUrl2……")
	flag.StringVar(&ClbAcl, "clb_acl", "clb", "clb_acl list 负载均衡1ID;安全组id1,安全组id2 负载均衡2ID;安全组id1,安全组id2")


	flag.Parse()
	if len(os.Args) > 0 {
		args := os.Args
		if args[2] != "list"{
			fmt.Println("错误：-参数后面的第一个参数必须是 'list'")
			os.Exit(1)
		}
	}
	if ClbAcl != "clb" {
		for _, arg := range flag.Args() {
			parameters := strings.Split(arg, ";")
			var var1 string
			var1 = parameters[0]
			aclids := strings.Split(parameters[1], ",") // 拆分安全组id
			tce.Clb_acl(var1,aclids)
		}
	}
	if redisBackup != "redis" {
		for _, arg := range flag.Args() {
			tce.Redis_backup(arg)
		}
	}
	if cdnUpdate != "cdn" {
		for _, arg := range flag.Args() {
			tce.Cdn_Update(arg)
		}
	}

}
