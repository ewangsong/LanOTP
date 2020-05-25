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
- 安装数据库并创建lanradius数据库

    ```
    create database lanotp default character set utf8;
    ```
- git中默认包含Linux的二进制文件
    ```
    cd /opt
    git clone https://github.com/ewangsong/LanOTP.git
    mv LanOTP lanotp
    cd /opt/lanotp
    ```
- 编写system启动脚本
  
    Web后台UI 
    
    ```
    vim /usr/lib/systemd/system/lanradiusct.service
    
    [Unit]
    Description=java tomcat project
    After=mariadb.service
    
    
    [Service]
    Type=forking
    User=root
    Group=root
    PIDFile=/run/lanradius-radiusct.pid
    ExecStart=/opt/lanradius/lanradius radiusct -d
    ExecReload=
    ExecStop=/opt/lanradius/lanradius stop
    PrivateTmp=true
    
    [Install]
    WantedBy=multi-user.target
    ```
  radius认证服务
    ```
    vim /usr/lib/systemd/system/lanadmin.service
    
    [Unit]
    Description=java tomcat project
    After=mariadb.service
    
    
    [Service]
    Type=forking
    User=root
    Group=root
    PIDFile=/run/lanradius-admin.pid
    ExecStart=/opt/lanradius/lanradius admin -d
    ExecReload=
    ExecStop=/opt/lanradius/lanradius stop
    PrivateTmp=true
    
    [Install]
    WantedBy=multi-user.target
    ```
 - ps：根据需求启动服务认证服务或者Web UI后台
    ```
    systemctl start lanradiusct    //radius认证服务
    systemctl start lanadmin       //Web UI后台管理
    ```
 - 添加开机启动
    ```
    systemctl enable lanradiusct    //radius认证服务
    systemctl enable lanadmin       //Web UI后台管理
    ```
 - 默认管理员账号密码
    ```
    账号：amdin
    密码：admin
    ```
###  命令行参数

```
admin			Start Lanradius             //启动后台管理程序
help        	Help about any command         //显示帮助选项
radiusct    	Start Lanradius                 //启动radius认证程序
stop        	Stop Lanradius                  //停止所有lanradius程序
version     	Print the version number of Lanradius   //显示版本
``` 

### 配置文件详解

```
appname = radiusweb             //app名称   
httpport = 8080                 //后台web端口
runmode = dev                   //变成模式
dbtype = "mysql"                //数据库类型
dbinfo = "root:123456@tcp(192.168.220.138:3306)/test?charset=utf8&loc=Local"                             //数据库配置
radiusport=":1812"               //radius服务端口
```
### 日志文件

```
/var/log/lanraiuds下
```
### 待做
1.	令牌列表用户链接问题
2.	自动提交信息
3.	模糊查询
