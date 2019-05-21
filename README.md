# 安装方法
## 1.下载项目
```bash
git clone https://git.imooc.com/coding-339/chat.git
```


## 2.项目配置
### 2.1 ide配置
建议采用goland IDE

ADD Configuration->左上角+->go build类型,右侧填写
```
files :{你的项目路径}/hellox.x/main.go
workdir:{你的项目路径}/hellox.x/
```
### 2.2 资源配置
将asset目录复制到对应项目下
如将asset复制到hellox.x下
### 2.3 数据库配置
修改service/init.go 中数据库配置文件
```cgo
drivename :="mysql"
DsName := "root:root@(192.168.0.102:3306)/chat?charset=utf8"
```
为你自己的数据库以及密码,格式如下
```
用户名:密码@(ip:port)/数据库名称?charset=utf8
```
### 2.4 页面入口地址
```
http://127.0.0.1:8080/user/login.shtml
```

## 3.依赖包安装

本项目依赖的`x/net`,`x/time`包可能会被墙，使用如下指令可以安装

```bash
$mkdir -p $GOPATH/src/golang.org/x/
$cd $GOPATH/src/golang.org/x/
$git clone https://github.com/golang/net.git net
$git clone https://github.com/golang/time.git time

```

其他包依赖

```bash
go get github.com/go-xorm/xorm
go get github.com/gorilla/websocket
go get gopkg.in/fatih/set.v0
go get github.com/go-xorm/xorm
```
# 联系作者
如有疑问请联系作者
```
关注公众号非正式解决方案  `betaidea`  后
即可获得作者微信号以及最新QQ群
提供答疑服务
提供其他附加服务如各种golang资料
各种系统源码
```
