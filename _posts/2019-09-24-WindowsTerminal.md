---
layout:     post                    # 使用的布局(不需要改)
title:     WindowsTerminal参数详解              # 标题
subtitle: 找个有西瓜和西红柿的菜市场,告诉他,去买一个西瓜,如果看见西红柿的话,就买两个 要是买回来一个西瓜和两个西红柿,他就不是真正的程序员 真正的程序员会买回来两个西瓜  #副标题
date:       2019-09-24 00:00:00 GMT+0800               # 时间
author:     Zen                      # 作者
header-img: img/photo/jiuzhaigou.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 常识
---


WindowsTerminal的配置文件在

----

`%userprofile%\AppData\Local\Packages\Microsoft.WindowsTerminal_8wekyb3d8bbwe\RoamingState\profiles.json`

----

```

{
    "globals" :
    {
        "alwaysShowTabs" : true,//是否永远显示上面的标签栏
        "defaultProfile" : "{08b5e85b-3615-463d-8df1-70bc4703843f}",//默认启动的配置文件,这是一个GUID,对应着下面的配置文件,这里填着哪个配置的GUID,Terminal启动时就会默认执行哪一个配置文件
        "initialCols" : 120,//表明启动时默认的终端是多长
        "initialRows" : 30,//表明启动时默认的终端是多宽
        "keybindings" : //终端的快捷键
        [
        	...
        ],
        "requestedTheme" : "system",
        "showTabsInTitlebar" : true,//标题栏显示标签
        "showTerminalTitleInTitlebar" : true,//在标题栏显示标签文字
        "wordDelimiters" : " ./\\()\"'-:,.;<>~!@#$%^&*|+=[]{}~?\u2502"
    },
    "profiles" :
    [
        {
            "acrylicOpacity" : 0.75,//不透明度
            "closeOnExit" : true,
            "colorScheme" : "Solarized Light",//颜色主题
            "commandline" : "powershell.exe",//对应启动的shell,这里powershell.exe表明启动的是PowerShell,当然你也可以换成pwsh,cmd甚至是VS的开发命令行\ssh直接连接远程主机或者虚拟机.只要是命令行的程序,这里都可以填入,但是要注意一下json的语法.这里要在多说一句,这里填入wsl命令就可以启动默认的Linux子系统,wsl ~就可以让启动目录是wsl的用户主目录
            "cursorColor" : "#657B83",//光标颜色
            "cursorShape" : "bar",//光标形状 bar==竖线 emptyBox==空心矩形 filledBox==实心矩形 vintage==下划线
            "fontFace" : "Sarasa Term SC",//就是字体了
            "fontSize" : 12,//字体大小
            "guid" : "{08b5e85b-3615-463d-8df1-70bc4703843f}",//guid就是这个配置的GUID,对于每一个配置,都应该使用不同的GUID,否则Terminal就会发生错误,GUID可以使用Powershell访问CLR生成,即在PowerShell中打出以下指令就可以获得一个新的GUID:[System.Guid]::NewGuid().ToString()
            "historySize" : 9001,
            "icon" : "ms-appdata:///roaming/powershell_32px.png",//是这个配置的图标,也就是在标题栏和新建中显示的图标,ms-appdata:///roaming/路径就是json储存的路径
            "name" : "PowerShell",//就是配置的名称,在新建菜单中显示的名称
            "padding" : "0, 0, 0, 0",
            "snapOnInput" : true,
            "startingDirectory" : "%USERPROFILE%",//启动目录
            "useAcrylic" : true,//毛玻璃效果
            "tabTitle" : "PowerShell"//标题栏文字
        },
        {
            "acrylicOpacity" : 0.75,
            "closeOnExit" : true,
            "colorScheme" : "Solarized Light",
            "commandline" : "pwsh",
            "cursorColor" : "#657B83",
            "cursorShape" : "bar",
            "fontFace" : "Sarasa Term SC",
            "fontSize" : 12,
            "guid" : "{f1c1bd17-a4e8-4497-acc2-4245b3b6bbd3}",
            "historySize" : 9001,
            "icon" : "ms-appdata:///roaming/powershell_32px.png",
            "name" : "PowerShell Core",
            "padding" : "0, 0, 0, 0",
            "snapOnInput" : true,
            "startingDirectory" : "%USERPROFILE%",
            "useAcrylic" : true,
            "tabTitle" : "PowerShell Core"
        },
        {
            "acrylicOpacity" : 0.85000002384185791,
            "closeOnExit" : false,
            "colorScheme" : "3024 Day",
            "commandline" : "wsl ~",
            "cursorColor" : "#4A4543",
            "cursorShape" : "bar",
            "fontFace" : "Sarasa Term SC",
            "fontSize" : 12,
            "guid" : "{78e390db-1bff-4533-9d7c-20f53d8bafa1}",
            "historySize" : 9001,
            "icon" : "ms-appdata:///roaming/ubuntu_32px.png",
            "name" : "Default WSL",
            "padding" : "0, 0, 0, 0",
            "snapOnInput" : true,
            "useAcrylic" : true
        },
        {
            "acrylicOpacity" : 0.85000002384185791,
            "closeOnExit" : false,
            "colorScheme" : "Night Owlish Light",
            "commandline" : "ssh mahoshojohcg@192.168.136.128",
            "cursorColor" : "#403F53",
            "cursorShape" : "bar",
            "fontFace" : "Sarasa Term SC",
            "fontSize" : 12,
            "guid" : "{50dc2393-bd60-4cbd-bd8f-64a7602e86ea}",
            "historySize" : 9001,
            "icon" : "ms-appdata:///roaming/debian.png",
            "name" : "Debian VM",
            "padding" : "0, 0, 0, 0",
            "snapOnInput" : true,
            "useAcrylic" : true
        }
    ],
    "schemes" : //配色方案
    [
    	...
    ]
}

```
