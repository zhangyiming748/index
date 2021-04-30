---
layout:     post                    # 使用的布局(不需要改)
title:      Git Clone 加速       # 标题
subtitle:   我是一只被禁足的安小鸟  #副标题
date:       2020-02-21   00:00:01 GMT+0800           # 时间
author:     Zen                      # 作者
header-img: img/photo/birdAngle.webp   #这篇文章标题背景图片
catalog: True                       # 是否归档
tags:                               #标签
    - Git
---

1. 查找域名对应的ip地址
`nslookup github.global.ssl.fastly.Net`
`nslookup github.com`
2. 修改hosts文件
`vi etc/hosts`
3. 重启或更新
`sudo /etc/init.d/networking restart`
`sudo dscacheutil -flushcache`
