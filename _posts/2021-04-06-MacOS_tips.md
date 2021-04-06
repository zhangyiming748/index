---
layout:     post                    # 使用的布局(不需要改)
title:      Catalina下关闭Big Sur更新提示             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-12-25 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - macOS
---

# 批量删除macOS读取外部驱动器产生的4k挂载点文件
`find yourdir/ -size 4k -exec rm -f {} \;`
