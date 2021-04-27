---
layout:     post                    # 使用的布局(不需要改)
title:      SSH 三步解决免密登录             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-04-27 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - linux
---
# 客户端生成公私钥
本地客户端生成公私钥:(一路回车默认即可)
```
$ ssh-keygen
//在用户目录.ssh文件夹下创建公私钥
```
# 上传公钥到服务器
这里测试用的服务器地址为:192.168.1.5
用户为:pi
 ```
 $ ssh-copy-id -i ~/.ssh/id_rsa.pub pi@192.168.1.5
 //将公钥写到服务器上的ssh目录下
```
# 测试免密登录
客户端通过ssh连接远程服务器,就可以免密登录了.
```
$ ssh pi@192.168.1.5
```
