---
layout:     post                    # 使用的布局(不需要改)
title:     树莓派3B制作AndroidTV             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-05-02 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                   # 是否归档
tags:                               #标签
    - 技巧
---
>Android TV is a version of the Android operating system designed for digital media players, set-top boxes, soundbars, and TVs and developed by Google. Serving as a replacement for Google TV, it features a user interface designed around content discovery and voice search, surfacing content aggregated from various media apps and services, and integration with other recent Google technologies such as Assistant, Cast, and Knowledge Graph.

## 需要的工具
- [x] Raspberry 3B
- [x] Micro SD Card
- [x] power line
- [ ] Lan
- [x] Keyborad
- [x] mouse

## 需要的软件
- [x] Android image
- [x] Gapps
- [x] Recovery2Boot
- [x] SD Card Formatter
- [x] Etcher

## 操作步骤
1. 下载安装`lineage-14.1-20180817-UNOFFICIAL-KonstaKANG-rpi3.img`
2. 下载GApps文件包备用
3. 开机进入开发者模式允许root权限
4. 树莓派进入recovery模式
5. 刷入`open_gapps-arm-7.1-tvstock-20200502.zip`
6. 刷入`lineage-14.1-rpi3-recovery2boot.zip`
