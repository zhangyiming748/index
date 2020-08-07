---
layout:     post                    # 使用的布局(不需要改)
title:      魔改Windows7支持UEFI       # 标题
subtitle:   2019年双11发展方向:把优惠券藏在大兴安岭丛林深处一颗四百年古树右后方300英尺的刘姓土拨鼠家的车库里,在2小时内找到土拨鼠一家并在车库前合影留念,即可获得5元优惠券  #副标题
date:       2019-11-13 00:59:59 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/jiuzhaigou.webp    #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 操作系统
---

对于有特殊需求的人群,Windows7才是他们需要的操作系统,现在的廉价笔记本很多都是纯efi了,Windows7却原生不支持uefi启动,以下是解决方法
----
1. 从Windows8的安装文件中提取
`Bootmgr.efi`
2. 文件重命名为
`BOOTX64.EFI`
3. 拷贝到win7ISO的
`\EFI\Boot\`下
没有BOOT文件夹新建一个
`Bootmgr.efi`
也可以从已经安装好的Win8/10系统获得
路径`%SystemRoot%\boot`
4. 将修改过的ISO文件[制作为启动盘](https://zhangyiming748.github.io/2019/05/16/make_a_bootable_usb_disk/)
进入安装界面[正常安装](https://zhangyiming748.github.io/2019/05/20/install_win7/)

----
以下是详细过程

Windows8原生支持UEFI,没问题.Windows7不一样,如果是U盘或移动硬盘安装,需要添加UEFI支持文件,否则不能以UEFI方式启动.

Windows7添加UEFI支持文件的方法:从Windows8的安装文件中提取`Bootmgr.efi`文件,重命名为`BOOTX64.efi`拷贝到win7安装文件的`\EFI\Boot\`下,没有BOOT文件夹新建一个`Bootmgr.efi`也可以从已经安装好的Win8/10系统获得,路径`C:\Windows\boot`

进入安装界面和正常安装相同

按`shift+F10`打开命令提示符

把MBR磁盘转换为GPT磁盘

键入

`diskpart` 打开diskpart工具

`list disk`列出系统拥有的磁盘

`select disk 0`选择目标磁盘选择0号磁盘(请根据磁盘大小,自行判断你的目标磁盘)

`clean`清空目标磁盘

`convert gpt`转换为GPT格式

`list partition`列出磁盘上的分区,因为我们刚转换成GPT格式,因此分区为空

`create partition efi size=200`建立EFI分区

`create partition msr size=128`建立MSR分区

`create partition primary size=102400`建立主分区,该分区用来安装win7 X64

`list partition`列出磁盘上的分区,检查一下是否建立好了分区,好了就关掉CMD

这时返回到了选择分区的界面,点刷新后,选择刚才那个主分区,不再赘述

## 彩蛋

鉴于许多人是花粉俱乐部来的,并且也大部分买的是AMD魔法书的用户
这里有一个接近成品的iso,首先声明,我只更改了USB和E(X)HCI驱动,后期你们非要用带有注入流氓软件功能的PE安装与我无关

[百度网盘](https://pan.baidu.com/s/1TIfK9hXU2yZPbYTNU6gCUA)
有缘人才能看到提取码:<font color="#FFFFFF"> jrnf </font>
