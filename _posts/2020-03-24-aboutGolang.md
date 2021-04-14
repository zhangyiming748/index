---
layout:     post                    # 使用的布局(不需要改)
title:      golang数据类型       # 标题
subtitle:   我是一只被禁足的安小鸟  #副标题
date:       2020-03-24  00:00:01 GMT+0800           # 时间
author:     Zen                      # 作者
header-img: img/photo/birdAngle.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - Golang
---
golang的基本数据类型包括
+ `bool`布尔类型,格式化输出`%t`
+ `int/uint/float64/complex64/byte/rune`数字类型格式化输出`%d%c%x%u等`
+ `string`字符串类型,值不可变,反引号忽略转义字符 **不可变数据类型**
+ `slice`切片类型,包括头指针,长度,容量
+ `map`映射类型,底层使用哈希表,扩容使用增量扩容,生成新的map,旧map标记已删除
+ `nil`空类型
