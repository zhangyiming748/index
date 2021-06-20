---
layout:     post                    # 使用的布局(不需要改)
title:      SublimeText技巧        # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-06-09 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - sublime
---

# 设置代码格式化快捷键
`[Preferences]->[Key Bindings]->[User]`
```
{
 "keys": ["alt+super+l"],
 "command": "reindent", 
 "args": {
    "single_line": false
    }
}

```
# 设置Package走代理
`[Preferences]->[Package Settings]->[Package Control]->[Settings]`
```
"http_proxy":"http://127.0.0.1:8889",
"https_proxy":"http://127.0.0.1:8889",
```