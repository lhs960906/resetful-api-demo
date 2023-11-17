# Demo 后端



Go http 标准库

第三方路由库 httprouter

Go 操作 MySQL



## 初始化项目

```shell
git clone git@github.com:lhs960906/resetful-api-demo.git -o demo
```



## 打印日志



## 项目配置管理

操作 MySQL 的时候，不能在程序中直接写死 MySQL 相关信息，需要用配置文件去定义连接 MySQL 的相关信息。

基于配置文件、基于环境（对容器友好）变量两种。

配置文件格式解析：

* etc/demo.json
* etc/demo.yaml
* etc/demo.xml
* etc/demo.toml（采用）。学习对应[toml解析库](https://github.com/BurntSushi/toml)。

基于配置文件的解析：

* etc/demo.env（采用）。学习对应[环境变量接续库](https://github.com/caarlos0/env)。



配置文件方式：/etc/demo.toml

```shell
[app]
name = "demo"
host = "0.0.0.0"
port = "8050"
key  = "this is your app key"

[mysql]
host = "127.0.0.1"
port = "3306"
username = "go_course"
password = "xxxx"
database = "go_course"

[log]
level = "debug"
dir = "logs"
format = "text"
to = "stdout"
```

基于环境变量：etc/demo.env

```shell
export MYSQL_HOST="127.0.0.1"
export MYSQL_PORT="3306"
export MYSQL_USERNAME="go_course"
export MYSQL_PASSWORD="xxx"
export MYSQL_DATABASE="go_course"
```



# 参考

[CMDB Demo](gitee.com/infraboard/go-course/blob)

[demo-api](https://gitee.com/infraboard/go-course/blob/master/day14/demo-api.md)



# 补充
iota
项目框架
日志框架
配置框架
assert包
连接数据库