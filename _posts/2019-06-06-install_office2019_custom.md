---
layout:     post                    # 使用的布局(不需要改)
title:      定制安装office2019               # 标题
subtitle:    上班吃外卖真的很不健康,珍爱生命,建议不要上班  #副标题
date:       2019-06-06 11:41:00              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 常识
---

代码如下

configure.xml
```
<?xml version="1.0"?>
-<Configuration>
	-<Add  SourcePath="H:\" ForceUpgrade="TRUE" AllowCdnFallback="TRUE" Channel="PerpetualVL2019" OfficeClientEdition="64">
		-<Product ID="ProPlus2019Volume">
			<Language ID="zh-cn"/>
			<Language ID="en-us"/>
			<ExcludeApp ID="Groove"/>
			<ExcludeApp ID="OneNote" />
			<ExcludeApp ID="Access" />
			<ExcludeApp ID="Lync" />
			<ExcludeApp ID="PowerPoint" />
			<ExcludeApp ID="Excel" />
			<ExcludeApp ID="OneDrive" />
			<ExcludeApp ID="Outlook" />
			<ExcludeApp ID="Publisher" />
			<ExcludeApp ID="Word" />
		</Product>
		-<Product ID="VisioPro2019Volume">
			<Language ID="zh-cn"/>
			<Language ID="en-us"/>
			<ExcludeApp ID="Groove"/>
		</Product>
		-<Product ID="ProjectPro2019Volume">
			<Language ID="zh-cn"/>
			<Language ID="en-us"/>
			<ExcludeApp ID="Groove"/>
		</Product>
	</Add>
	<Property Value="0" Name="SharedComputerLicensing"/>
	<Property Value="TRUE" Name="PinIconsToTaskbar"/>
	<Property Value="0" Name="SCLCacheOverride"/>
	<Updates Enabled="TRUE"/>
</Configuration>
```
其中
+ ExcludeApp表示不安装,相反IncludeApp表示安装
+ SourcePath="H:\"表示安装镜像挂载的位置,比如H:\
+ ProPlus2019Volume表示安装你常见的office组件
+ VisioPro2019Volume表示安装VISIO
+ ProjectPro2019Volume表示安装Project
+ 剩下的靠猜也能猜出来
