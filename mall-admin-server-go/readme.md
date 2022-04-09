# mall-admin-server-go
## 项目简介
该项目实现了微信商城管理平台的后端。

微信商城管理平台前端 仓库链接[：微信商城前端管理平台仓库链接](https://github.com/TeanLee/mall-admin/tree/master/mall-admin-frontend)

## 开发环境
-   **[GoLand](https://www.jetbrains.com/go/)**

## 运行方式一
1. `git clone https://github.com/TeanLee/mall-admin.git`
2. `cd mall-admin-server-go`
3. 在 GoLand 中打开导入 clone 的项目
4. 执行 `go run main.go`

## 运行方式二
1. 在本机中安装 [Docker](https://www.docker.com/)
2. 在本文件夹下执行 `docker-compose up --build`

## 项目组织结构
````
mall-admin-server-go
├── api -- Controller 层
├── config -- MySQL、Redis、Session 的连接和初始化
├── env -- 环境变量配置（本地运行、线上运行的默认环境，在 Dockerfile 中或集群中指定运行环境）
├── middleware -- middleware 层
├── model -- 实体类层，与数据库交互
├── MySQL -- 用于 Dockerfile 定制数据库，管理平台数据初始化
├── router -- 定义项目重的后端接口路由
├── service -- service 层
````

## 技术选型

**后端技术**

| 技术 | 说明 | 官网 |
|  ----  | ----  | ----  |
| Go |Go 语言 | <https://go.dev/> |
| Gin |Go (Golang) 编写的 Web 框架 | <https://github.com/gin-gonic/gin> |
| GORM | Golang 的 ORM框架 | <https://gorm.io/index.html> |
| Swagger-UI | 文档生成工具 | <https://github.com/swagger-api/swagger-ui> |


**容器技术**

| 技术 | 说明 | 官网 |
|  ----  | ----  | ----  |
| Docker | 容器 | https://www.docker.com/ |
| Kubernetes | 容器编排引擎 | https://kubernetes.io/ |