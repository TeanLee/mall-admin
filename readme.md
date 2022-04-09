# 商城类微信小程序的后台管理平台
对应的商城类微信小程序实现：https://github.com/TeanLee/hema
## mall-admin-front
商城类管理平台的前端

技术栈：Vue + Element UI + Sass + HTML + JS +CSS 

## mall-admin-server-go
商城类管理平台的后端

技术栈：Go + Gin + GORM + MySQL + Redis

## 项目简介
该项目实现了微信商城管理平台的前后端。

详细功能介绍及演示地址：https://juejin.cn/post/7084021257041084430

实现的功能点：
- 权限管理
    - [ ] 管理员注册
    - [ ] 管理员登录
    - [ ] 用户权限修改（普通管理员、超级管理员）
- 商品管理
    - [ ] 商品列表（修改、删除）
    - [ ] 添加商品
    - [ ] 商品统计（图表统计）
- 分类管理
    - [ ] 商品分类列表（修改、删除）
- 商城普通用户管理
    - [ ] 用户管理（用户信息查看、修改）
    - [ ] 添加用户


<br>

## 运行方式
1. 在本机中安装 [Docker](https://www.docker.com/)
2. 在本文件夹下执行 `docker-compose up --build`

<br>

**前端技术**
| 技术 | 说明 | 官网 |
|  ----  | ----  | ----  |
| JavaScript | JavaScript语言 | https://www.javascript.com/ |
| Vue | 前端框架 | https://vuejs.org/ |
| Axios | 前端HTTP框架 | <https://github.com/axios/axios> |
| ElementUI | UI组件库 | https://element.eleme.cn/#/zh-CN |
| Sass | 样式拓展语言 | https://sass-lang.com/ |


<br><br>

**后端技术**
| 技术 | 说明 | 官网 |
|  ----  | ----  | ----  |
| Go |Go 语言 | <https://go.dev/> |
| Gin |Go (Golang) 编写的 Web 框架 | <https://github.com/gin-gonic/gin> |
| GORM | Golang 的 ORM框架 | <https://gorm.io/index.html> |
| Swagger-UI | 文档生成工具 | <https://github.com/swagger-api/swagger-ui> |

<br><br>

**容器技术**
| 技术 | 说明 | 官网 |
|  ----  | ----  | ----  |
| Docker | 容器 | https://www.docker.com/ |
| Kubernetes | 容器编排引擎 | https://kubernetes.io/ |

<br><br>

**数据库**
| 技术 | 说明 | 官网 |
|  ----  | ----  | ----  |
| MySql | 关系型数据库 | https://www.mysql.com/ |
| Redis | 非关系型数据库 | https://redis.io/ |
