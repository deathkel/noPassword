## noPassword
一个密码生成器，用于生成随机字符密码，使用golang编写，可生成多平台二进制可执行程序
## 简单描述
* 这是一个简单的密码生成工具，目的在于为每个需要登录的应用设置自己的不同的密码，且不在任何地方存储自己的密码，每次需要密码时通过指定格式的方式生成密码。
* 程序会根据 执行程序后面的所有参数生成对应的随机字符串（随机字符串包括0~9,a~z,A~Z）
### 原理
应用名称 + 生成密码算法 + 任意密码 * N = 唯一密码
### 格式示例
* 应用名称 + 独立密码
* 应用名称 + 子应用名称 + 独立密码
* 其他任意格式
### 使用示例
```bash
./noPassword 网易 qweasdzxc

5cD2mT1kS7oA9yA1eN3sC0fL3fN5nV0a

### 指定密码长度

./noPassword -l 16 网易 qweasdzxc

5cL0aI0hL5tW6jM0
```

## 环境搭建
1. 依赖组件：
   golang依赖管理工具:[dep](http://cf.meitu.com/confluence/pages/viewpage.action?pageId=35854688)
3. 设置GOPATH工作目录
   sudo vim ~/.bash_profile
   export GOPATH="/www/go"
   source ~/.bash_profile
4. 安装golang依赖库
   `dep ensure`

#### 使用
1. 编译与启动程序
`cd main && go build noPassword.go`
`./noPassword 网易 qweasdzxc`


2 编译其他平台
[参考此处](https://blog.csdn.net/panshiqu/article/details/53788067)


### 在Alfred中使用
1. 将plugin\flred中的noPassword.alfredworkflow扩展加入alfred的workflow
2. 命令示例 `np 网易 123456`,按回车复制获得的密码 
![alfred-example](/resource/img/alfred-example.png)
