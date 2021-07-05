---
layout:     post                    # 使用的布局(不需要改)
title:      一键设置Linux/Macos代理      # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2021-07-05 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - Proxy
---



```
#!/bin/bash
hostip=$(cat /etc/resolv.conf | grep nameserver | awk '{ print $2 }')
wslip=$(hostname -I | awk '{ print $1 }')
port=8889
socks=1089
PROXY_HTTP="http://${hostip}:${port}"

set_proxy(){
    export http_proxy="${PROXY_HTTP}"
    export HTTP_PROXY="${PROXY_HTTP}"

    export https_proxy="${PROXY_HTTP}"
    export HTTPS_proxy="${PROXY_HTTP}"

    git config --global http.proxy "${PROXY_HTTP}"
    git config --global https.proxy "${PROXY_HTTP}"

    echo "Host github.com" > ~/.ssh/config
    echo "HostName github.com" >> ~/.ssh/config
    echo "User git" >> ~/.ssh/config
    echo "ProxyCommand nc -v -x ${hostip}:${socks} %h %p" >> ~/.ssh/config
}

unset_proxy(){
    unset http_proxy
    unset HTTP_PROXY

    unset https_proxy
    unset HTTPS_PROXY

    git config --global --unset http.proxy
    git config --global --unset https.proxy
    
    echo "" > ~/.ssh/config
}

test_setting(){
    echo "Host ip:" ${hostip}
    echo "WSL ip:" ${wslip}
    echo "Current proxy:" $https_proxy
}

if [ "$1" = "set" ]
then
    set_proxy

elif [ "$1" = "unset" ]
then
    unset_proxy

elif [ "$1" = "test" ]
then
    test_setting
else
    echo "Unsupported arguments(set/unset/test)"
fi
```
