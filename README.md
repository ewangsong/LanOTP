## 关于
- LANOTP是一个开源的Radius OTP服务软件。

- LANOTP支持标准RADIUS协议。

- 编写参考：ToughRADIUS、SoftRadius、GO-OTP

## 安装（目前仅支持Linux）

---
### 编译安装
 编译需要go环境需自己安装下载

- go get github.com/astaxie/beego
- go get -u github.com/go-sql-driver/mysql   
    ```
    在go源目录下创建文件夹用来git clone 源代码
    mkdir lanotp
    cd lanotp
    git clone https://github.com/ewangsong/LanOTP.git
    cd LanOTP
    go build -o lanotp main.go
    cp -r LanOTP /opt/lanotp
    ```
### 快捷安装
- 安装数据库并创建lanotp数据库

    ```
    create database lanotp default character set utf8;
    ```
 - 默认管理员账号密码
    ```
    账号：amdin
    密码：admin
    ```
###  命令行参数
```
-start
    	前台启动
  -startd
    	后台启动
  -stop
    	关闭程序
  -v	查看版本
``` 

### 配置文件详解

```
appname = lanotp                //app名称   
httpport = 9091                 //后台web端口
runmode = dev                   //变成模式
dbtype = "mysql"                //数据库类型
dbinfo = "root:123456@tcp(192.168.220.138:3306)/test?charset=utf8&loc=Local"                             //数据库配置
radiusport=":1819"               //radius服务端口
```
### 日志文件
```
/var/log/lanotp下
```
### 待做
1.	令牌列表用户链接问题
2.	自动提交信息
3.	模糊查询
