# 本地安装goctl
1.本地将go-zero项目克隆下来：  `git clone git@github.com:i4de/go-zero.git`
2.到目录go-zero\tools\goctl 下 执行命令： `go install`  
3.后续执行下面的各种goctl命令即可

# 环境初始化

protoc/protoc-gen-go/protoc-gen-grpc-go 依赖可以通过
`goctl env check -i -f` 一键安装

# api文件编译方法
命令执行路径: ithings\src\papisvr\
```shell script
goctl api go -api http/papi.api  -dir ./  --style=goZero
```
