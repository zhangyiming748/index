---
layout:     post                    # 使用的布局(不需要改)
title:      关闭Big Sur更新提示             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-11-13 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - macOS
---

![aWsmVJ.jpg](https://www.apple.com/newsroom/images/product/os/macos/standard/apple_macos-bigsur-availability_redesign_11122020_big.jpg.large_2x.jpg)

Big Sur 虽然是是一个很好的新系统,但是对于开发者来说它的安全性真的是不敢恭维,从Mojave升级到Catalina后硬盘整体变成加密盘的软连接,举个例子,从python3.8升级到3.9，常规方法升级后pip3指向的还是3.8，想手动更改到3.9，只能通过关闭系统完整性检查，再手动更改软连接实现，安全性提升还带来很多麻烦，对于开发者来说这并不是一件好事，我并不想升级，但是不敢保证哪天一不小心就点上了升级

如果你也被它所困扰,那就按照如下步骤进行操作:

1. 打开"系统偏好设置",点"软件更新".

2. 取消选择"自动保持我的Mac最新"

3. 点击"高级"按钮

4. 取消所有的勾选

5. 点"好"

6. 退出系统设置,打开终端命令行,分别执行如下3步操作:

`sudo softwareupdate --ignore "macOS Big Sur"`

`defaults write com.apple.systempreferences AttentionPrefBundleIDs 0 `

`killall Dock`

----
恢复

`sudo /usr/sbin/softwareupdate --reset-ignored`
