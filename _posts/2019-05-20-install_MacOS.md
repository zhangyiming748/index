---
layout:     post                    # 使用的布局(不需要改)
title:      制作启动盘安装Macos              # 标题
subtitle:   浔阳江头夜送客,硫氢化铁显红色  #副标题
date:       2019-05-20              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - macOS
---


![1](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/1.webp)

这种大版本升级,我总是喜欢干净的全新安装,所以当然是使用USB启动盘来安装全新的系统,当初的系统还是可以直接下载系统制作U盘的,时过境迁,苹果已经不提供单独的镜像下载了,所以这次就来个一劳永逸的方法,老少通吃!
事前准备:一个大于8G的u盘
开始制作(以macOS 10.14 Mojave为例)
+ 需要把u盘格式化为`Mac OS X 扩展（日志式）` `GUID 分区图``
1. 下載 macOS High Sierra:
![2](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/2.webp)
2. 下载完成macOS Mojave后按下command(⌘)+ Q关闭跳出来的安装信息
![3](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/3.webp)
3. 打开终端(路径:应用程序->工具->terminal)
![4](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/4.webp)
4. 在终端中输入下列命令,并加上空格
![5](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/5.webp)
5. 在Finder->应用程序中找到刚下载的[安装macOS Mojave],并按下右键选择[显示包内容]
![6](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/6.webp)
6. 进入 Contents > Resources,找到 createinstallmedia.
![7](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/7.webp)
7. 將 第四步 中的 createinstallmedia 拖到 Terminal 窗口中
![8](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/8.webp)
![9](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/9.webp)
8. 输入下列命令,并加上空格
![10](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/10.webp)
9. 將 U盘插上电脑,并直接拖拽至Terminal中 (假设U盘卷标为USB)
![11](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/11.webp)
![12](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/12.webp)
10. 在Terminal中输入下列命令,并加上空格
![13](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/13.webp)
11. 打开Finder-> 应用程序,将第一步中的[安裝 macOS Mojave]拖动到Terminal中
![14](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/14.webp)
![15](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/15.webp)
12. 按下Enter并输入本机密码
![16](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/16.webp)
13. 密码输入完成后,按下'y' 开始制作
![17](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/17.webp)
14. 开始制作时会初始化这个U盘,所以注意备份U盘资料,制作过程大约20分钟(DD方式写入),具体时间取决于U盘体制
![18](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/18.webp)
15. 制作完成,就可以关闭Terminal
![19](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/19.webp)
16. 即可看见桌面上出现了Install macOS Mojave的系统盘
![20](https://raw.githubusercontent.com/zhangyiming748/zhangyiming748.github.io/master/img/installMacOS/20.webp)

#### FAQ

###### Mojave 安装包太小无法制作?
正常情况下AppStore下载来的安装包大约6G,如果发现过小,请先执行下载回来的应用程序,然后真正的系统DMG才会开始下载,然后再从第四部开始制作
###### 懒人模式
如果你觉得上面的方法太复杂搞不懂,那就直接使用懒人安装吧
请先把要制作的U盘命名为USB,之后开启Terminal,把下列的粘贴上去,再按下 return,接下來照着 Step 12. 继续下去即可.

```
sudo /Applications/Install\ macOS\ Mojave.app/Contents/Resources/createinstallmedia --volume /Volumes/USB --applicationpath /Applications/Install\ macOS\ Mojave.app
```
### 后记
接下来只要在插上U盘的MBP上重启同时按住option听到Ding~就可以开始安装系统了
每次大版本更新我都会全新安装避免日后因为升级安装出现一些难以理解的怪现象
