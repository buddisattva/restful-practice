# RESTful API Practice

## Deployment
```shell=
$ docker run --name restful-mysql -e MYSQL_ROOT_PASSWORD=5566 -p 3306:3306 -d mysql:8

# 在 MySQL instance 中建立名為 restful 的 database
# 在 restful 中執行 database\migrations\20200921152355_create_albums.up.sql 內容

$ cd /path/to/this-project
$ cd app
$ go run *.go
```

## API

> GET localhost:8000/v1/albums
- 列出全部資源
> GET localhost:8000/v1/albums/{{id}}
- 列出指定 ID 的資源
> POST localhost:8000/v1/albums
```json=
{
    "genre": "yoyo",
    "name": "diy"
}
```
- 新增一筆 genre 為 yoyo、name 為 diy 的資源
> DELETE localhost:8000/v1/albums/{{id}}
- 刪除指定 ID 的資源（冪等行為，刪除一次與多次，都會輸出相同結果）
> PUT localhost:8000/v1/albums/{{id}}
```json=
{
    "genre": "VVVV",
    "name": "5555"
}
```
- 修改指定 ID 的資源屬性（冪等行為）
