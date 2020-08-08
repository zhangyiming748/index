---
layout:     post                    # 使用的布局(不需要改)
title:      关闭Catalina更新提示             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-08-25 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - macOS
---

![aWsmVJ.jpg](https://s1.ax1x.com/2020/08/07/aWsmVJ.jpg)

Catalina每隔一段时间就会自动提示更新,系统更新图标上会显示红色的更新提示,在通知中心也会经常弹出通知,这样就会觉得很烦.

如果你也被它所困扰,那就按照如下步骤进行操作:

1. 打开"系统偏好设置",点"软件更新".

2. 取消选择"自动保持我的Mac最新"

3. 点击"高级"按钮

4. 取消所有的勾选

5. 点"好"

6. 退出系统设置,打开终端命令行,分别执行如下3步操作:

`sudo softwareupdate --ignore "macOS Catalina"`

`defaults write com.apple.systempreferences AttentionPrefBundleIDs 0 `

`killall Dock`

----
恢复

`sudo /usr/sbin/softwareupdate --reset-ignored`
