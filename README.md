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
`cd main && go build`

1.1 说明: 程序会根据 执行程序后面的所有参数生成对应的随机数

* 最简单的使用方式 网站名称 + 一个密码(类似onePassword)
```bash
./noPassword 网易 qweasdzxc

5cD2mT1kS7oA9yA1eN3sC0fL3fN5nV0a
```

* 指定密码长度
```bash
./noPassword -l 16 网易 qweasdzxc

5cL0aI0hL5tW6jM0
```

* 编译其他平台
[参考此处](https://blog.csdn.net/panshiqu/article/details/53788067)