---
layout:     post                    # 使用的布局(不需要改)
title:      闲聊逻辑锁               # 标题
subtitle:   众鸟高飞尽,孤云独去闲.奇变偶不变,符号看象限  #副标题
date:       2019-05-29              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 常识
---

硬盘逻辑所的原理就是在MBR分区表中把扩展分区的那一部分通过修改十六进制表将逻辑驱动器的地址指针做成死循环(类似循环链表)

破解方法大致是在汇编环境下修改JMP,然后PE下用DiskGenius等软件修改分区表16进制文件恢复扩展分区线性结构

晦涩难懂的方法在[这里](https://www.52pojie.cn/thread-844097-1-1.html),不容易学会但是讲的很细致

对于新手来说,上面的方法很难实现

那么,跳出来想一下,既然修复很难,为什么不从源头上规避呢?

硬盘逻辑锁,顾名思义跟"逻辑"有关,一块硬盘(MBR)可以分成四个分区

扩展分区中可以再分无数个逻辑分区

看到"逻辑"了吗?

![1](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/LogicLock/1.webp)

所以

1. 不要使用扩展分区,四个主分区还觉得不够用,教你个方法:只分一个区,根目录下新建26个文件夹命名从A到Z,满足你的心理需求

2. 如果还觉得不过瘾,我就是喜欢我的电脑上逻辑驱动器像春运一样多,不要使用MBR分区表,使用比较新的GPT分区表,支持高达128个主分区
