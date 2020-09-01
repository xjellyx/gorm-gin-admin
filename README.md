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

English | [简体中文](./README-zh_CN.md)

# Project Guidelines

- Web UI Framework：[element-ui](https://github.com/ElemeFE/element)  
- Server Framework：[gin](https://github.com/gin-gonic/gin) 
- Grom Framework: [gorm](https://github.com/go-gorm/gorm)
## 1. Basic Introduction
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
