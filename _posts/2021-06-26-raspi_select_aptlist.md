---
layout:     post                    # 使用的布局(不需要改)
title:      raspi4B更换源        # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-06-26 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: True                     # 是否归档
tags:                               #标签
    - Raspberry
---


# 简介
2020年上半年,Raspberry Pi 4 发布了 8GB 内存版本,同时带来 64 位镜像,相较于32位镜像,在性能方面有很大提升.

操作系统镜像方面,现在 Raspberry Pi 支持 64 位,其包含与常规 32 位镜像相同的应用程序集和相同的桌面环境,不过只是针对 Debian arm64 构建的.同时,现在开始,32 位和 64 位操作系统镜像都不再使用"Raspbian"这一名称,改用新的统一的名称:Raspberry Pi OS,这样更便于用户识记.

树莓派镜像烧录不作过多介绍,本博文的Raspbian系统对应Debian10(buster),镜像站选用清华大学开源软件镜像站(tuna).

# 64位系统配置(arm64版本)
64位镜像可以直接使用debian的系统源,首先需要编辑`/etc/apt/sources.list`使用#号注释原内容,然后在末尾添加Debian的系统源:

```
# 默认注释了源码镜像以提高 apt update 速度,如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ buster main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ buster main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-updates main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-updates main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-backports main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ buster-backports main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian-security buster/updates main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security buster/updates main contrib non-free
```
接着需要配置树莓派官方的软件源,编辑/etc/apt/sources.list.d/raspi.list,同样的,注释原内容,在末尾添加如下内容:
`deb http://mirrors.tuna.tsinghua.edu.cn/raspberrypi/ buster main ui`

# 32位系统配置(armhf版本)
两者配置方法一致,都是修改上述两个文件(64位系统也可以下述系统源),具体如下:

编辑`/etc/apt/sources.list`注释原文件内容,用以下内容取代:
```
deb http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ buster main non-free contrib rpi
deb-src http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ buster main non-free contrib rpi
```
编辑`/etc/apt/sources.list.d/raspi.list`注释原文件内容,用以下内容取代:
`deb http://mirrors.tuna.tsinghua.edu.cn/raspberrypi/ buster main ui`
# 系统更新
系统源和软件源配置完成以后,使用下述两条命令更新:
```
sudo apt update
sudo apt upgrade -y

```
