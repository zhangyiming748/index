---
layout:     post                    # 使用的布局(不需要改)
title:     为Win10-UWP应用使用设置代理             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-04-26 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                   # 是否归档
tags:                               #标签
    - 技巧
---
Win 10 的 UWP应用(应用商店下载的APP,如OneDrive\我的世界基岩版等),默认是不走代理的(沙盒的网络隔离特性:禁止APP访问localhost ),即使使用了「小飞机」软件,UWP软件仍然无法得到加速.
网上有很多解决此类问题的方法,要不就是使用 Fiddler 软件进行设置,要不就是使用CheckNetIsolation命令来解除限制,但是需要通过注册表来获取应用的SID,都比较麻烦.
其实有一个办法,不需要下载软件,也不需要打开注册表.

----

## 解除限制的方法

1. 获取名称
  按`Win+R`输入
  `C:\Users\%username%\AppData\Local\Packages`
  打开后能看到的每一个文件夹就代表一个UWP程序包,文件夹的名称就是程序包的名称.
  ![J2v83D.png](https://s1.ax1x.com/2020/04/26/J2v83D.png)
  例如:Microsoft.Office.OneNote_8wekyb3d8bbwe
2. 解除限制

  使用的是CheckNetIsolation命令.该命令的语法是:
  ```
  CheckNetIsolation LoopbackExempt [operation] [-n=] [-p=]
        操作列表:
            -a  -  向环回免除列表中添加 AppContainer 或程序包系列.
            -d  -  从环回免除列表中删除 AppContainer 或程序包系列.
            -c  -  清除环回免除的 AppContainer 和程序包系列的列表.
            -s  -  显示环回免除的 AppContainer 和程序包系列的列表.

        参数列表:
            -n= - AppContainer 名称或程序包系列名称.
            -p= - AppContainer 或程序包系列安全标识符(SID).
            -?  - 显示 LoopbackExempt 模块的此帮助消息.
  ```
以OneNote为例,在命令行里输入:`CheckNetIsolation.exe LoopbackExempt -a -n="Microsoft.Office.OneNote_8wekyb3d8bbwe"`命令执行后,系统会提示完成,该UWP软件无法访问localhost的限制已经解除
或者
计算机\HKEY_CURRENT_USER\SOFTWARE\Classes\Local Settings\Software\Microsoft\Windows\CurrentVersion\AppContainer\Mappings\S-1-15-2-1333615684-3476610416-3051081704-793991562-341351578-2034363004-642090109aaaa
后
CheckNetIsolation.exe loopbackexempt -a -p=S-1-15-2-2551677095-2355568638-4209445997-2436930744-3692183382-387691378-1866284433
