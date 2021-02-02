---
layout:     post                    # 使用的布局(不需要改)
title:      通过修改实例路径,改变硬件信息             # 标题
subtitle:   娱乐大师常规用法 #副标题
date:       2021-01-30 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - 闲谈
---
1. 设备管理器打开显示适配器找到显卡右键属性找到设备实例路径,如图
[![yAKG4K.png](https://s3.ax1x.com/2021/01/31/yAKG4K.png)](https://imgchr.com/i/yAKG4K)
2. 打开注册表编辑器定位到
`计算机\HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Enum\PCI\VEN_1002&DEV_15D8&SUBSYS_3E0D19E5&REV_C1\4&35feb52a&0&0041`
3. DeviceDesc键即为设备显示名称,可以随意修改,之后再用鲁大师什么的看就是改好的名字
