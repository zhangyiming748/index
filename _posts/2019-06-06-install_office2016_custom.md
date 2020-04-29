---
layout:     post                    # 使用的布局(不需要改)
title:      定制安装office2016               # 标题
subtitle:    熬夜真的很伤身体,所以建议通宵  #副标题
date:       2019-06-06 11:42:00              # 时间
author:     Zen                      # 作者
header-img: img/pet/supremelysab-787607-unsplash.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 常识
---

代码如下

configure.xml
```
<Configuration>
	<Add SourcePath="H:\" OfficeClientEdition="64" Branch="Current">
		<Product ID="O365ProPlusRetail">
			<Language ID="zh-CN" />
			<ExcludeApp ID="Access" />
			<ExcludeApp ID="Groove" />
			<ExcludeApp ID="InfoPath" />
			<ExcludeApp ID="Lync" />
			<ExcludeApp ID="Publisher" />
			<ExcludeApp ID="SharePointDesigner" />
			<ExcludeApp ID="OneNote" />
			<ExcludeApp ID="Outlook" />
		</Product>
		<Product ID="VisioProRetail">
			<Language ID="zh-CN" />
		</Product>
		<Product ID="proplusretail">
			<Language ID="zh-cn" />
			<ExcludeApp ID="Access" />
			<ExcludeApp ID="Groove" />
			<ExcludeApp ID="InfoPath" />
			<ExcludeApp ID="Lync" />
			<ExcludeApp ID="Publisher" />
			<ExcludeApp ID="SharePointDesigner" />
			<ExcludeApp ID="OneNote" />
			<ExcludeApp ID="Outlook" />
		</Product>
	</Add>
</Configuration>
```
其中
+ ExcludeApp表示不安装,相反IncludeApp表示安装
+ SourcePath="H:\"表示安装镜像挂载的位置,比如H:\
+ ProPlus2019Volume表示安装你常见的office组件
+ VisioPro2019Volume表示安装VISIO
+ 剩下的靠猜也能猜出来
