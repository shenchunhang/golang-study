---
title: 验证码设计及mysql max函数
date: 2020-06-07 15:28:51
tags: 
- mysql
description: 哥们我挺无语的, 遇上历史遗留问题了
---

## 起因
事情是这样的, 一线有反馈存在验证码的问题, 频率为偶现, 接到反馈后, 我还有点懵, 有问题就有问题嘛, 还是偶现问题

## 定位
我先自己测试环境试了一下, 连着10个验证码都没有这个问题, 又到生产环境试了10个, 也没有这个问题, 哦嗬, 那没法了, 问题没法复现, 那先看下日志吧, 问了一下一线的同事, 没给啥信息出来, 那就只能代码走读了, 然后就一层一层网上扒, 一阵看下来, 发现这个 查询验证码的sql 问题很大

{% code %}
 select max(id), verifyCode, used, phoneNumber, creteTime ...
 from sms where verifyCode =  limit 1;
{% endcode %}

max函数只对结果集的对应参数列有效, 和其他列无关系, 这个sql查询到的结果
只能查询到verifyCode对应最早的一条数据, 这样的问题居然到我这里才爆出来, 运气挺好
6位的验证码, 也就是这种实现逻辑下, 只有899999个有效值, 哈哈哈, 应该还是当初代码检视太随意了

## 解决方案
### 方案一
直接修改这个sql, max函数一点作用都没有
修改为
{% code %}
 select id, verifyCode, used, phoneNumber, creteTime ...
 from sms where verifyCode =  order by id desc limit 1;
{% endcode %}
这种方式简单, 不需要其他中间件, 改起来也快, 但是不太容易扩展
### 方案二
借用redis, key自带过期时间的, 能查询到就代表有效, 之前的设计里面, 还分为 无效验证码和过期验证码, 真没必要, 纯粹只有无效验证码