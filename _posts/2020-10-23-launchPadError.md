---
layout:     post                    # 使用的布局(不需要改)
title:      macOS启动台无法重命名              # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-10-23  08:00:00            # 时间
author:     zen                      # 作者
header-img: img/photo/birdAngle.webp     #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - macOS
---


1. 通过Finder进入`~/Library/Application\ Support/Dock`文件夹
2. 删除所有db文件`rm *.db`
3. 重启`sudo killall Dock`
