---
layout:     post                    # 使用的布局(不需要改)
title:      redis常用命令             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-04-12 00:00:03 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                    # 是否归档
tags:                               #标签
    - 学习
---
# String
|命令|语法|返回值|
|:---:|:---:|:---:|
|set|set key value|OK(存在直接覆盖)|
|get|get key|value / nil|
|del|del key...key|删除的个数|
