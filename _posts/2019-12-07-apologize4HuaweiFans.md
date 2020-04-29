---
layout:     post                    # 使用的布局(不需要改)
title:      我错了,我把大家带跑偏了       # 标题
subtitle:   我一直都想把北京房子卖了,然后回到老家,换个大房子,买上一辆车,过上慢节奏走生活,可房东一直不同意  #副标题
date:       2019-12-07 00:59:59 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/jiuzhaigou.webp    #这篇文章标题背景图片
catalog: False                      # 是否归档
tags:                               #标签
    - 杂谈
---

之前发布的一篇[blog](https://zhangyiming748.github.io/2019/11/18/RebuildWirelessCard/)和在花粉论坛上的若干帖子,跟着我思维被带跑偏的人也有几个
现在来完整说一遍我犯的错误

----

我一开始的思路是:
1. 我的电脑上有一个写着正确数据(SN\MAC\DEVICEID)的NFC贴片
2. 我要拿到这个数据
3. 拿到数据后修改为我新网卡的MAC地址
4. 重新写到新的NFC芯片上
其中第三步是困扰了我很长时间的问题
现已确定MagicBook Pro连接时确实会显示
![RFID](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/Honor/MagicBookProRFID.jpg)<center>现已确定MagicBook Pro连接时确实会显示</center>
卡在这一步已经很长时间了,痛定思痛,突然茅厕顿开
为什么不换一种方法?我可是有电脑的人啊

跳出思维定势推倒重来:
1. 我有电脑
2. 我电脑识别我的网卡
3. 电脑管家可以生成新的二维码
4. 我可以通过二维码生成正确值的NFC芯片
5. 虽然整个过程中我不知道NFC中的DEVICESID数据,但是电脑和手机都知道,这就足够了

----

对于曾经跟我一起绞尽脑汁搞到devicesId的老铁们再次表示歉意
发个红包,高抬贵手,支付宝扫一下,买个QQ糖吃

![Red](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/Honor/red.png)
