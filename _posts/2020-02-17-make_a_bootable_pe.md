---
layout:     post                    # 使用的布局(不需要改)
title:      用Rufus不费脑制作PE              # 标题
subtitle:   我是一只被禁足的安小鸟  #副标题
date:       2020-02-17    00:00:01 GMT+0800          # 时间
author:     Zen                      # 作者
header-img: img/photo/birdAngle.webp   #这篇文章标题背景图片
catalog: True                       # 是否归档
tags:                               #标签
    - Windows
---
首先说明PE只是用来维护系统而不是重装系统的工具,不多说,你细品
----
原料:
  - Rufus
  - 根据个人喜好选择的PE镜像
  - U盘(容量由镜像大小决定)

----

# 天才第一步

插入U盘,打开Rufus,选择PE,选择MBR或GPT,制作,打完收工

# 配图

![打开Rufus](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/openRufus.png?raw=true)<center>打开Rufus</center>
![选择这个U盘](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/selectDevices.png?raw=true)<center>选择这个U盘</center>
![选择已经下载好的PE镜像](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/selectISO.png?raw=true)<center>选择已经下载好的PE镜像</center>
![选择MBR/GPT](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/selectPartTabel.png?raw=true)<center>选择MBR/GPT</center>
![不要更改默认设置](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/default.png?raw=true)<center>不要更改默认设置</center>
![开始制作](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/start.png?raw=true)<center>开始制作</center>
![最后一次确认](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/makeSure.png?raw=true)<center>最后一次确认</center>
![制作完成](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/finish.png?raw=true)<center>制作完成</center>
关闭电脑后,打开电源长按F12(部分新款Thinkpad需要开机先按回车之后再按F12,部分华硕机型热键是DEL),选择这个U盘,enter进入
![启动PE](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/startPE.png?raw=true)<center>启动PE</center>
![进入PE](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/Make_Bootable_PE/intoPE.png?raw=true)<center>进入PE</center>
然后就像使用正常电脑一样,把电脑里重要的文件复制出来,粘贴到外部存储器上(移动硬盘什么的),然后就可以重装这台电脑了
可以参考
[系统盘安装原版Win10](https://zhangyiming748.github.io/2019/05/19/install_win10/)
和
[制作系统盘](https://zhangyiming748.github.io/2019/05/16/make_a_bootable_usb_disk/)
# 途中使用的软件和镜像下载
首先说明软件与本人无关
>根据2002年1月1日<计算机软件保护条例>规定:为了学习和研究软件内含的设计思想和原理,通过安装,显示,传输或者存储软件等方式使用软件的,可以不经软件著作权人许可,不向其支付报酬!

光卡PE:[b29p](https://pan.baidu.com/s/1PrTQ4aqhAVYllVexk7TDjg)

Rufus:[wz46](https://pan.baidu.com/s/1C8Dq_GMeCith9qNTptOcaQ)
