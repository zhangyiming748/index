---
layout:     post                    # 使用的布局(不需要改)
title:      raspi蓝牙控制       # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-06-28 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: True                     # 是否归档
tags:                               #标签
    - Raspberry
---

```
bluetoothctl # 启动蓝牙程序
power on # 输入power on 命令打开控制器电源。默认是关闭的
devices # 获取要配对设备的 MAC 地址
scan on # 如果设备未在清单中列出，输入 scan on 命令设置设备发现模式
agent on # 命令打开代理
pair 00:90:4B:20:62:99 # 开始配对（支持 tab 键补全）
trust 00:90:4B:20:62:99 # 如果使用无 PIN 码设备，再次连接可能需要手工认证
connect 00:90:4B:20:62:99 # 建立连接
exit # 退出bluetoothct
```
