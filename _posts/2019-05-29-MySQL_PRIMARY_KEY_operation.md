---
layout:     post                    # 使用的布局(不需要改)
title:      MySQL主键操作               # 标题
subtitle:      #副标题
date:       2019-05-29 11:43:00              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 数据库
---

在我们使用MySQL的时候,有时会遇到须要更改或者删除MySQL的主键,我们能够简单的使用alter table table_name drop primary key;来完成.以下我使用数据表table_test来作了样例.
1. 首先创建一个数据表table_test:
```
create table table_test
(`id` varchar(100) NOT NULL,
`name` varchar(100) NOT NULL,
PRIMARY KEY (`name`))
ENGINE=MyISAM DEFAULT CHARSET=gb2312;
```
2. 如果发现主键设置错了,应该是id是主键,但如今表里已经有好多数据了,不能删除表再重建了,仅仅能在这基础上改动表结构.
先删除主键

`alter table table_test drop primary key;`

然后再增加主键

`alter table table_test add primary key(id);`

注:在增加主键之前,必须先把反复的id删除掉.
