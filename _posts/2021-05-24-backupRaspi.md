---
layout:     post                    # 使用的布局(不需要改)
title:      树莓派系统整体备份           # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-05-24 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Raspberry
---

看到的一个youtuber的视频介绍的工具挺不错,正好手头手头一个树莓派设备身兼数职:作为Linux平台廉价的硬件跑高负载任务不心疼,作为网络机顶盒刷kodi看电视直播,作为openwrt科学上网,作为打印机服务器等等,每次转变一个角色就要重刷系统重新安装软件确实很麻烦,做好几个备份每次只要恢复就直接处于就绪状态是个省时间的好主意

----

<iframe width="560" height="315" src="https://www.youtube.com/embed/k61s3_7nyS0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>


# 以下是我的操作

我的磁盘状态,需要把sda备份到sdb
```
$ df -h
Filesystem      Size  Used Avail Use% Mounted on
/dev/root       117G  8.5G  103G   8% /
devtmpfs        1.7G     0  1.7G   0% /dev
tmpfs           1.8G     0  1.8G   0% /dev/shm
tmpfs           1.8G   10M  1.8G   1% /run
tmpfs           5.0M  4.0K  5.0M   1% /run/lock
tmpfs           1.8G     0  1.8G   0% /sys/fs/cgroup
/dev/sda1       253M   49M  204M  20% /boot
tmpfs           361M  8.0K  361M   1% /run/user/1000
/dev/sdb1       670G  435G  235G  65% /media/pi/Gloway720
```

执行备份命令

`$ sudo dc3dd if=/dev/sda of=/media/pi/Gloway720/backup.img`

从固态硬盘(NVME)复制到固态硬盘(SATA)

均插在USB3.0Port上,鼠标键盘均已拔出,使用网线连接,操作使用远程ssh连接

备份过程中任何操作都会有几率让硬盘断电报错

报错信息

```
writing to `/media/pi/Gloway720/backup.img': Input/output error
```
使用带辅助供电的USB线缆重新接驳状况改善,但依然不能在备份过程中对目标磁盘进行IO操作,否则依然报错

备份会生成与硬盘相同大小的IMG文件,我一开始还天真的

`$ sudo dc3dd if=/dev/sda of=/home/pi/backup.img`
