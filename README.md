## AzuLLia - Core

![](./docs/Logo.png)

**AzuLLia-Core用于快速构建Blog内容管理接口**

### 开始！

**获取仓库源码**

```shell
git clone https://github.com/ExMikuPro/AzuLLia.git
```

> 在编译之前需要在```./service/type.go```文件中配置一下框架基本配置参数:
>

找到常量```DBAddress```&```DBAddress```

```js
const DBAddress = "mongodb://username:password@host:port/database" // 数据库地址
const DBDataBase = "AzuLLia"                                       // 数据库名称
```

修改完成后进行编译

```shell
docker build -t azullia:latest .
```

### 目录结构

 ```
AzuLLia
├── Dockerfile
├── README.md
├── docker-compose.yml
├── docs
│  ├── Logo.png
│  ├── docs.go
│  ├── swagger.json
│  └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── service
│  ├── api.go
│  ├── database.go
│  ├── file.go
│  ├── function.go
│  ├── router.go
│  └── type.go
├── staticFile
│  ├── README.md
│  └── favicon.ico
└── uploadFile
    └── README.md

 ```

### 开发进度

1. [x] 编写 dockerFile
2. [ ] 自动化安装程序
3. [ ] 优化数据库增删改查
4. [ ] 实现基本接口功能
5. [ ] 整理数据库结构
6. [ ] 编写项目接口文档
7. [ ] 绘制项目Logo