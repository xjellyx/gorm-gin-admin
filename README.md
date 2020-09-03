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

English | [简体中文](./README_zh.md)

# Project Guidelines
[Online Demo](http://39.98.44.155:80)

[Font Een](https://github.com/olongfen/user_admin.git)
- Web UI Framework：[element-ui](https://github.com/ElemeFE/element)  
- Server Framework：[gin](https://github.com/gin-gonic/gin) 
- Grom Framework: [gorm](https://github.com/go-gorm/gorm)
## 1. Basic Introduction
### 1.1 Project structure
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
#### 1.1.1 Generate  api docs 
```
    cd grom-gin-admin
    go get -u github.com/swaggo/swag/cmd/swag
    swag init
```

### 1.2 Environment step
#### 1.2.1 Install golang
- If you have proxy to see  [olongfen.github.o](https://olongfen.github.io/#/note/fedora%E8%A3%85%E6%9C%BA%E5%90%8E%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE?id=%e5%ae%89%e8%a3%85golang)
- Else to see [blog.olongfen.ltd](http://blog.olongfen.ltd:9001/#/note/fedora%E8%A3%85%E6%9C%BA%E5%90%8E%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE?id=%E5%AE%89%E8%A3%85golang)

### 1.2.2 Install project Environment
In this project, my environment install by docker.
``` 
   git clone https://github.com/olongfen/gorm-gin-admin.git
```
```
    cd gorm-gin-admin
    docker-compose up -d .
```

### 1.2.3 Project configure step

- admin and general user rsa key
   ```
    admin.key admin.pub
    user.key user.pub
   ```
- casbin model
  ```model_casbin.conf``` 
- project config 
    
 ``` 
  when you run project the first that will be auto creating, and then you can run service, you can edit project config file 
when project is running.     
  ```  

### 1.2.4 Run Service
``` 
    if you don't have config file which name project.config.yaml. you shuold run  project first auto generate config file,
then you try again can be running.
    command: go run main.go
```

## 2. Technical selection

- Frontend: using `Element-UI` based on vue，to code the page.
- Backend: using `Gin` to quickly build basic RESTful API. `Gin` is a web framework written in Go (Golang).
- Database: `PostgreSQL` ，using `gorm2.0` to implement data manipulation.
- Cache: using `Redis` to implement the recording of the JWT token of the currently active user and implement the multi-login restriction.
- API: using Swagger of Gin to auto generate APIs docs。
- Config: using `gopkg.in/yaml.v2` to implement `yaml` config file。
- Log: using `github.com/sirupsen/logrus` record logs。

## 3. Features
- Authority management: Authority management based on `jwt` and `casbin`. 
- User management: The system administrator assigns user roles and role permissions.
- Role management: Create the main object of permission control, and then assign different API permissions to the role.
- Menu management: User dynamic menu configuration implementation, assigning different menus to different roles.
- API management: Different users can call different API permissions.

## 4. To-do list

- [ ] upload & export Excel
- [ ] record manager actions
- [ ] RPC 
- [ ] Cache token
