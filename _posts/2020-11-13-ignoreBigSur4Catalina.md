---
layout:     post                    # 使用的布局(不需要改)
title:      Catalina下关闭Big Sur更新提示             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-11-13 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - macOS
---
伴随着macOS Catalina 10.15.5的发布以及适用于macOS Mojave以及High Sierra的安全更新,苹果对安全软件更新策略做出进一步调整,从而让用户更难忽视这些可用的软件更新,以便于让系统保持最新状态
我没钱买M1芯片的电脑是我的错,但你让我强制升级x86的电脑到Big sur就有点过分了,之前用其他硬盘安装过,体验非常不好,可能这个系统目前只适合ARM,但是
在使用带"--ignore"的softwareupdate命令时候,不再忽略macOS的重要新版本
在安装Security Update 2020-003之后,这项更新策略同样适用于macOS Mojave和macOS High Sierra.
![](https://static.cnbetacdn.com/thumb/article/2020/0529/8e869eac1451303.jpg)
尊重是互相的,你不拿消费者当人,我也没必要留着你
![aWsmVJ.jpg](https://www.apple.com/newsroom/images/product/os/macos/standard/apple_macos-bigsur-availability_redesign_11122020_big.jpg.large_2x.jpg)

Big Sur 虽然是是一个很好的新系统,但是对于开发者来说它的安全性真的是不敢恭维,从Mojave升级到Catalina后硬盘整体变成加密盘的软连接,举个例子,从python3.8升级到3.9,常规方法升级后pip3指向的还是3.8,想手动更改到3.9,只能通过关闭系统完整性检查,再手动更改软连接实现,安全性提升还带来很多麻烦,对于开发者来说这并不是一件好事,我并不想升级,但是不敢保证哪天一不小心就点上了升级

如果你也被它所困扰,那就按照如下步骤进行操作:

打开终端,选择你喜欢的编辑器输入
`sudo sublime /etc/hosts`
添加以下文本
```
127.0.0.1 swscan.apple.com
127.0.0.1 swdist.apple.com
127.0.0.1 swquery.apple.com
127.0.0.1 swcdn.apple.com
127.0.0.1 swdownload.apple.com
```
屏蔽掉升级服务器 然后
```
//删除图标的小红点
defaults write com.apple.systempreferences AttentionPrefBundleIDs 0
killall Dock
```
