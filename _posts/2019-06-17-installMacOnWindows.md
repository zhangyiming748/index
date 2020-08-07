---
layout:     post                    # 使用的布局(不需要改)
title:      在Windows上用VMware安装Mac虚拟机         # 标题
subtitle:    一部手机的寿命大约在三到五年,只有人类寿命的二十分之一. 手机只是你生活的一部分,而你却是它的全部. 请放下身边的杂事,多陪陪你的手机.   #副标题
date:       2019-06-17 08:02:00              # 时间
author:     Zen                      # 作者
header-img: img/photo/jiuzhaigou.webp   #这篇文章标题背景图片
catalog: True                       # 是否归档
tags:                               #标签
    - 虚拟机
---
在Windows上用VMware安装Mac虚拟机,对买不起MBP的我们来说也算是一种安慰,毕竟也是体验了一回完整的MacOS了,不过这其中的难点在我看来大致有两点,解决方案会在文章中提到,不过写教程的过程中有些地方因为已经安装了没法截图,只要按照步骤来,不需要看图也能解决问题
# 下载VMware
VMware workstation Pro 11/12/14/15

Vmware workstation Player 7/12/14/15

都可以,不过建议Pro15

**去[官网](https://www.vmware.com/cn.html)下载**

**去[官网](https://www.vmware.com/cn.html)下载**

**去[官网](https://www.vmware.com/cn.html)下载**

重要的事情说三遍,下载链接就摆在那里,不会下载是你自己的问题,就这智商,要啥自行车?
# 安装VMware
安装过程对一些人来说也是个问题
不要抖机灵非要安装软件到非系统盘,原因不多说,知道悔改的看[这篇文章](https://zhangyiming748.github.io/2019/05/23/HardDisk_Partition/)

# 下载破解工具
原作者代码库在[这里](https://github.com/DrDonk/unlocker),自行下载Zip包或clone到本地,也可以用Github Desktop,总之用你会的方法
然后得到一个像这样的文件夹结构
```
│  .gitattributes
│  .gitignore
│  darwin.md
│  dumpsmc.exe
│  dumpsmc.py
│  gettools.exe
│  gettools.py
│  LICENSE
│  lnx-install.sh
│  lnx-uninstall.sh
│  lnx-update-tools.sh
│  readme.txt
│  test-unlocker.py
│  unlocker.exe
│  unlocker.py
│  win-build.cmd
│  win-install.cmd
│  win-test-install.cmd
│  win-uninstall.cmd
│  win-update-tools.cmd
│  
└─.git
    │  config
    │  description
    │  HEAD
    │  index
    │  packed-refs
    │  
    ├─branches
    ├─hooks
    │      applypatch-msg.sample
    │      commit-msg.sample
    │      fsmonitor-watchman.sample
    │      post-update.sample
    │      pre-applypatch.sample
    │      pre-commit.sample
    │      pre-push.sample
    │      pre-rebase.sample
    │      pre-receive.sample
    │      prepare-commit-msg.sample
    │      update.sample
    │      
    ├─info
    │      exclude
    │      
    ├─logs
    │  │  HEAD
    │  │  
    │  └─refs
    │      ├─heads
    │      │      master
    │      │      
    │      └─remotes
    │          └─origin
    │                  HEAD
    │                  
    ├─objects
    │  ├─56
    │  │      266d360f3da9f922766101055bd78ffa3724bf
    │  │      
    │  ├─info
    │  └─pack
    │          pack-bafc0500670b8e8721b27120beb394c0829d78c1.idx
    │          pack-bafc0500670b8e8721b27120beb394c0829d78c1.pack
    │          
    └─refs
        ├─heads
        │      master
        │      
        ├─remotes
        │  └─origin
        │          HEAD
        │          
        └─tags
```
# 安装破解工具
用管理员身份运行其中的`win-install.cmd`安装

如果没有报错,万事大吉
如果有,从自身找原因
# 在VMware上安装MacOS
安装选择典型,稍后安装操作系统,然后选择系统界面就能看到苹果系统选项了
如果你不懂怎么配置,这不是本教程考虑的问题,先学会怎么用虚拟机

所有信息配置完成后,先不要着急启动
# 修改信息
在你选择虚拟机的安装目录里,大概的目录结构是
```
2019/06/11  19:19               686 macOS 10.13-0.vmdk
2019/06/11  19:27               686 macOS 10.13-1.vmdk
2019/06/11  19:43               686 macOS 10.13-2.vmdk
2019/06/11  19:57               686 macOS 10.13-3.vmdk
2019/06/13  18:05           270,840 macOS 10.13.nvram
2019/06/13  18:05    31,184,715,776 macOS 10.13.vmdk
2019/03/15  21:23                 0 macOS 10.13.vmsd
2019/06/13  18:05             3,229 macOS 10.13.vmx
2019/06/13  17:40               374 macOS 10.13.vmxf
2019/06/11  21:33           279,966 vmware-0.log
2019/06/11  20:04           281,797 vmware-1.log
2019/06/11  19:57           257,453 vmware-2.log
2019/06/13  18:05           389,305 vmware.log
```
打开`.vmx`配置文件
大概是这样
```
.encoding = "GBK"
config.version = "8"
virtualHW.version = "16"
pciBridge0.present = "TRUE"
pciBridge4.present = "TRUE"
pciBridge4.virtualDev = "pcieRootPort"
pciBridge4.functions = "8"
pciBridge5.present = "TRUE"
pciBridge5.virtualDev = "pcieRootPort"
pciBridge5.functions = "8"
pciBridge6.present = "TRUE"
pciBridge6.virtualDev = "pcieRootPort"
pciBridge6.functions = "8"
pciBridge7.present = "TRUE"
pciBridge7.virtualDev = "pcieRootPort"
pciBridge7.functions = "8"
vmci0.present = "TRUE"
smc.present = "TRUE"
hpet0.present = "TRUE"
ich7m.present = "TRUE"
usb.vbluetooth.startConnected = "TRUE"
board-id.reflectHost = "TRUE"
firmware = "efi"
displayName = "macOS 10.13"
guestOS = "darwin17-64"
nvram = "macOS 10.13.nvram"
virtualHW.productCompatibility = "hosted"
powerType.powerOff = "soft"
powerType.powerOn = "soft"
powerType.suspend = "soft"
powerType.reset = "soft"
tools.syncTime = "TRUE"
sound.autoDetect = "TRUE"
sound.virtualDev = "hdaudio"
sound.fileName = "-1"
sound.present = "TRUE"
numvcpus = "2"
memsize = "4096"
sata0.present = "TRUE"
sata0:0.fileName = "macOS 10.13.vmdk"
sata0:0.present = "TRUE"
sata0:1.deviceType = "cdrom-image"
sata0:1.fileName = "C:\Users\zhang\Downloads\macos in virtualbox\macOSsierra.iso"
sata0:1.present = "TRUE"
usb.present = "TRUE"
ehci.present = "TRUE"
usb_xhci.present = "TRUE"
ethernet0.connectionType = "nat"
ethernet0.addressType = "generated"
ethernet0.virtualDev = "e1000e"
ethernet0.present = "TRUE"
extendedConfigFile = "macOS 10.13.vmxf"
vvtd.enable = "TRUE"
floppy0.present = "FALSE"
vhv.enable = "TRUE"
gui.lastPoweredViewMode = "fullscreen"
numa.autosize.cookie = "20012"
numa.autosize.vcpu.maxPerVirtualNode = "2"
uuid.bios = "56 4d b6 b0 b5 8a 67 47-a4 2e bc 67 e7 a4 f6 68"
uuid.location = "56 4d b6 b0 b5 8a 67 47-a4 2e bc 67 e7 a4 f6 68"
sata0:0.redo = ""
pciBridge0.pciSlotNumber = "17"
pciBridge4.pciSlotNumber = "21"
pciBridge5.pciSlotNumber = "22"
pciBridge6.pciSlotNumber = "23"
pciBridge7.pciSlotNumber = "24"
usb.pciSlotNumber = "32"
ethernet0.pciSlotNumber = "160"
sound.pciSlotNumber = "33"
ehci.pciSlotNumber = "34"
usb_xhci.pciSlotNumber = "192"
vmci0.pciSlotNumber = "35"
sata0.pciSlotNumber = "36"
vmotion.checkpointFBSize = "134217728"
vmotion.checkpointSVGAPrimarySize = "268435456"
ethernet0.generatedAddress = "00:0c:29:a4:f6:68"
ethernet0.generatedAddressOffset = "0"
vmci0.id = "-408619416"
monitor.phys_bits_used = "43"
cleanShutdown = "TRUE"
softPowerOff = "TRUE"
usb_xhci:6.speed = "2"
usb_xhci:6.present = "TRUE"
usb_xhci:6.deviceType = "hub"
usb_xhci:6.port = "6"
usb_xhci:6.parent = "-1"
usb_xhci:7.speed = "4"
usb_xhci:7.present = "TRUE"
usb_xhci:7.deviceType = "hub"
usb_xhci:7.port = "7"
usb_xhci:7.parent = "-1"
toolsInstallManager.updateCounter = "20"
toolsInstallManager.lastInstallError = "21004"
checkpoint.vmState = ""
sata0:2.deviceType = "rawDisk"
sata0:2.fileName = "macOS 10.13-3.vmdk"
sata0:2.redo = ""
tools.upgrade.policy = "useGlobal"
usb_xhci:4.present = "TRUE"
usb_xhci:4.deviceType = "hid"
usb_xhci:4.port = "4"
usb_xhci:4.parent = "-1"
sata0:1.startConnected = "FALSE"
tools.remindInstall = "TRUE"
```
在任意位置添加新的一行`smc.version = 0`

保存退出
# 大功告成
回到VMware,启动虚拟机,坐和放宽,根据硬件配置不同,启动有的时候很慢
