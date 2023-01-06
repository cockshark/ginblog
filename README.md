# ginblog
first gin project practice

### 初始化项目目录
```markdown
gin-blog/
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   ├── article.go
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── routers
│   ├── api
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime
```
- conf： 用于存储配置文件
- middleware： 应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers： 路由逻辑处理
- runtime： 应用运行时数据
- sql：存放建表语句