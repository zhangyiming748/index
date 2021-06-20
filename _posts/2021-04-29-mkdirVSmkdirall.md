---
layout:     post                    # 使用的布局(不需要改)
title:      os.Mkdir os.MkdirAll区别             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-04-29 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Golang
---
Mkdir 用于创建单个目录
```
err:=os.Mkdir("./dir1",os.ModePerm)
if err!=nil{
   fmt.Println(err)
}
```
MkdirAll 用于创建文件夹路径
```
err=os.MkdirAll("./dir1/dir2",os.ModePerm)
if err!=nil{
   fmt.Println(err)
   }
err为nil
//成功创建dir1/dir2文件路径
```
