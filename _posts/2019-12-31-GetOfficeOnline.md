---
layout:     post                    # 使用的布局(不需要改)
title:      在线安装全功能正版Office            # 标题
subtitle:   2020年消灭贫困人口指的是在2020年,所以大家还有一年时间翻身 #副标题
date:       2019-12-31 23:59:59 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/jiuzhaigou.webp    #这篇文章标题背景图片
catalog: True                       # 是否归档
tags:                               #标签
    - 技巧
---

## Office官网得到xml

皇帝的网址链接:

<font color="#FFFFFF">https://config.office.com/deploymentsettings</font>

1. 选择自己需要安装的组件
2. 全部选择完成后点击右上方导出
3. 选择xml格式
4. 接受许可条件
5. 选择文件名
6. 导出

## Microsoft官网得到部署工具

皇帝的网址链接:

<font color="#FFFFFF">https://www.microsoft.com/en-us/download/details.aspx?id=49117</font>

## 解压部署工具

双击下载的部署工具,解压到任何你喜欢的位置,比如桌面;会得到setup.exe和几个xml文件,如果你之前自己生成过xml,就不需要这几个默认的xml,最好别用,改首选语言挺麻烦的

## 运行部署工具


以管理员身份运行cmd或powershell

+ cmd:

`cd %userprofile%\Desktop`

`setup.exe /configure configuration.xml`

+ powershell:

`cd %userprofile%\Desktop`

`./setup.exe /configure ./configuration.xml`

等待完成即可

## 激活

如果运气好赶上了第一波,有可能直接就是大客户批量激活,不过现在可能连车尾灯都看不见了

不过我比较善良,以下是部分皇帝的神秘代码

<font color="#FFFFFF">if %i%==1 set KMS_Sev=kms7.MSGuides.comif<br>%i%==2 set KMS_Sev=kms8.MSGuides.comif<br>%i%==3 set KMS_Sev=kms9.MSGuides.comif<br>%i%==4 goto notsupported</font>

不敢发全,毕竟涉及到版权
