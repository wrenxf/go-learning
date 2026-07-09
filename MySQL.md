# MySQL

## DDL-数据库操作

客户端连接:mysql [-h 127.0.0.1] [-P 3306] -u root -p

查询所有数据:show databases;

创建数据库:create database [if not exists] 数据库名 [default charset 字符集] [排序规则]

删除数据库: drop database [if exists] 数据库名