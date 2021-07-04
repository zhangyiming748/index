---
layout:     post                    # 使用的布局(不需要改)
title:      解决windows预览体验计划页面空白问题      # 标题
subtitle:   垂死病中惊坐起 笑问win11何时来 #副标题
date:       2021-07-04 00:01:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Windows
---

早上醒来,惊闻微软悄没声的发布了windows11,并且将于下周推送Insider预览版本

这次官方提供的把win本变成Mac的机会,我怎么会错过呢

泡好咖啡,速度打开  设置->Windows预览体验计划

卧槽?怎么是空白的啥也没有呢?

兄弟莫慌,在左下角搜索中搜索powershell或者直接从开始菜单中找到windows powershell

右键以管理员身份运行

在powershell中分别执行如下四个命令

```
$path = "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\DataCollection"
# Telemetry level: 1 - basic, 3 - full
$value = "3"
New-ItemProperty -Path $path -Name AllowTelemetry -Value $value -Type Dword -Force
New-ItemProperty -Path $path -Name MaxTelemetryAllowed -Value $value -Type Dword -Force
```

执行完重启电脑,重新查看windows预览体验计划
