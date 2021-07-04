---
layout:     post                    # 使用的布局(不需要改)
title:      Windows11OnARM卸载多余UWP      # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-07-04 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Raspberry
---

|要删除的软件|命令|
|:---|---:|
|MSN天气|Remove-AppxPackage Microsoft.BingWeather_1.0.6.0_arm__8wekyb3d8bbwe|
|人脉|Remove-AppxPackage Microsoft.People_10.1909.12456.0_arm64__8wekyb3d8bbwe|
|使用技巧|Remove-AppxPackage Microsoft.Getstarted_10.2.41172.0_arm64__8wekyb3d8bbwe|
|便笺|Remove-AppxPackage Microsoft.MicrosoftStickyNotes_4.0.2.0_arm__8wekyb3d8bbwe|
|必应新闻|Remove-AppxPackage Microsoft.BingNews_1.0.6.0_arm__8wekyb3d8bbwe|
|搜索|Remove-AppxPackage -AllUsers Microsoft.Windows.Search_1.16.0.22000_neutral_neutral_cw5n1h2txyewy(目前无法删除)|
|地图|Remove-AppxPackage Microsoft.WindowsMaps_1.0.22.0_arm64__8wekyb3d8bbwe|
|录音机|Remove-AppxPackage Microsoft.WindowsSoundRecorder_1.0.38.0_arm64__8wekyb3d8bbwe|
|相机|Remove-AppxPackage Microsoft.WindowsCamera_2020.503.58.0_arm__8wekyb3d8bbwe|
|计算器|Remove-AppxPackage Microsoft.WindowsCalculator_10.2012.21.0_arm__8wekyb3d8bbwe|
|闹钟和时钟|Remove-AppxPackage Microsoft.WindowsAlarms_1.0.36.0_arm64__8wekyb3d8bbwe|
|你的手机|Remove-AppxPackage Microsoft.YourPhone_0.19051.7.0_arm__8wekyb3d8bbwe|
|获取帮助|Remove-AppxPackage Microsoft.GetHelp_10.2008.32311.0_arm__8wekyb3d8bbwe|
|Xbox|Get-AppxPackage \*Xbox\* \| Remove-AppxPackage|
|邮件和日历|Remove-AppxPackage microsoft.windowscommunicationsapps_16005.12827.20400.0_arm__8wekyb3d8bbwe|
