---
layout:     post                    # 使用的布局(不需要改)
title:     CLion实现远程调试             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-04-17 00:00:08 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                   # 是否归档
tags:                               #标签
    - 技巧
---

由于Windows下的gcc工具实在是不好用,而且由于众所周知的原因,安装起来也是非常的麻烦,因为手头正好有个闲置的树莓派,所以我用了另一种思路,远程调试

2020版的CLion这个功能非常方便

以下是我的做法
---

## 工具
Windows系统的电脑

Jetbrains的Clion

树莓派3B

局域网环境

## 步骤

#### 新建一个带远程调试的工具链

1. file➡Settings➡Build,Execution,Deployment➡Toolchains点击左侧加号新建连接
2. 从下拉菜单选择Remote Host并点击Credentials右侧的齿轮图标配置远端连接,在打开的对话框中,将ssh的用户名密码配置进去以登录远端机器

![JZz2NT.png](https://s1.ax1x.com/2020/04/17/JZz2NT.png)

#### 建立相应CMake文件

1. file➡Settings➡Build,Execution,Deployment➡cmake点击加号新建一个cmake配置

![JZzR4U.png](https://s1.ax1x.com/2020/04/17/JZzR4U.png)

## 效果

![JZzgEV.png](https://s1.ax1x.com/2020/04/17/JZzgEV.png)
