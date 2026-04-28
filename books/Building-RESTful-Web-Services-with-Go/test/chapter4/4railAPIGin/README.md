# API 设计(CRUD)
HTTP动词      路径                 操作        资源
POST        /v1/trains           Create      Train
POST        /v1/stations         Create      Station
GET         /v1/trains/id        Read        Train
GET         /v1/stations/id      Read        Station
POST        /v1/schedule         Create      Route
DELETE      /v1/trains/id        Delete      Train
DELETE      /v1/stations/id      Delete        Station

# 引入同级模块
```bash
cd /path/to/railAPI
go mod edit --replace dbutils=../2rail/dbutils
go get dbutils
```
