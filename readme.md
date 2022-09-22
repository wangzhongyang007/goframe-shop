# 项目启动失败的原因
1. Go或者goframe安装的版本不一致
2. 没有配置相关的mysql和redis

# 出现问题可以联系我

## 微信

wangzhongyang1993

## 我的博客

https://juejin.cn/user/2189882892232029/posts

## 我的公众号

程序员升级打怪之旅

# 交叉编译

## for Linux

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

## for windows

```
CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build
```

# 部署流程

1. 本地提交git
2. 因为远程服务器已经安装了go环境，直接执行部署脚本：

```
/home/shop/setup.sh
```

# 热更新启动项目

```
gf run main.go
```