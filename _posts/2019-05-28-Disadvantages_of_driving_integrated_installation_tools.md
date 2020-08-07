---
layout:     post                    # 使用的布局(不需要改)
title:      驱动集成安装工具的弊端               # 标题
subtitle:   停车坐爱枫林晚,从此君王不早朝  #副标题
date:       2019-05-28              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 常识
---

驱动集成安装工具就是那些所谓一键安装的\帮助你安装驱动的软件,首先说个本身就有流氓软件性质的这类软件,XX精灵,隶属于金山全家桶,安装后右下角弹窗诱导安装金山毒霸,而且使用时一不小心就会被在所有浏览器捆绑主页毒霸网址大全,使用的是JavaScript脚本,常规的恢复注册表方法根本没有效果,根除的方法只能是在%systemroot%下30000多个文件夹中删除这个可能隐藏在任何一个子文件夹里的JS文件,大多数的驱动工具的盈利模式都是广告推广,你用这个软件时间越长,这个软件能从广告商那里得到的利润就越多,那么怎么才能抓住用户呢?通过让自己成为一款良心软件的方式显然是慢而且困难的,那么从另一个角度讲,只要不被用户卸载就相当于抓住用户了,除了大部分软件把"卸载"写得特别小\利用三重否定疑问句询问你是否不要不卸载该软件及其下载的文件,然而给我写下这篇日志最大的决心是卸载相关驱动的uninstall.exe也就是驱动自带的卸载工具直接被这类软件单独删除,而驱动卸载只有通过uninstall.exe才能真正的卸载所有相关的文件和注册表,而且驱动升级实质上也是通过卸载->重新安装新版完成的,驱动集成安装工具恰恰是把这些卸载程序屏蔽或删除了,导致用户卸载了驱动集成安装工具后,之前安装的驱动都无法卸载/升级了,这就是图一时方便的后果.

这类软件不能一棍子打死,还是有比较好的官方工具,比如戴尔的DELL UPDATE,联想ideaPad的联想驱动管理,联想ThinkPad的SYSTEM Update,如果普通用户无法辨别哪一个软件可信,从经济的角度上考虑一下,你买电脑就相当于买了相关的服务,所以电脑官网的肯定是没问题的,至于网上看起来免费的驱动安装工具,非亲非故的,还要给员工发工资,它为什么免费?
想通之后,下一个问题就是去哪里下载驱动了,如果你善于利用百度,接下来说的你都懂

+ Intel
英特尔官方的驱动安装工具,如果你的电脑是组装台式机,使用了Intel的CPU\无线模块\固态硬盘或者其他的英特尔产品,这个工具可以傻瓜式安装驱动

![1](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/DriverSoftware/1.webp)

+ AMD
AMD官方的驱动下载网站,如果你的电脑是组装台式机,使用了AMD的APU\GPU,这个网址可以下载到驱动,使用时需要一定的中文理解能力

![2](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/DriverSoftware/2.webp)

+ NVIDIA
nVidia官方的驱动下载网站,如果你的电脑是组装台式机,使用了英伟达的GPU,这个网址可以下载到驱动,使用时需要一定的中文理解能力和逻辑思维.

![3](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/DriverSoftware/3.webp)

+ DELL
戴尔的官方驱动下载网页,戴尔的品牌台式机和笔记本都可以在这里下载全部的驱动,唯一的难点就是要找到服务标签/快速服务代码(电脑上的贴纸),但是很多人视而不见

![4](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/DriverSoftware/4.webp)

+ Lenovo
联想家用机驱动下载工具,ideapad的台式机和笔记本用这个软件可以下载全部驱动,无脑安装就行,没有捆绑

![5](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/DriverSoftware/5.webp)

+ ThinkPad
虽然已经被联想收购,但还是有自己的驱动安装工具,thinkPad笔记本移动图形工作站,thinkvantage都可以用这个工具,这里有个很有意思的事,在香港官网下载的system update找出来的驱动会比大陆版的驱动更全面

![6](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/DriverSoftware/6.webp)

### 世界上还有很多品牌的电脑这里不能一一列举,原则就是找到品牌官网输入自己电脑的型号去下载驱动,上面的例子如果你没看懂的话,那就放弃吧.
