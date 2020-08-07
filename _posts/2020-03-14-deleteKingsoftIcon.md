---
layout:     post                    # 使用的布局(不需要改)
title:      删除此电脑中金山云顽固图标       # 标题
subtitle:   我是一只被禁足的安小鸟  #副标题
date:       2020-03-14   00:00:11 GMT+0800           # 时间
author:     Zen                      # 作者
header-img: img/photo/birdAngle.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 技巧
---

我不是针对金山,我是觉得所有不提供给用户自行选择关闭的软件都是垃圾
----

![两个无法直接删除的图标](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/jinshanCloud/1.png?raw=true)

`计算机\HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace`

![找到注册表删除此电脑](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/jinshanCloud/2.png?raw=true)

![效果](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/jinshanCloud/3.png?raw=true)

`计算机\HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Explorer\Desktop\NameSpace\{7AE6DE87-C956-4B40-9C89-3D166C9841D3}`

![找到注册表删除侧边栏](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/jinshanCloud/4.png?raw=true)

![效果](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/jinshanCloud/5.png?raw=true)

----
附上验证CLSID的方法shell:::{ED7BA470-8E54-465E-825C-99712043E01C}
