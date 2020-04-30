---
layout:     post                    # 使用的布局(不需要改)
title:      用启动盘安装Windows7              # 标题
subtitle:   欲渡黄河冰塞川,石灰石加稀硫酸  #副标题
date:       2019-05-20              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 操作系统
---
首先声明,这是不带任何推广软件和广告的原版系统,不要认为那些一键装系统的软件真的好心让你免费一键装系统,基本上的套路都是捆绑主页,你打开一次他挣两分钱,各种流氓管家卸不掉,各种必要的运行库和系统功能精简掉,如果你真觉得这些你都不在乎,可以线下来问问我还有什么,因为这是大部分人的利益,说出来断了大家财路.
安装原版系统自然会比盗版系统花费的时间长,这也是为什么电脑店的奸商不给你装正版的原因

言归正传,进入正题
要制作一个可启动的U盘,制作方法在[另一篇blog](https://zhangyiming748.github.io/2019/05/16/make_a_bootable_usb_disk/)

制作好启动盘之后,以下是单项选择题,三个启动U盘的方法你选择其中一个:
1. 可以通过电脑上的快捷启动按键来启动你的U盘.(常见F12)
(就是在按下电源键的时候一直敲敲敲敲,一直敲敲敲就可以出现可启动设备的选项,然后选择U盘来启动)
2. F1或F2热键进入主板设置

----
当你成功启动U盘的时候,会出现操作系统的安装界面,就可以开始安装你的操作系统了
出现画面之后什么都不要点,先按shift+F10,进入下面的命令提示符,按照图中代码输入
如果shift+F10不好使就试试Fn+shift+F10 有些主板默认Fn反转
![1](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installWindows7/1.webp)
关于磁盘分区的弊端,以前的blog应该有,自己去看,而且不论你多么想分区,都不要在这里分区,进系统之后再分

![2](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installWindows7/2.webp)
其中有一步选择自定义或者升级安装,一定要选择自定义

点下一步
就会开始安装操作系统了

![3](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installWindows7/3.webp)
然后电脑会自动重启,,然后会出现很多设置,,自己选择自己喜欢的设置,然后就安装完成了!

如果以上这些你还看不懂或者不想看懂的话,这里有整个过程的视频

<iframe width="560" height="315" src="https://www.youtube.com/embed/czfcIeUVMEo" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

如果你还抱怨为什么我电脑上和你显示的不一样的话,那么恭喜你,你是填鸭式教育的成果,我猜接下来你要问的问题就像你在网吧问网管:'我的电脑屏幕上为什么比别人少了一个快捷方式?你是不是看不起我'之类的,学习的意义不在于你学到了什么,而是学会如何学习

如果你还是没懂我什么意思的话,你赢了
