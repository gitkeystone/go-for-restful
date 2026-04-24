# 启动
```bash
go run main.go
```

# 测试
```bash
curl -X POST -v -i http://127.0.0.1:8080/users \
-H 'Content-type: application/json' \
-H 'Accept: application/xml' \
-d '{"Id": "1", "Name": "fanguiju"}'

curl -X GET -v -i http://127.0.0.1:8080/users/1
```

# 接口文档
```bash
# 搜索：http://workstation.percxh.com:8080/apidocs.json
http://workstation.percxh.com:8080/apidocs
```