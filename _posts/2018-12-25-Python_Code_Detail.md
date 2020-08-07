---
layout:     post                    # 使用的布局(不需要改)
title:      Python代码详解              # 标题
subtitle:   查重这个事太不合逻辑了吧?其他人论文说1+1=2,我写论文也说1+1=2,那我这个就算重复,那我应该怎么办?我说1+1=3?  #副标题
date:       2018-12-25              # 时间
author:     Zen                      # 作者
header-img: img/pet/black.webp   #这篇文章标题背景图片
catalog: False                       # 是否归档
tags:                               #标签
    - 论文
---
fromPython.py
```
# coding=UTF-8

import requests
from bs4 import BeautifulSoup
import codecs
import time
from from_python_to_mysql.connect import mysqlto

DOWNLOAD_URL = 'http://movie.douban.com/top250' //目标网站的URL,今天我们就是要盘...爬它


def download_page(url):
    headers = {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36' //请求头,没人会刻意记住,复制粘贴就好
    }
    proxy = {  //代理设置,我真实的ip已经被封了
        'http': '127.0.0.1:1080',
        'https': '127.0.0.1:1080'
    }
    return requests.get(url, headers=headers).content //得到请求的HTML页面返回给下面


def parse_html(html):
    soup = BeautifulSoup(html, features="html.parser") //美味汤的规范写法
    movie_list_soup = soup.find('ol', attrs={'class': 'grid_view'})
    movie_name_list = []
    title_list = []
    movie_id_list = []
    for movie_li in movie_list_soup.find_all('li'): //循环体内查找每页的电影,过滤规则自己写
        detail = movie_li.find('div', attrs={'class': 'hd'})
        movie_url = detail.find('a')['href']
        movie_id = movie_url[-8:-2]
        time.sleep(1)  # 延时函数,再被封就不止一个小时了
        movie_title, movie_director, movie_date, movie_time, movie_introduction = parse_detail_html(movie_url) //进入电影详情页获取电影详情,导演主演啥的
        time.sleep(1)  # 延时函数,再被封就不止一个小时了
        # if movie_id == '129512': //运行过程中一个自作聪明的电影名字
        mysqlto(movie_id, movie_title, movie_director, movie_date, movie_time, movie_introduction) //将获取到的信息传到另一个方法
        # connect.mysqlto(movie_id, movie_title, 'dsddsds', 'dsddsds', 'dsddsds', 'dsddsds')
        time.sleep(1)  # 延时函数,再被封就不止一个小时了
        title = detail.find('span', attrs={'class': 'title'}).getText()
        print('completely')
        movie_name_list.append(title) //将获取到的电影名添加到列表
        title_list.append(movie_title)
        movie_id_list.append(movie_id)

        time.sleep(1)

    next_page = soup.find('span', attrs={'class': 'next'}).find('a')
    if next_page:
        return movie_name_list, DOWNLOAD_URL + next_page['href']//如果有下一页,递归到下一页
    return movie_name_list, None


def parse_detail_html(movie_url): //用来获取上面每个电影详情页的信息
    html = download_page(movie_url)
    soup = BeautifulSoup(html, features="html.parser")
    movie_title = soup.find('span', attrs={'property': 'v:itemreviewed'}).get_text()
    movie_title = movie_title.replace("'", "\\'")
    time.sleep(1)#延时函数,再被封就不止一个小时了
    movie_director = soup.find('a', attrs={"rel": "v:directedBy"}).get_text()
    time.sleep(1)#延时函数,再被封就不止一个小时了
    movie_introduction = soup.find('span', attrs={"property": "v:summary"}).get_text()
    time.sleep(1)#延时函数,再被封就不止一个小时了
    movie_date = soup.find('span', attrs={"property": "v:initialReleaseDate"}).get_text()
    # 延时函数,再被封就不止一个小时了
    movie_time = soup.find('span', attrs={"property": "v:runtime"}).get_text()
    time.sleep(1)#延时函数,再被封就不止一个小时了

    # connect.mysqlto(movie_id=movie_id, movie_title=movie_title, movie_director=movie_director, movie_date=movie_date,movie_time=movie_time, movie_introduction=movie_introduction)
    return movie_title, movie_director, movie_date, movie_time, movie_introduction


if __name__ == '__main__':
    url = DOWNLOAD_URL

    with codecs.open('movies.txt', 'wb', encoding='utf-8') as fp: //写到文件里,主要做测试,方便看到某个电影是否确实写入
        while url:
            html = download_page(url)
            movies, url = parse_html(html)
            fp.write('\n'.join(movies))
            # fp.write('\n'.join(m_id))
            time.sleep(1)
```
toMysql.py
```
#!/usr/bin/python3

import pymysql
import pymysql as mysql  # pip install PyMySQL Python3的安装方式
# 引入python中的traceback模块,跟踪错误
import traceback


# 引入sys模块


def mysqlto(movie_id, movie_title, movie_director, movie_date, movie_time, movie_introduction):
    # 打开数据库连接
    db = pymysql.connect("localhost", "root", "163453", "douban_movie")

    # 使用cursor()方法获取操作游标
    cursor = db.cursor()

    # SQL 插入语句
    # title = movie_title
    # director=movie_director
    # introduction = movie_introduction
    sql = "INSERT INTO t_doubanmovie(movieID,movieName, \
           movieDirector, movieDate,movieTime,movieIntroduction) \
           VALUES ('%s', '%s', '%s', '%s', '%s',  '%s')" % \
          (movie_id, movie_title, movie_director, movie_date, movie_time, movie_introduction)
    print(movie_id, movie_title, movie_director, movie_date, movie_time, movie_introduction)
    try:
        # 执行sql语句
        cursor.execute(sql)
        # 执行sql语句
        db.commit()
    except:
        # 发生错误时回滚
        db.rollback()
        print('异常')
        traceback.print_exc()
    # 关闭数据库连接
    db.close()


if __name__ == '__main__':
    mysqlto(movie_id='1', movie_title='a', movie_director='f', movie_date='6', movie_time='11', movie_introduction='k')
    mysqlto(movie_id='2', movie_title='b', movie_director='g', movie_date='7', movie_time='12', movie_introduction='l')
    mysqlto(movie_id='3', movie_title='c', movie_director='h', movie_date='8', movie_time='13', movie_introduction='m')
    mysqlto(movie_id='4', movie_title='d', movie_director='i', movie_date='9', movie_time='14', movie_introduction='n')
    mysqlto(movie_id='5', movie_title='e', movie_director='j', movie_date='0', movie_time='15', movie_introduction='o')
```
fromMySql.py
```
import pymysql
import pymysql as mysql  # pip install PyMySQL Python3的安装方式
# 引入python中的traceback模块,跟踪错误
import traceback
from from_mysql_to_elasticsearch.import_elasticsearch import toES


def GetMySQL():
    # 打开数据库连接
    db = pymysql.connect("localhost", "root", "163453", "douban_movie")

    # 使用cursor()方法获取操作游标
    cursor = db.cursor()
    # SQL 查询语句
    sql = "SELECT * FROM t_doubanmovie"
    try:
        # 执行SQL语句
        cursor.execute(sql)
        # 获取所有记录列表
        results = cursor.fetchall()
        for row in results:
            mid = row[0]
            mtitle = row[1]
            mdirector = row[2]
            mdate = row[3]
            mtime = row[4]
            mintroduction = row[5]
            # 打印结果
            print("mid=%s,mtitle=%s,mdirector=%s,mdate=%s,mtime=%s,mintroduction=%s" % \
                  (mid, mtitle, mdirector, mdate, mtime,mintroduction))
            toES(mid, mtitle, mdirector, mdate, mtime,mintroduction)
    except:
        print("Error: unable to fetch data")
    finally:
        # 关闭数据库连接
        db.close()
if __name__ == '__main__':
    GetMySQL()

```
toElasticSearch.py
```
# coding=UTF-8

from datetime import datetime
from elasticsearch import Elasticsearch
es = Elasticsearch()


def toES(mid, mtitle, mdirector, mdate, mtime,mintroduction): //在上一段代码中调用,逐条导入
    doc = [
        {"index": {}},
        {'mid': mid, 'mtitle': mtitle, 'mdirector': mdirector, 'mdate': mdate,'mtime':mtime, 'mdate': mdate,'mintroduction':mintroduction}
    ]
    # doc = [
    #     {'index': {'_index': 'indexName', '_type': 'typeName', '_id': 'idValue'}}
    #     {'name': 'jack', 'sex': 'male', 'age': 10}
    #     {'delete': {'_index': 'indexName', '_type': 'typeName', '_id': 'idValue'}}
    #     {"create": {'_index': 'indexName', "_type": 'typeName', '_id': 'idValue'}}
    #     {'name': 'lucy', 'sex': 'female', 'age': 20}
    #     {'update': {'_index': 'indexName', '_type': 'typeName', '_id': 'idValue'}}
    #     {'doc': {'age': '100'}}
    # ]
    es.bulk(index='movietest', body = doc)
```
