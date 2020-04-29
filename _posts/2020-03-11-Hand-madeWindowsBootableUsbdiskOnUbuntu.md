---
layout:     post                    # 使用的布局(不需要改)
title:      Ubuntu下手工制作Windows启动盘       # 标题
subtitle:   我是一只被禁足的安小鸟  #副标题
date:       2020-03-10    00:00:01 GMT+0800          # 时间
author:     Zen                      # 作者
header-img: img/photo/birdAngle.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 操作系统
---
[原文地址](https://www.onetransistor.eu/2014/09/make-bootable-windows-usb-from-ubuntu.html)
----
# 梗概

#### 准备USB驱动器
1. 启动`GParted`
2. 选择要制作的u盘,卸载所有已经挂载的分区
3. 新建分区表,可选msdos(相当于MBR)或gpt(相当于仅UEFI),选择后apply应用更改
4. 新建分区,MBR对应NTFS,GPT对应FAT32,如果是MBR别忘记标记分区为可启动

#### 复制ISO文件
1. 挂载ISO并把所有文件复制到U盘分区
2. 如果挂载不了可以尝试解压缩后复制

#### 使u盘可启动
1. MBR可启动
    1. `sudo grub-install --target=i386-pc --boot-directory="/media/<username>/<drive_label>/boot" /dev/sdX`
    说明: 其中`/media/<username>/<drive_label>/boot`替换为你自己的目录,`/dev/sdX`替换为你自己的u盘设备,切记是设备而不是分区(dev/sdx1)
    等待完成
    2. 创建一个文本文件`grub.cfg`保存到u盘`boot/grub`文件夹下,命令为`touch /media/<username>/<drive_label>/boot/grub/grub.cfg`
    文件内容为

    ```
    default=1  
    timeout=15
    color_normal=light-cyan/dark-gray
    menu_color_normal=black/light-cyan
    menu_color_highlight=white/black

    menuentry "Start Windows Installation" {
        insmod ntfs
        insmod search_label
        search --no-floppy --set=root --label <USB_drive_label> --hint hd0,msdos1
        ntldr /bootmgr
        boot
    }

    menuentry "Boot from the first hard drive" {
        insmod ntfs
        insmod chain
        insmod part_msdos
        insmod part_gpt
        set root=(hd1)
        chainloader +1
        boot
    }
    ```

    3. 用vi或者你喜欢的编辑器写入就好
    4. 制作完成,这种方法支持legacy bios和uefi csm
2. GPT可启动
    1. 检查u盘下efi文件夹中是否有文件`bootx64.efi`或`bootia32.efi`,如果有,制作完成可以使用了
    2. 如果没有,挂载之前的ISO找到install.wim或install.esd解压到单独的文件夹或者直接挂载
    3. 进入目录`./windows/BOOT/efi`提取出`bootmgfw.efi`重命名为`bootx64.efi`放入u盘中的`efi/boot`目录,制作完成

----
## 应观众要求,以下是图文教程

![insertUdisk](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/insertUdisk.png?raw=true)<center>插入U盘</center>
![listDisk](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/listDisk.png?raw=true)<center>查看当前磁盘状态</center>
![deleteFirst](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/deleteFirst.png?raw=true)<center>重建分区表前必须删除所有分区</center>
![cleanPartition](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/cleanPartition.png?raw=true)<center>清除U盘上全部分区</center>
![createPartable](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/createPartable.png?raw=true)<center>新建分区表</center>
![partitionAttribute](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/partitionAttribute.png?raw=true)<center>新建分区的属性</center>
![createPartitionPrimary](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/createPartitionPrimary.png?raw=true)<center>U盘建立主分区</center>
![mountCommand](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/mountCommand.png?raw=true)<center>挂载ISO命令</center>
![mountISO2Dir](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/mountISO2Dir.png?raw=true)<center>挂载ISO到指定目录</center>
![moveISO2Udisk](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/moveISO2Udisk.png?raw=true)<center>从ISO移动文件到U盘</center>
![udiskTree](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/udiskTree.png?raw=true)<center>制作完成后的目录结构</center>
![selectSystem](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/selectSystem.jpg?raw=true)<center>临时指定启动设备界面</center>
![booting](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/booting.jpg?raw=true)<center>正在进入系统</center>
![intoSystem](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/intoSystem.jpg?raw=true)<center>进入系统</center>
![aigoUdisk](https://github.com/zhangyiming748/zhangyiming748.github.io/blob/master/img/ubuntu/makeBootableUSB/aigoUdisk.jpg?raw=true)<center>十年前买的爱国者U盘</center>
