---
layout:     post                    # 使用的布局(不需要改)
title:      Raspberry作为samba服务             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-01-15 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: True                     # 是否归档
tags:                               #标签
    - Raspberry
---

## 首先设置硬盘的自动挂载
我的设备是从旧电脑上拆下来的NGFF固态硬盘,设备文件名`/dev/sda`,分区`sda1`已经格式化为`ext4`

首先手动挂载一遍`sudo mount /dev/sda1 /mnt`(如果已经被自动挂载需要先卸载,自动挂载点文件路径太长)

接下来打开`/etc/fstab`添加一行`/dev/sda1 /mnt  ext4  defaults  0 0`然后`mount -a`验证一下写的没问题(否则开机进入emergency环境)之后重启
## 确保没有问题后卸载设备进行samba安装
安装samba程序
`sudo apt-get install samba samba-common-bin`
安装过程中如果提示需要安装额外的包选择`yes`
## 进入磁盘新建用于samba的文件夹

```
sudo mkdir disk
sudo chown -R root:users /mnt/disk
sudo chmod -R ug=rwx,o=rx /mnt/disk
```

## 修改samba的配置文件`/etc/samba/smb.conf`

在
```
####### Authentication #######```
```
下面加上这行文字
`security = user`
把
```
#======================= Share Definitions =======================
[homes]
    comment = Home Directories
    browseable = no
# By default, the home directories are exported read-only. Change the
# next parameter to 'no' if you want to be able to write to them.
    read only = yes
```
中的`yes`改成`no`

在文件末尾添加
```
[public]
  comment = public storage # 名称随意
  path = /mnt/disk/public # 路径是你打算用来开放的磁盘路径
  valid users = @users
  force group = users
  create mask = 0660 # 连接到服务的用户创建文件时的文件权限
  directory mask = 0771 # 连接到服务的用户创建目录时的目录权限
  read only = no
```
## 重启smb服务或重启电脑
`sudo /etc/init.d/samba-ad-dc restart`
## 新建一个samba用户
`sudo smbpasswd -a pi`
这就是后期用来访问smb服务的用户名和密码
## 验证
![通过其他电脑连接](https://z3.ax1x.com/2021/01/15/sw6IVs.png)
