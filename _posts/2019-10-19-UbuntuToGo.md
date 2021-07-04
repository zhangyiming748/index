---
layout:     post                    # 使用的布局(不需要改)
title:      安装UbuntuToGo的血泪教训             # 标题
subtitle:    #副标题
date:       2019-10-19 04:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/jiuzhaigou.webp    #这篇文章标题背景图片
catalog: True                       # 是否归档
tags:                               #标签
    - Linux
---

实在太困了明天早上写正文

----

原教程在IT之家,自己去查,我只说遇到的问题

分区的时候教程并没有说明,两个分区都必须是主分区,扩展分区和逻辑驱动器的形式无疑是会安装失败的

挂载chroot环境后教程提供的直接安装方法会导致提示`已在其他命令中包括或者已被废弃`实际上只是因为安装源没有更新,两个小时之后我终于意识到这个问题

最终虽然能在虚拟机上成功以两种方式启动(uefi/Legacy)

却不能在单EFI电脑上成功运行
