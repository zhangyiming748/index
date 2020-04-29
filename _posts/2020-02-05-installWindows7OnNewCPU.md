---
layout:     post                    # 使用的布局(不需要改)
title:      在5代以上的CPU上安装Win7         # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-02-05 00:00:01 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: Ture                      # 是否归档
tags:                               #标签
    - 操作系统
---


# Step 1 下载安装镜像
+ 去[这里](msdn.itellyou.cn)自行下载win7x64版本安装镜像或[直接下载](
ed2k://|file|cn_windows_7_ultimate_with_sp1_x64_dvd_u_677408.iso|3420557312|B58548681854236C7939003B583A8078|/)

# Step 2 处理安装镜像
+ 使用任何你擅长的工具把镜像文件中`resource\install.wim`提取出来
+ 使用[DISM++](http://cdn.chuyu.me/Dism++10.1.1001.10_d4ba4eb035254b3326d6adc6638bc9c8daea7018.zip)或任何你喜欢的方式挂载这个wim包到本机的任意一个空文件夹下,这里要注意你的绝对路径不要含有任何中文和特殊符号
+ 以dism++为例,左侧找到驱动管理(我已经很长时间没用Windows了,不过大概是这个意思)
+ 把Intel或AMD的XHCL芯片组驱动文件解压添加适合自己的到当前wim下
+ dism++ 任务栏找到`卸载映像`此时wim文件处理完毕
+ 解压iso中全部文件为一个文件夹,这里的install.wim替换成刚才制作好的wim
+ 还是以dism++为例,你也可以选择你喜欢的工具(我个人喜欢oscdimg)有个打包镜像的工具,选择这个新文件夹打包,就会生成一个新的iso文件

# Step 3 制作启动盘
+ [制作启动盘方法](https://zhangyiming748.github.io/2019/05/16/make_a_bootable_usb_disk/)
+ 一定要选择GPT(非CSM)
+ 不要抖机灵修改别的东西

# Step 4 处理启动盘 添加uefi支持
+ [方法](https://zhangyiming748.github.io/2019/11/13/MakeWindows7SupportUEFI/)是分割线上半部分

# Step 5 使用启动盘开始安装

+ [方法](https://zhangyiming748.github.io/2019/11/13/MakeWindows7SupportUEFI/)是分割线下半部分

# Step 6 安装注意事项
+ 系统盘启动后第一个界面千万要冷静,不要把当年无脑下一步安装流氓软件的劲头拿出来
+ 接下来在命令提示符中不可省略的命令有

```
diskpart //打开磁盘工具
list disk //列出当前可用磁盘
select disk 0 //选择你的磁盘号,不一定是0
clean //初始化磁盘
convert gpt //转换成GUID分区表
exit //退出磁盘工具
exit //退出命令提示符
```
+ 选自定义
+ <font size="15"><b>千万别分区<br>千万别分区<br>千万别分区<br>千万别分区<br>千万别分区<br>千万别分区<br>千万别分区</b></font>

# Step 7 等待完全安装好进入OOBE阶段再去操作

# Step 8 打完收工

----

如果你成功了,请给我寄一箱N95口罩
