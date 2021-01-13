![hfunc](docs/images/hfunc.png)

# hfunc

hfunc 一套轻量级 Go 微服务框架，包含大量微服务相关框架及工具。  

## Goals

打造一款简单通用，集各家之所长的微服务框架

### Principles

* 简单：不过度设计，代码平实简单；


## Features
* Logger：标准日志接口，可方便集成三方 log 库；
* Web: 封装gin，开箱即用;
* Orm: 集成gorm, 开箱即用，也可以使用其他orm框架;
* Server：进行基础的 Server 层封装, 支持tcp,udp,http等协议的使用;

## Getting Started
### Required
- [go](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

### Install hfunc
```
# 安装生成工具
go get github.com/hnit-acm/hfunc/tool/hfunc
```
### Create a service
```
# 创建项目模板
hfunc new helloworld
```

## Community
* QQ Group: 

## License
