package tce

import (
	"fmt"
	"tce-shell/conf"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
)

func Cdn_Update(path string) {
	//// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey
	credential := common.NewCredential(
		conf.AK,
		conf.SK,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := cdn.NewClient(credential, "", cpf)

	// 实例化一个请求对象
	request := cdn.NewPurgePathCacheRequest()
	request.Paths = common.StringPtrs([]string{path})
	request.FlushType = common.StringPtr("delete")

	// 返回的resp是一个PurgePathCacheResponse的实例，与请求对象对应
	response, err := client.PurgePathCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	//输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
