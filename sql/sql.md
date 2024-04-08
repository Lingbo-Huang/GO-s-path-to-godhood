```sql
SELECT * FROM students ORDER BY score DESC;
SELECT age, AVG(score) AS avg_score FROM students GROUP BY age;
SELECT age, AVG(score) AS avg_score FROM students GROUP BY age WITH ROLLUP;
SELECT * FROM students ORDER BY score ASC;
SELECT age, AVG(score) AS avg_score FROM students GROUP BY age;
```


```sql
CREATE TABLE `user` (
`id` int(11) NOT NULL DEFAULT '0',
`username` varchar(255) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=myisam DEFAULT CHARSET=utf8;

CREATE TABLE `user2` (
`id` int(11) NOT NULL DEFAULT '0',
`username` varchar(255) NOT NULL,
PRIMARY KEY(`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8
```

```sql
select id, name from employee;
CREATE TABLE `employee` (
`id` int(11) NOT NULL,
`name` varchar(255) DEFAULT NULL,
`age` int(11) DEFAULT NULL,
`date` datetime DEFAULT NULL,
`sex` int(1) DEFAULT NULL,
PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

select id, name from employee where name='jay' limit 1;
select id, name from employee where id > 10000 limit 10
select id, name from employee order by id limit 10000, 10
select userId, name from user where userId like '123%';
```

```sql
select * from (select * from tab1 where id > 2) t1 left join tab2 t2 on t1.size = t2.size;
alter table user add index idx_address_age (address, age)
```

```sql
insert into user(name, age) values 
<foreach collection="list" item="item" index="index" separator=",">
          (#{item.name}, #{item.age})                      
</foreach>
```

```sql
select DISTINCT name from user;
```