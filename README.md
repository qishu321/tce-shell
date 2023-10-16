<p align="center">
  <a href="https://github.com/qishu321/tce-shell">
    <img src="https://avatars.githubusercontent.com/u/95009146?s=400&u=0984e6a6a761fa007f6ad459abbb1ee9786424b8&v=4" alt="Logo" width="180" height="180">
  </a>

  <h1 align="center">tce-shell</h1>
  <p align="center">
   本项目使用go开发，实现一些常用调用腾讯云api操作。
    <br />
     <br />
  </p>
## 技术栈

#### 后端 Golang 1.19

#### 已实现的功能
- cdn更新
- clb安全组更新
- redis备份

```bash
##
C:\Users\w\go\src\tce-shell>go run main.go  --help
Usage of C:\Users\w\AppData\Local\Temp\go-build3214296154\b001\exe\main.exe:
  -cdn_update string
        cdn_update list CdnUrl1 CdnUrl2…… (default "cdn")
  -clb_acl string
        clb_acl list 负载均衡1ID;安全组id1,安全组id2 负载均衡2ID;安全组id1,安全组id2 (default "clb")
  -redis_backup string
        redis_backup list 实例1 实例2 实例3…… (default "redis")

```

## 部署方法

###安装编译

```go
# clone
git clone https://github.com/qishu321/tce-shell.git
修改conf里config.ini里的ak和sk
然后可以go run main.go -clb_acl --help
也可以./tce-shell -clb_acl --help

```


#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。