---
layout:     post                    # 使用的布局(不需要改)
title:      证书方式激活WindowsThinPC             # 标题
subtitle:    #副标题
date:       2019-11-01 00:00:01 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/jiuzhaigou.webp    #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 操作系统
---

1. 确保三个文件放在`c:\`
![文件图](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/ActiveWindowsThinPC/activeWindowsThinPC.png)
2. 使用管理员身份运行cmd
3. 输入命令`slmgr.vbs /ilc c:\pkeyconfig-embedded.xrm-ms`出现对话框后点`OK`
4. 输入命令`slmgr.vbs /ilc c:\Security-SPP-Component-SKU-Embedded-VLBA-ul.xrm-ms`出现对话框后点`OK`
5. 输入命令`slmgr.vbs /ilc c:\Security-SPP-Component-SKU-Embedded-VLBA-ul-oob.xrm-ms`出现对话框后点`OK`
6. 输入命令`slmgr.vbs /ipk XGY72-BRBBT-FF8MH-2GG8H-W7KCW`出现对话框后点`OK`
7. 输入命令`slmgr.vbs /xpr`出现对话框后点`OK`激活成功

----
打完收工
