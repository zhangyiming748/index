---
layout:     post                    # 使用的布局(不需要改)
title:      Windows更新死循环解决      # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-07-20 00:00:05 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Windows
---

在Win10更新或升级的过程中，有时进展可能不太顺利，尤其是反复出现更新失败提示“无法完成更新，正在撤销更改”时很让人失望。不过这个问题也能解决，而且也不难，一个命令就可以搞定。

进入WindowsRE环境
执行以下命令

`DISM /image:系统盘符:\ /cleanup-image /revertpendingactions`
