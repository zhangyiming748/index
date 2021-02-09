---
layout:     post                    # 使用的布局(不需要改)
title:     移除华为云盘在此电脑上的顽固图标            # 标题
subtitle:   看看华为云盘是先解决问题还是先解决提出问题人 #副标题
date:       2021-02-09  00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                  # 是否归档
tags:                               #标签
    - 华为
---

>--我想提issue<br>
--你需要先安装windows1803以上的版本<br>
--安装了<br>
--你要安装华为云盘10.3版本才有资格提出反馈<br>
--哪里能安装10.3版本？<br>
--不好意思，我们还没发布<br>

这种推脱话术不高级，直接说必须在Windows11上运行才有权提出反馈多好

华为云服务客户端虽说还在Beta测试,我在那个群里,提过很多反馈,但到如今,该改的还是没改,今天重新安装竟又发现了问题

安装后在此电脑中的图标不能通过常规手段删除,就算某网盘都能在设置中取消在此电脑中显示图标入口,华为竟然到如今还不支持,不知道是技术不支持还是人为不想支持,既然如此,那就只能先君子后小人了

[![安装前此电脑的状态](https://s3.ax1x.com/2021/02/09/yaWqeI.png)](https://imgchr.com/i/yaWqeI)
<center>安装前此电脑的状态</center>

[![安装后初次登陆前](https://s3.ax1x.com/2021/02/09/yaWHOA.png)](https://imgchr.com/i/yaWHOA)
<center>安装后初次登陆前</center>

[![打开华为云盘同步功能后](https://s3.ax1x.com/2021/02/09/yafplQ.png)](https://imgchr.com/i/yafplQ)
<center>打开华为云盘同步功能后</center>

<h1>从这里开始,华为变味了</h1>

[![此电脑上华为云盘的注册表位置](https://s3.ax1x.com/2021/02/09/yaW7yd.png)](https://imgchr.com/i/yaW7yd)
<center>此电脑上华为云盘的注册表位置</center>


`计算机\HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\{$CLSID}`
这里CLSID在云盘每次启动的时候都会变,防止自己被干掉

删除之后刷新一下此电脑

[![此电脑删除后](https://s3.ax1x.com/2021/02/09/yaWoSe.png)](https://imgchr.com/i/yaWoSe)
----
不过以上其实都是无用功
华为云盘项目组不改,用户做什么都只是饮鸩止渴
如果你kill进程后重新打开
那么一切都会恢复原来的样子
[![重新登陆后图标回归](https://s3.ax1x.com/2021/02/09/yafSSg.png)](https://imgchr.com/i/yafSSg)
<center>风里雨里,具有流氓行为的软件图标一直在这里等你</center>
