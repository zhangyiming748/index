---
layout:     post                    # 使用的布局(不需要改)
title:      忘记MacBook密码解决方法              # 标题
subtitle:   我扌's your problem? #副标题
date:       2019-07-09  18:00:00            # 时间
author:     Base on ITHome                      # 作者
header-img: img/post-bg-2015.jpg    #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - Macintosh OS
---
其实相较于Windows,macOS整天哪哪都需要验证密码,更不容易忘,但这不是连MacBook Air都有指纹识别了嘛,忘记密码的风险也就大了一些,了解一些技巧或多一手准备还是很有必要的.
和Windows一样,方法很多,Command + R + Terminal + resetpassword.对不对呢,也对,可以解决问题,但并非是所有情况下的最优解.下面和大家分享几种解决方案.

----

# 使用Apple ID重设密码(已开启允许)
很像利用Microsoft账户重置PIN码,这是个人认为最理想的方式了,方便而安全.
使用该方式的前提是用户已经登录Apple ID/iCloud账户,且已经在"系统偏好设置→用户与群组"中,开启"允许用户使用Apple ID重设密码".

![1](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/macOSpasswd/1.png)

忘记密码时,只需要在登录界面输错三次,系统便会跳出密码提示和使用Apple ID重设密码选项.需要注意的是,点击文本是没有用的,需要点击右侧类似"Play键"的按钮才会有反应

点击后系统会要求输入Apple ID和密码来验证.

填写正确的Apple ID和密码后,系统会提示更改登录密码将会创建新的钥匙串,之前的还在,但是仍然需要旧密码才能访问数据,所以喜欢用自动生成的乱码做密码的用户可以着手找回密码了

![2](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/macOSpasswd/2.png)

同意后就可以创建新密码了,点击"重设密码"完成,即可用新鲜出炉的密码登录系统了.

# 用其他管理员账户更改密码

简言之用另一个管理员权限强制更换当前密码,相当于Linux的su用户更改user用户密码,可行性不高,这里略过,有兴趣底下留言

# macOS恢复功能

终于到了大家喜闻乐见的"Command(⌘) + R"了,有点Win RE的味道,这也是第一个复杂到要脱离系统的情况

首先是下面这个界面肯定没错,有人喜欢称其为"恢复模式",有人喜欢叫"实用工具界面",官方说法是"从macOS恢复功能/macOS Recovery启动",总之要从看到"macOS实用工具"窗口开始.

开机同时按住Command(⌘) + R",直到进入以下界面

![3](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/macOSpasswd/3.png)

点击"实用工具 → 终端",就可以调出"终端"窗口.

![4](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/macOSpasswd/4.png)

只需键入

`resetpassword`

中间没有空格,回车即可.

系统处理完会弹出"重设密码向导",首先"选择要重设密码的用户".

之后输入新的密码,点击下一步即可完成,重启后可用新的密码登录.

有一点需要注意,旧版本的系统中,如果登录了Apple ID,创建新密码前需要验证,10.14将这一步骤放到了后面,登录系统后需要重新登录Apple ID.

# 单用户模式

最方便的方法,没有之一

开机同时按住
Command(⌘) + S
读条完成后转到一个DOS界面,这就是类Unix系统的单用户模式了,可以简单理解成超级管理员的专用简易办公室.

![5](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/macOSpasswd/5.png)

上面一大堆都不用管,我们需要先按一下回车,看到
`localhost :/ root#`
就可以输入命令行了.


首先输入

`fsck -y`

"-"前有空格,下同,回车,这行的作用是执行硬盘检测,可以忽略;
待再次出现
`localhost :/ root#`
后接着输入

`mount -uaw /`

"/"前有空格,回车,这一行用来加载文件系统;
输入

`rm /var/db/.AppleSetupDone`

区分大小写,只有第一个"/"前有空格,回车,这行用来删除初始化设置时macOS生成的隐藏文件,删除后再次进入系统就可以新建一个管理员账户,并且完整保留之前账户的数据;
输入

`reboot`

回车重启,命令行会自动执行.

稍作等待就可以看到语言选择界面,放心,数据都还在.

----

其他高级方法可以去看原文,这里不再赘述
