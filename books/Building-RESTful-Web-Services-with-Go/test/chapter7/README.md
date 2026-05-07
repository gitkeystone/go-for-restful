# PostGreSQL Docker
```bash
docker run --name postgres-local -p 5432:5432 \
  -v ~/.postgres-data:/var/lib/postgresql/data  \
  -e POSTGRES_PASSWORD='dq4A@#19' -d postgres:14.22
  
  
# 进入容器中的psql shell
docker exec -it postgres-local psql -U postgres   
\?
```

# pgAdmin 4(图形用户界面-GUI)
> https://hub.docker.com/r/dpage/pgadmin4/

# 用户
\du
```bash
CREATE USER gituser with PASSWORD 'dq4A@#19';
ALTER USER gituser CREATEDB CREATEROLE;
DROP USER gituser;
```

> https://www.postgresql.org/docs/10/sql.html

# 数据库
```postgresql
\? 帮助
CREATE DATABASE mydb;

# 查询表
\c mydb \dt

\c database_name	连接到指定数据库
\dt	列出当前数据库中的所有表
\d table_name	查看表的结构（列、类型、索引等）
\d+ table_name	查看更详细的表信息
\x	切换扩展显示模式（垂直显示行内容）
\q	退出psql命令行
```

# 表
```postgresql
# 创建
CREATE TABLE products (
    product_no integer,
    name text,
    price numeric
);

# 插入  
INSERT INTO products VALUES (1, 'Rice', 5.99);

# 更新
UPDATE products SET price = 10 WHERE price = 5.99;

# 删除
DELETE FROM products WHERE price = 5.99;



```

# 驱动
```bash
go install github.com/lib/pq

```


