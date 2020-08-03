---
layout:     post                    # 使用的布局(不需要改)
title:     在linux上制作windows启动盘            # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-07-10  00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: ture                  # 是否归档
tags:                               #标签
    - 操作系统
---

经常有初学者开始尝试双系统的时候在安装linux的步骤时直接选择liunx默认双系统方案--覆盖windows的grud,这样会导致windows系统在某些错误的时候不会出现在系统选单,最简单的方法就是重新安装windows,但是大部分人还不知道怎么在linux上制作系统盘,其实,在linux下制作系统盘比在windows上容易的不是一点半点,接下来就讲解一下两种系统启动方式的制作方法:
----

## UEFI
现在大部分的电脑都是支持UEFI启动甚至只支持UEFI启动的
----

1. 下载GParted分区编辑器

`sudo apt install gparted`
![UKGOgg.png](https://s1.ax1x.com/2020/07/10/UKGOgg.png)

2. 插入u盘,在GPatred中找到

![UKGL8S.png](https://s1.ax1x.com/2020/07/10/UKGL8S.png)
确保当前的设备未被挂载
![UKGvuj.png](https://s1.ax1x.com/2020/07/10/UKGvuj.png)

3. 选择菜单栏->设备->创建分区表

格式选择gpt

![UKGzbn.png](https://s1.ax1x.com/2020/07/10/UKGzbn.png)

4. 创建新分区,格式选择fat32
![UKGXvQ.png](https://s1.ax1x.com/2020/07/10/UKGXvQ.png)

5. 挂载下载好的iso,打开
![UKJpEq.png](https://s1.ax1x.com/2020/07/10/UKJpEq.png)
![]()

6. 复制全部文件到u盘上
![UKGxDs.png](https://s1.ax1x.com/2020/07/10/UKGxDs.png)

7. 制作完毕

以上步骤看起来繁琐,主要是给一些没有自主思考能力的人看的,其实这些步骤可以总结一句话:
> 解压windows的安装镜像到gpt分区表\fat32分区的u盘上

## BIOS(Legacy)
比较老的电脑会是只支持BIOS启动,包括在主板中能看到`EFI shell`字样的
----

1. 下载制作工具

`sudo add-apt-repository ppa:nilarimogard/webupd8`

`sudo apt update`

`sudo apt install woeusb`

2. 相信自己的语言理解能力