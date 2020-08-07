---
layout:     post                    # 使用的布局(不需要改)
title:      Android VS IOS             # 标题
subtitle:   征集副标题 #副标题
date:       2019-12-19 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/jiuzhaigou.webp    #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 杂谈
---
Android可以自由导入导出文件.

Android可以通过更多渠道获取应用.

Android可以自己进行刷机自定义系统.

Android可以安装软件插件.

Android可以不用交30%智商税.

Android热点不会自动关闭.

Android有多功能NFC.

Android可以自定义语音助手.

Android有自动点击器.

Android有自动关机.

Android可以自定义动态壁纸.

Android图片文件夹分类明确清晰.

Android不用受到百度云智障限速.

Android自带准确的流量统计.

Android可以自定义桌面.

Android有悬浮窗.

Android有息屏显示.

Android有实时网速监控.

Android可以自定义状态栏图标项目.

Android各种方法支持长截图.

Android有息屏快捷手势.

Android各种姿势清理垃圾文件.

Android可以设置图案解锁.

Android可以自定义动画幅度.
……
以上都是原生Android和iOS 13 Developer Beta 4的对比.很多深度定制\自定义系统的黑科技均未列举.

不接受抬杠.以上内容全部属实,你不会弄是你的事情.
----
# 关于iPhone不卡
+ Android优先响应程序
就像pc,在跑某个大型程序的时候,可能任务管理器已经显示无响应,但是程序依然在工作,工作结束后才能恢复响应,形象比喻就是:我现在要专心忙我的事,没时间跟你互动,等我忙完了会通知你.而从用户角度看,屏幕没反应了,过一会才好
+ iOS优先响应屏幕
不论工作有没有完成,先给你出个图看,然后来个过渡动画,只要屏幕一直在动,脑残粉根本不知道实际上程序一直在停止响应

# 关于权限和隐私
+ Android在使用权限前会告知
+ iOS对于一些人来说,iPhone应用不会出现不给权限就不让用的情况,让这些人感到了守护住了自己隐私的安全感和优越感,比起Android,经常会出现一些不给权限就不让用的APP,所以iOS养老.实际上呢?举个例子
```
    //获取通讯录
    @IBAction func getContact(_ sender: Any) {
        ContactManager().getContacts()
    }
```
```    
    //获取定位信息
    @IBAction func getLocation(_ sender: Any) {

        LocationManager.shared.getLocationInfo { (error, locationModel) in

           print(locationModel?.latitude, locationModel?.longitude, locationModel?.placemark?.addressDictionary!["FormattedAddressLines"],  locationModel?.placemark?.addressDictionary!["Name"])

            }
    }
```

数据从哪来呢?国区AppleID信息存在云上贵州
而且之前AppStore的信用卡盗刷事件中奖的大部分都是云上贵州并且开启iCloud钥匙串的用户

所以,Android的策略是我需要的权限我问你给不给,你不给我就不让你用;iOS的策略是我不需要征求你意见,你不给我就加一行代码去信息中心获取
# 关于缓存
人均985211的这里大都知道程序运行时必须有一部分缓存是要留在手机里的,无论什么系统
但是经常会有推广联盟的十五秒魔性广告点一下玩一年装备不花一分钱,Android手机可以找到缓存目录删除后新建同名文件设置权限700,阻止程序再次创建这个缓存广告的文件夹,iOS呢?先来个文件管理器?
