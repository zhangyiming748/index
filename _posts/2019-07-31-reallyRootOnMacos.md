---
layout:     post                    # 使用的布局(不需要改)
title:      macOS开启真正的root用户              # 标题
subtitle:   她趴在桌上问我:喂,班上有你喜欢的人吗?我看了她一眼淡淡的说:有啊.她心里一阵失落但还是装作无所谓的问道:谁啊?你猜.她把全班其他女生都猜完了我却都摇头,我轻笑摸她头:傻瓜,你确定都猜完了吗?她听后愣了一会,害羞的转过头没再说话.这时我温柔的伏在她耳边轻声说道:还有男生的名字你没猜! #副标题
date:       2019-07-31  08:00:00            # 时间
author:     Base on Xie Bruce                      # 作者
header-img: img/post-bg-2015.jpg    #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - Macintosh OS
---

[这是原文地址](https://www.xiebruce.top/809.html/comment-page-1#comment-289)
之所以要盗过来,就是怕这个博文哪天消失了,留个备份,因为这确实是我需要的功能

----

右键`访达`选择`前往文件夹`或者`shift+command+G`

输入`/System/Library/CoreServices/Applications`
![1](https://img.xiebruce.top/2019/01/26/3d58bab278d09bab1fd8e910600c1699.png)

找到`目录实用工具`并打开
![2](https://img.xiebruce.top/2019/01/26/ae23e53657cf74d121dde5d9dc6d834b.png)

点击左下角的锁,并打开

![3](https://img.xiebruce.top/2019/01/26/1c6349f22caaff534c70c93fbb4f74a5.jpg)

点击顶部菜单栏的`编辑` `启用root用户`

![4](https://img.xiebruce.top/2019/01/26/6a1f5a338e436671bb55fe1f28efa47e.jpg)

给root用户设置密码,设置成功后root用户开启
![5](https://img.xiebruce.top/2019/01/26/e73ed3fcf7358fc1006f9392cd22ddd0.jpg)

开启之后就能在用户登录界面的其他用户里选择root登陆了

终端环境默认root权限

在管理员用户下的`sudo su` `sudo -i` `sudo -s`并不是真正的root,即使能做的已经很多了

类比正版Windows系统的administrator账户

----

没有需求的人尽量别开
