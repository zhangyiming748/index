---
layout:     post                    # 使用的布局(不需要改)
title:      Git常用commit参数       # 标题
subtitle:   有些人看上去风风光光,背地里都是我替他收蚂蚁森林的能量   #副标题
date:       2019-12-10 00:59:59 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/jiuzhaigou.webp    #这篇文章标题背景图片
catalog: False                      # 是否归档
tags:                               #标签
    - 杂谈
---

`git add main.go`
开始跟踪新文件/把已跟踪的文件放到暂存区

`git commit`
`git commit -m "message"`
提交修改到本地

`git push git@url.git gethistory_api`
推送提交的修改到远程

`git clone git@url.git <dirName>`
克隆现有代码库

`git status`
查看文件状态

`git branch -a`
查看当前分支

`git clone -b <branchName> git@url.git`
克隆指定分支的代码库

`git checkout -b feature_x`
创建一个叫做`feature_x`的分支,并切换过去:

`git branch -d feature_x`
删掉分支

`git branch -vv`
输出映射关系

`git branch -u origin/dev`
将当前本地分支与远程分支建立映射关系,dev为远程分支名

`git branch --unset-upstream`
撤销当前本地分支与远程分支的映射关系
