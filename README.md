# iris-demo
> iris框架demo

## 开发及发布
### 工具
* golang(1.7+)

#### 如何开始
1. 下载并安装[Golang](https://golang.org/dl/)

#### 项目位置

项目路径应该在 本机go的src目录下 {gopath}/src/maizuo.com/back-end/iris-demo, 如果项目路径正确的路径,会导致不然访问第三方包资源

#### 启动模板前的准备

1.项目配置文件在config目录下, 本地启动项目使用的是config.json这个配置文件

2.在配置文件中, 修改server下的port和host, 确保项目运行在可用的端口上(如果出现其他人不能访问的情况,注意将server下的host修改为 0.0.0.0)

3.如果需要连接grpc服务端, 先要启动服务端项目,确保相关的服务端的port和host配置是正确的

4.如果不需要grpc相关业务, 可以在第一层的目录的main函数中, 注释掉initialize.SetupRPC() 避免启动报错(同理如果不需要redis也可以在此处注释掉)

5.本地运行项目 isDevelopment配置 确保为true, 输出日志会打印在控制台

#### 启动项目

- 启动项目
##### linux或mac系统下:
```
make go
```
##### win系统下安装工具支持gun命令

[MinGW](http://www.mingw.org/wiki/getting_started)

##### 如果安装失败,或者不想安装,也可以完整命令执行启动

```
go run main.go -conf ./config/local
```

#### 启动遇到缺包问题

解决方式
1.通过go get下载缺少的包(部分包需要翻墙) 

```
eg: go get http://github.com/xxxxxxx
```

2.下载不下来的包 通过golangtc等第三方网站下载

> http://golangtc.com/download/package

3.直接拷贝其他同事已经下载好的包


#### 测试项目是否正常启动链接
demo访问
> http://localhost:8080/api/demo/eno

grpc接口查询

> http://localhost:8080/api/user/200000666/address


#### 项目结构说明

```
|____iris-demo						项目根目录
| |____config       				配置文件目录
| | |____dev.json					测试环境配置   
| | |____local.json					本地环境配置
| | |____prod.json					正式环境配置
| |____Godeps						go项目打包描述目录
| | |____Godeps.json
| | |____Readme
| |____Makefile
| |____README.md
| |____src 							代码根目录
| | |____server 					后端项目根目录
| | | |____controller 				controller层 处理业务参数
| | | | |____demoController.go
| | | | |____main.go
| | | |____data 					data层 从其他系统获取数据 包括db,第三方接口等
| | | | |____dataInter 				data层接口
| | | | | |____demoDataInterface.go
| | | | |____demoData.go 			
| | | | |____test 					data层测试
| | | | | |____demoData_test.go
| | | | | |____main.go
| | | |____entity 					entity
| | | | |____result.go
| | | |____errcode 					全局错误集合
| | | | |____main.go
| | | |____initialize 				项目初始化层: 处理项目启动时需要初始化的内容:redis,db,log等
| | | | |____config.go              初始化加载项目配置
| | | | |____context.go             初始化项目上下文环境
| | | | |____error.go               全局panic异常拦截
| | | | |____logger.go              打开日志记录连接
| | | | |____redis.go               建立redis连接
| | | | |____rpc.go                 建立grpc连接
| | | | |____server.go              启动http服务
| | | |____proto 					protobuf协议生成的对象
| | | | |____protobuf
| | | | | |____user.pb.go
| | | |____route 					route层: 处理路由地址 分派到对应的controller方法
| | | | |____api 					api: 后端接口相关
| | | | | |____main.go
| | | | |____web 					web: 前端页面相关
| | | |____service 					service层: 处理业务逻辑
| | | | |____demoService.go
| | | | |____serviceInter 			service层接口
| | | | | |____demoServiceInterface.go
| | | | |____test 					service层测试
| | | |____test 					
| | | | |____demo_test.go
| | | | |____test_test.go
| | | |____timer 					定时器相关
| | | | |____healthCheck.go
| | | | |____timer.go
| | | | |____timerDemo.go
| | | |____util 					项目工具类
| | | | |____base64.go
| | | | |____http.go
| | | | |____ip.go
| | | | |____log.go
| | | | |____md5.go
| | | | |____rateLimiter.go
| | | | |____redis.go
| | | |____common 					项目常量包
```

