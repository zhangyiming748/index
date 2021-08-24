---
layout:     post                    # 使用的布局(不需要改)
title:      Atom设置代理    # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-08-24 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: True                     # 是否归档
tags:                               #标签
    - Proxy
---

# Atom 设置和取消代理
### 设置Shadowsocks代理
`apm config set http-proxy socks5:127.0.0.1:1089`
`apm config set https-proxy socks5:127.0.0.1:1089`
### 取消ssl
`apm config set strict-ssl false`
### 取消代理
`apm config set http-proxy null`
`apm config set https-proxy null`
### 查看代理设置
`apm config get http-proxy`
`apm config get https-proxy`
