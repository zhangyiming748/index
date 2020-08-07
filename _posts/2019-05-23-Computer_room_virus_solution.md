---
layout:     post                    # 使用的布局(不需要改)
title:      机房病毒(俗称)解决方法              # 标题
subtitle:   我寄愁心与明月,奈何明月照沟渠  #副标题
date:       2019-05-23              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 常识
---
具体表现为u盘插入公共电脑后文件部分消失,残存文件图标改变,u盘容量依然是全部文件

这里假设A1为健康u盘 A2为染毒u盘 B为带毒公共电脑 C为个人干净电脑

A1在插入B后变为A2,此时正确的解决方法是找到一台C,插入u盘,cmd输入:
```
cd d:  //(这里假设D是u盘根目录,自行更改)
```

```
attrib -s -h -a /S /D *.*
del /S /Q *.exe
```
如果在这之前你不小心点击了u盘中任何看似文件夹的可执行程序,那么你现在这台电脑C已经成为了B
解决方法如下:
注册表编辑 [HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run] 右边open＝异常的exe文件 删除即可
