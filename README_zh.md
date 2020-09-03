﻿
<div align=center>
<img src="https://github.com/olongfen/gorm-gin-admin/blob/master/docs/go.jpeg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.14-blue"/>
<img src="https://img.shields.io/badge/gin-1.6.3-lightBlue"/>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.12.0-green"/>
<img src="https://img.shields.io/badge/gorm-1.20.0-red"/>
<img src="https://img.shields.io/badge/casbin-2.11.2-yellow"/>
<img src="https://img.shields.io/badge/redis-6.15.9-lightGree"/>
</div>

简体中文 | [English](./README.md)

# 项目指导文档
[Online Demo](http://39.98.44.155:80)

[Font Een](https://github.com/olongfen/user_admin.git)
- Web UI Framework：[element-ui](https://github.com/ElemeFE/element)  
- Server Framework：[gin](https://github.com/gin-gonic/gin) 
- Grom Framework: [gorm](https://github.com/go-gorm/gorm)
## 1. 基本介绍
### 1.1 项目结构
```
    │  ├─conf               (Config file)
    │  ├─docs  	            （swagger APIs docs）
    │  ├─log                 (log file)
    │  ├─public              (public static file)
            │  ├─static      (head icon)
    ├─src
        │  ├─controller     (Controller)
        │  ├─middleware      (Middleware）
        │  ├─models         （Model entity）
        │  ├─pkg            （Project private package）
            │  ├─adapter    (Casbin adapter)
            │  ├─app        (Gin service response) 
            │  ├─codes      (Response code)
            │  ├─error      (Project private error)
            │  ├─gredis     (Redis)
            │  ├─query      (Songo parase to SQL line)
            │  ├─setting    (Project setting)
        │  ├─router         (Router)
        │  ├─rpc            (RPC)
        │  ├─service        (services)
        │  └─utils	        （common utilities）
    
```
#### 1.1.2 生成API文档
```
    cd grom-gin-admin
    go get -u github.com/swaggo/swag/cmd/swag
    swag init
```
访问http://localhost:8050/swagger/index.html，即可查看swagger文档

### 1.2 环境配置
#### 1.2.1 Golang 安装
- 国外代理可以浏览该界面查看golang安装文档  [olongfen.github.o](https://olongfen.github.io/#/note/fedora%E8%A3%85%E6%9C%BA%E5%90%8E%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE?id=%e5%ae%89%e8%a3%85golang)
- 国内朋友点击这个界面查看golang安装文档 [blog.olongfen.ltd](http://blog.olongfen.ltd:9001/#/note/fedora%E8%A3%85%E6%9C%BA%E5%90%8E%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE?id=%E5%AE%89%E8%A3%85golang)

### 1.2.2 项目环境配置
本人在这个项目的环境配置通过docker安装，需要在主机安装配置自己安装PostgreSQL和Redis即可
``` 
   git clone https://github.com/olongfen/gorm-gin-admin.git
```
```
    cd gorm-gin-admin
    docker-compose up -d .
```

### 1.2.3 项目配置说明

- 通过RSA 密钥对来验证用户会话权限
   ```
    admin.key admin.pub
    user.key user.pub
   ```
- casbin 模型文件
  ```model_casbin.conf``` 
- 项目配置文件 
    
 ``` 
    当你的conf目录下还没有配置文件的时候，你可以运行一次项目，会自动生成配置文件，再运行一次就可以开机服务了，修改项目配置无需重启项目，配置热加
载每10s一次，可以自己在pkg/settiong/下修改热加载时间   
  ```  

### 1.2.4 Run Service
``` 
     go run main.go
```

## 2. 技术栈

- 前端：用基于`vue`的`Element-UI`构建基础页面。
- 后端：用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架。
- 数据库：采用`PostgreSQL`，使用`gorm2.0版本`实现对数据库的基本操作,
- 缓存：准备开发使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用`gopkg.in/yaml.v2`实现`yaml`格式的配置文件。
- 日志：使用`github.com/sirupsen/logrus`实现日志记录。

## 3. 主要功能
- 权限管理：基于`jwt`和`casbin`实现的权限管理 
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 条件搜索：增加条件搜索示例。
- restful示例：可以参考用户管理模块中的示例API。 

## 4. 计划任务

- [ ] 导入，导出Excel
- [ ] 管理员操作记录
- [ ] RPC给其他项目调用 
- [ ] token缓存机制
