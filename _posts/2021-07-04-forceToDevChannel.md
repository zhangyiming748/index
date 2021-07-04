---
layout:     post                    # 使用的布局(不需要改)
title:      强行加入Windows insider Dev channel      # 标题
subtitle:   垂死病中惊坐起 笑问win11何时来 #副标题
date:       2021-07-04 00:02:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Windows
---

>IT之家 6 月 28 日消息 微软上周正式发布了全新的 Windows 11 操作系统,并将在本周(美国时间)推送 Insider Preview 预览版,之后在圣诞节期间推出正式版且在明年支持升级.

>当然,我们还是建议满足完整的硬件要求的 PC 去升级,毕竟体验会更好一些.此外,微软也表示:
由于这些设备不满足新的硬件要求,因此可能存在影响 Windows 11 在这些 PC 上的体验的问题和 bug,这些问题和 bug 可能无法修复.
如果这些 PC 中的任何一台出现问题而需要回滚到 Windows 10,你可以使用 MCT 工具返回,只是这些 PC 不再允许再次升级到 Windows 11 Insider Preview 版本.它们将被视为新 PC,并且将严格执行上面的最低硬件要求.
一旦 Windows 11 正式到来,这些 PC 可选择退出测试,并且将无法接收之后的 Windows 11 Insider Preview 版本.此外,他们还必须使用微软提供的 ISO 全新安装以回到 Windows 10 预览通道.
此外,微软还将在今年夏天晚些时候发布 Beta 通道的 Windows 11 Insider Preview 版本,Beta 测试中不满足 Win11 硬件要求的 PC 将被移至 Release Preview 频道,其中部分 PC 可自行返回 Beta 通道但风险自负.
感谢IT之家网友热心线索投递,即使不满足最低硬件要求, 2021 年 6 月 24 日之前所有已经在其 PC 上安装 Dev Channel 版本的用户仍可选择升级 Dev 版 Windows11,只是在未来会有限制.

### 对于 6 月 24 日之前没有加入 DEV 通道却又想测试 Win11 的小伙伴,我们这里提供一种方法,不过修改注册表有一定风险,该方法仅供参考,后果自负:

首先,加入 Insider 计划并选择 Release Preview,毕竟你也没有其他选择.

重启,然后打开注册表编辑器.

在左侧导航中找到:
`HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\WindowsSelfHost\UI\Selection`
![1](https://z3.ax1x.com/2021/07/04/RfmjM9.jpg)
1. 将 UIBranch 中的值更改为 Dev
2. 将 ContentType 中的值更改为 Mainline
3. 将 Ring 中的文本更改为 External

然后导航到:
`HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\WindowsSelfHost\Applicability`
![2](https://z3.ax1x.com/2021/07/04/RfmOxJ.jpg)
1. 将 BranchName 中的值更改为 Dev
2. 将 ContentType 中的值更改为 Mainline
3. 将 Ring 中的值更改为 External
4. 退出注册表编辑器,重启.

OK,若你重启之后发现现在变成了 DEV 通道那么恭喜你,现在只需耐心等待微软准时推送测试版即可;万一操作失败的话则建议放弃电脑
