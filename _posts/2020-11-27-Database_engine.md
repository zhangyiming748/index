---
layout:     post                    # 使用的布局(不需要改)
title:      Mysql 存储引擎的区别和比较             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-11-27 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - database学习
---
#常用存储引擎

## MyISAM存储引擎
MyISAM基于ISAM存储引擎,并对其进行扩展.它是在Web\数据仓储和其他应用环境下最常使用的存储引擎之一.MyISAM拥有较高的插入\查询速度,但不支持事务.
MyISAM主要特性有:
1. 大文件(达到63位文件长度)在支持大文件的文件系统和操作系统上被支持.
2. 当把删除和更新及插入操作混合使用的时候,动态尺寸的行产生更少碎片.这要通过合并相邻被删除的块,以及若下一个块被删除,就扩展到下一块自动完成.
3. 每个MyISAM表最大索引数是64,这可以通过重新编译来改变.每个索引最大的列数是16
4. NULL被允许在索引的列中,这个值占每个键的0~1个字节
5. 可以把数据文件和索引文件放在不同目录(InnoDB是放在一个目录里面的)
MyISAM引擎使用B+Tree作为索引结构,叶节点的data域存放的是数据记录的地址.下图是MyISAM索引的原理图:
![原理图](https://img-blog.csdn.net/20170705170330879?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemdyZ2Zy/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
这里设表一共有三列,假设我们以Col1为主键,则上图是一个MyISAM表的主索引(Primary key)示意.可以看出MyISAM的索引文件仅仅保存数据记录的地址.在MyISAM中,主索引和辅助索引(Secondary key)在结构上没有任何区别,只是主索引要求key是唯一的,而辅助索引的key可以重复.如果我们在Col2上建立一个辅助索引,则此索引的结构如下图所示:
![结构](https://img-blog.csdn.net/20170705170516932?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemdyZ2Zy/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
同样也是一颗B+Tree,data域保存数据记录的地址.因此,MyISAM中索引检索的算法为首先按照B+Tree搜索算法搜索索引,如果指定的Key存在,则取出其data域的值,然后以data域的值为地址,读取相应数据记录.
MyISAM的索引方式也叫做"非聚集"的,之所以这么称呼是为了与InnoDB的聚集索引区分.

## InnoDB存储引擎
InnoDB是事务型数据库的首选引擎,支持事务安全表(ACID),支持行锁定和外键,上图也看到了,InnoDB是默认的MySQL引擎.
InnoDB主要特性有:
1. InnoDB给MySQL提供了具有提交\回滚和崩溃恢复能力的事物安全(ACID兼容)存储引擎.InnoDB锁定在行级并且也在SELECT语句中提供一个类似Oracle的非锁定读.这些功能增加了多用户部署和性能.在SQL查询中,可以自由地将InnoDB类型的表和其他MySQL的表类型混合起来,甚至在同一个查询中也可以混合
2. InnoDB是为处理巨大数据量的最大性能设计.它的CPU效率可能是任何其他基于磁盘的关系型数据库引擎锁不能匹敌的
3. InnoDB存储引擎完全与MySQL服务器整合,InnoDB存储引擎为在主内存中缓存数据和索引而维持它自己的缓冲池.InnoDB将它的表和索引在一个逻辑表空间中,表空间可以包含数个文件(或原始磁盘文件).这与MyISAM表不同,比如在MyISAM表中每个表被存放在分离的文件中.InnoDB表可以是任何尺寸,即使在文件尺寸被限制为2GB的操作系统上
4. InnoDB支持外键完整性约束,存储表中的数据时,每张表的存储都按主键顺序存放,如果没有显示在表定义时指定主键,InnoDB会为每一行生成一个6字节的ROWID,并以此作为主键.
虽然InnoDB也使用B+Tree作为索引结构,但具体实现方式却与MyISAM截然不同.

第一个重大区别是InnoDB的数据文件本身就是索引文件.从 上文知道,MyISAM索引文件和数据文件是分离的,索引文件仅保存数据记录的地址.而在InnoDB中,表数据文件本身就是按B+Tree组织的一个索 引结构,这棵树的叶节点data域保存了完整的数据记录.这个索引的key是数据表的主键,因此InnoDB表数据文件本身就是主索引.
![主索引](https://img-blog.csdn.net/20170705170833096?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemdyZ2Zy/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
上图是InnoDB主索引(同时也是数据文件)的示意图,可以看到叶节点包含了完整的数据记录.这种索引叫做聚集索引.因为InnoDB的数据文件本身 要按主键聚集,所以InnoDB要求表必须有主键(MyISAM可以没有),如果没有显式指定,则MySQL系统会自动选择一个可以唯一标识数据记录的列 作为主键,如果不存在这种列,则MySQL自动为InnoDB表生成一个隐含字段作为主键,这个字段长度为6个字节,类型为长整形.
第二个与MyISAM索引的不同是InnoDB的辅助索引data域存储相应记录主键的值而不是地址.换句话说,InnoDB的所有辅助索引都引用主键作为data域.例如,下图为定义在Col3上的一个辅助索引:
![辅助索引](https://img-blog.csdn.net/20170705171044159?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemdyZ2Zy/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
这里以英文字符的ASCII码作为比较准则.聚集索引这种实现方式使得按主键的搜索十分高效,但是辅助索引搜索需要检索两遍索引:首先检索辅助索引获得主键,然后用主键到主索引中检索获得记录.
了解不同存储引擎的索引实现方式对于正确使用和优化索引都非常有帮助,例如知道了InnoDB的索引实现后,就很容易明白为什么不建议使用过长的字段作为 主键,因为所有辅助索引都引用主索引,过长的主索引会令辅助索引变得过大.再例如,用非单调的字段作为主键在InnoDB中不是个好主意,因为 InnoDB数据文件本身是一颗B+Tree,非单调的主键会造成在插入新记录时数据文件为了维持B+Tree的特性而频繁的分裂调整,十分低效,而使用 自增字段作为主键则是一个很好的选择.

## MEMORY存储引擎
MEMORY存储引擎将表中的数据存储到内存中,未查询和引用其他表数据提供快速访问.
MEMORY主要特性有:
1. MEMORY表的每个表可以有多达32个索引,每个索引16列,以及500字节的最大键长度
2. MEMORY存储引擎执行HASH和BTREE缩影
3. 可以在一个MEMORY表中有非唯一键值
4. MEMORY表使用一个固定的记录长度格式
5. MEMORY不支持BLOB或TEXT列
6. MEMORY支持AUTO_INCREMENT列和对可包含NULL值的列的索引
7. MEMORY表在所由客户端之间共享(就像其他任何非TEMPORARY表)
8. MEMORY表内存被存储在内存中,内存是MEMORY表和服务器在查询处理时的空闲中,创建的内部表共享
9. 当不再需要MEMORY表的内容时,要释放被MEMORY表使用的内存,应该执行DELETE FROM或TRUNCATE TABLE,或者删除整个表(使用DROP TABLE)

## Archive存储引擎
不同的存储引擎都有各自的特点,以适应不同的需求,如下表所示:
![](https://img-blog.csdn.net/20170705172036010?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvemdyZ2Zy/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

|功能|MYISAM|Memory|InnoDB|Archive|
|:--:|:---:|:---:|:---:|:---:|
|存储限制|256TB|RAM|64TB|None|
|支持事务|No|No||No|
|支持全文索引|Yes|No|No|No|
|支持数索引|Yes|Yes|Yes|No|
|支持哈希索引|No|Yes|No|No|
|支持数据缓存|No|N/A|Yes|No|
|支持外键|No|No|Yes|No|

# 存储引擎的选择
InnoDB :如果要提供提交\回滚\崩溃恢复能力的事务安全(ACID兼容)能力,并要求实现并发控制,InnoDB是一个好的选择


MyISAM:如果数据表主要用来插入和查询记录,则MyISAM(但是不支持事务)引擎能提供较高的处理效率

Memory:如果只是临时存放数据,数据量不大,并且不需要较高的数据安全性,可以选择将数据保存在内存中的Memory引擎,MySQL中使用该引擎作为临时表,存放查询的中间结果.数据的处理速度很快但是安全性不高.

Archive:如果只有INSERT和SELECT操作,可以选择Archive,Archive支持高并发的插入操作,但是本身不是事务安全的.Archive非常适合存储归档数据,如记录日志信息可以使用Archive

使用哪一种引擎需要灵活选择,一个数据库中多个表可以使用不同引擎以满足各种性能和实际需求,使用合适的存储引擎,将会提高整个数据库的性能

## InnoDB 和 MyISAM之间的区别:
1. InnoDB支持事物,而MyISAM不支持事物

2. InnoDB支持行级锁,而MyISAM支持表级锁

3. InnoDB支持MVCC, 而MyISAM不支持

4. InnoDB支持外键,而MyISAM不支持

5. InnoDB不支持全文索引,而MyISAM支持.(X)
