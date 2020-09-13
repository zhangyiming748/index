---
layout:     post                    # 使用的布局(不需要改)
title:      Mojave原生读写NTFS             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-09-11 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - macOS
---

+ 插上磁盘
![wNKJVe.png](https://s1.ax1x.com/2020/09/11/wNKJVe.png "桌面上会出现盘符,但这样打开只能读不能写")
+ 打开终端
![wNM1Wn.png](https://s1.ax1x.com/2020/09/11/wNM1Wn.png "查看磁盘的 Name")
<center>如图Elements就是我磁盘的Name</center>
+ 然后打开vim
![wNQ2j0.png](https://s1.ax1x.com/2020/09/11/wNQ2j0.png "sudo vim /etc/fstab")
+ 写入指令
![wNleUg.png](https://s1.ax1x.com/2020/09/11/wNleUg.png "LABEL=Elements	none	ntfs	rw,auto,nobrowse")
+ 保存文本后重启电脑
这时候会发现桌面和finder里面都没有磁盘
![wN8QxA.png](https://s1.ax1x.com/2020/09/11/wN8QxA.png "磁盘是挂在/Volumes下")
+ 为了方便使用在桌面创建快捷方式
![wNGpeP.png](https://s1.ax1x.com/2020/09/11/wNGpeP.png "在桌面创建快捷方式")
+ 卸载磁盘
![wNG8SJ.png](https://s1.ax1x.com/2020/09/11/wNG8SJ.png "或者直接用命令卸载")
+ 缺点
磁盘的省电功能不可用,通电过程中不会休眠,即使长时间未读写
