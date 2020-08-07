---
layout:     post                    # 使用的布局(不需要改)
title:      MySQL基本命令               # 标题
subtitle:      #副标题
date:       2019-05-29 11:42:00              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - mysql学习
---

1. `show databases;`
查看所有的数据库,等同于select schema_name from information_schema.schemata\G.\G 替换;,以纵向报表的形式输出结果,有利于阅读.

2. `status`
查看mysql数据库的运行状态

3. `use` 选择数据库 例如 use information_schema

4. `show tables`
查看数据库中的表

5. `desc  table_name;`
查看表结构

6. `show table status from db like 条件`
查看表状态
