create database library;

use library;

create table users(
      uid int primary key auto_increment, -- 用户id 自增长主键
      uname varchar(20), -- 用户名
      phone_number varchar(20), -- 用户密码
      upassword varchar(20), -- 用户密码

      index index_name(uname), -- 用户名普通索引
      unique index_phone_number(phone_number) -- 唯一索引
)auto_increment=1000;

-- 给用户表的密码字段加上非空约束
alter table users modify upassword varchar(20) not null;

-- 管理员表
create table librarian(
      lid int primary key auto_increment,-- 管理员id 自增长主键
      lname varchar(20),
      lphone_number varchar(20),
      lpassword varchar(20),

      unique index_lphone_number(lphone_number) -- 唯一索引
)auto_increment=2000;

alter table librarian modify lpassword varchar(20) not null;


-- 书库表
create table book(
     bid  int primary key auto_increment,
     bname varchar(20),
     author varchar(20),
     inventory int not null,


     index index_bname(bname),
     index index_author(author)
)auto_increment=3000;

-- 中间借阅表
CREATE TABLE `borrow` (
      `brid` int(11) NOT NULL AUTO_INCREMENT COMMENT '借阅id',
      `uid` int(11) DEFAULT NULL COMMENT '用户id',
      `bid` int(11) DEFAULT NULL COMMENT '图书id',
      `borrow_date` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '借书日期',
      `return_date` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '应该归还日期',
      `real_date` datetime NOT NULL DEFAULT '1000-01-01 00:00:00' COMMENT '实际归还日期',
      PRIMARY KEY (`brid`),
      KEY `uid` (`uid`),
      KEY `bid` (`bid`),
      CONSTRAINT `borrow_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`),
      CONSTRAINT `borrow_ibfk_2` FOREIGN KEY (`bid`) REFERENCES `book` (`bid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
