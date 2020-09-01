﻿
<div align=center>
<img src="https://github.com/olongfen/gorm-gin-admin/docs/go/jpeg" width=300" height="300" />
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
[Online Documentation-todo]()

[Development Steps](https://www.gin-vue-admin.com/docs/help) (Contributor:  <a href="https://github.com/LLemonGreen">LLemonGreen</a> And <a href="https://github.com/fkk0509">Fann</a>)
- Web UI Framework：[element-ui](https://github.com/ElemeFE/element)  
- Server Framework：[gin](https://github.com/gin-gonic/gin) 
- Grom Framework: [gorm](https://github.com/go-gorm/gorm)
## 1. Basic Introduction

### 1.1 Project Introduction

[Online Demo]()

username：admin

password：111111

### 1.2 Contributing Guide
#### 1.2.1 Issue Guidelines


#### 1.2.2 Pull Request Guidelines

- Fork this repository to your own account. Do not create branches here.

- Commit info should be formatted as `[File Name]: Info about commit.` (e.g. `README.md: Fix xxx bug`)

- <font color=red>Make sure PRs are created to `develop` branch instead of `master` branch.</font>

- If your PR fixes a bug, please provide a description about the related bug.

- Merging a PR takes two maintainers: one approves the changes after reviewing, and then the other reviews and merges.


## 3. Technical selection

- Frontend: using `Element-UI` based on vue，to code the page.
- Backend: using `Gin` to quickly build basic RESTful API. `Gin` is a web framework written in Go (Golang).
- Database: `PostgreSQL` ，using `gorm2.0` to implement data manipulation.
- Cache: using `Redis` to implement the recording of the JWT token of the currently active user and implement the multi-login restriction.
- API: using Swagger of Gin to auto generate APIs docs。
- Config: using `gopkg.in/yaml.v2` to implement `yaml` config file。
- Log: using `github.com/sirupsen/logrus` record logs。

## 5. Features
- Authority management: Authority management based on `jwt` and `casbin`. 
- User management: The system administrator assigns user roles and role permissions.
- Role management: Create the main object of permission control, and then assign different API permissions and menu permissions to the role.
- Menu management: User dynamic menu configuration implementation, assigning different menus to different roles.
- API management: Different users can call different API permissions.

## 6. To-do list

- [ ] upload & export Excel
- [ ] e-chart
- [ ] workflow, task transfer function
- [ ] frontend independent mode, mock
