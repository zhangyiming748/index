---
layout:     post                    # 使用的布局(不需要改)
title:      Sublime设置ss代理      # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-07-20 00:00:07 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Proxy
---

Sublime设置http/https代理
Preference -> Package Setting -> Package Control -> Setting-User添加以下设置
```
"http_proxy": "socks5://127.0.0.1:1080",
"https_proxy": "socks5://127.0.0.1:1080",
```
