# 我自己项目所上传的站点位置：http://r3mix4lles.cn:9999/ (暂时不开放)
本站是从@chengka 大佬的[Singo](https://github.com/bydmm/singo)框架的学习项目所fork下来进行修改的。
这个项目为自己的毕设所开发，是系统的后端部分。
## 目前相对原项目所做的修改
+ 用户验证方式从session改为了token，使用的是[jwt-go](https://github.com/dgrijalva/jwt-go)。降低服务器负担，更加适应REST API的设计风格
+ 使用[file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs)来进行接口请求、响应记录，设置为输出到文件
+ 对用户增添了权限等级
+ 添加审核视频以及其它管理员后台相关的服务内容
+ 为视频添加了
    + 审核状态
    + 分区
+ 添加音频媒体
+ 新建审核日志模型，用来记录审核情况
+ 添加了评论功能
+ 添加了个人信息页
+ 上传接口可以根据上传文件的类型来确定上传到OSS的位置
+ 为视频审核调用生成雪碧图的接口
+ 为跨域设置了环境变量CORS
### 下面是原项目的说明内容
***

# G站: https://www.gourouting.com

欢迎来到[G站](www.gourouting.com)，本站是[Singo](https://github.com/bydmm/singo)框架的学习项目。

## 项目地址

https://github.com/bydmm/giligili

## 项目目的

本项目代码并不是为了真正经营一个视频站项目而编写。

本项目的主要目的是为了方便大家学习怎么用Golang编写前后端分离的纯后端项目

## 重要: 如何运行

#### 1.学习Go Module管理依赖

本项目已经迁移到使用Go Module来管理依赖,和视频的开始有所不同! 所以按照视频的方法是跑不起来的。

请参考本视频了解什么是Go Module：https://www.bilibili.com/video/av63052644/

Go Module会让你未来面对各种依赖问题迎刃而解，所以学习和使用对你是非常有价值的

#### 2.配置数据库

本项目依赖于任何网站项目都会使用的Mysql和Redis，所以你需要提前安装和启动这两个服务。

如果你是windows用户，可以快速的解决mysql和redis安装的问题,通过: PHPStudy。

本视频用几分钟教会你如何使用PHPStudy，https://www.bilibili.com/video/av64485001/

如果你是OSX或者linux的硬核用户，相必启动Mysql和Redis对你不是问题😁

#### 3.配置环境变量

> 设置环境变量，你可以参考singo框架的文档: https://singo.gourouting.com/quick-guide/set-env.html

由于每个用户的电脑环境不同，所以我们通过环境变量来改变着些容易变化的属性。

你需要复制项目根目录下的.env.example文件，然后建立.env文件，然后把内容帖进去

```ini
MYSQL_DSN="user:password@tcp(ip:port)/dbname?charset=utf8&parseTime=True&loc=Local" # mysql连接串
REDIS_ADDR="127.0.0.1:6379" # redis地址
REDIS_PW="" # redis密码(可以不填)
REDIS_DB="" # redis数据库(可以不填)
SESSION_SECRET="youneedtoset" # session密钥，开发环境可以不用改
GIN_MODE="debug" # 服务状态，开发环境不用改
# 下面是OSS对象存储的参数
# 参考本视频来管理上传文件：https://www.bilibili.com/video/av60189734/
OSS_END_POINT="oss-cn-hongkong.aliyuncs.com" # OSS端点
OSS_ACCESS_KEY_ID="xxx"
OSS_ACCESS_KEY_SECRET="qqqq"
OSS_BUCKET="lalalal"

```

#### Windows CMD 系统启动指令

```bash
set GOPROXY=https://mirrors.aliyun.com/goproxy/
set GO111MODULE=on

go run main.go
```

#### Windows Powershell 系统启动指令

```bash
$env:GOPROXY = 'https://mirrors.aliyun.com/goproxy/'
$env:GO111MODULE = 'on'

go run main.go
```

#### linux / OSX 系统启动

```bash
export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GO111MODULE=on

go run main.go
```

## 视频实况系列教程

[让我们写个G站吧！Golang全栈编程实况](https://space.bilibili.com/10/channel/detail?cid=78794)

## Singo框架

使用Singo开发Web服务，用最简单的架构，实现够用的框架，服务海量用户

https://github.com/bydmm/singo

## 神奇的接口文档

服务启动后: http://localhost:3000/swagger/index.html

接口文档位于项目swagger目录下。请阅读目录内的文档