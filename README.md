ji## AzuLLia - Core

![](./docs/Logo.png)

**AzuLLia-Core用于快速构建Blog内容管理接口**

### 开始！

**获取仓库源码**

```shell
git clone https://github.com/ExMikuPro/AzuLLia.git
```


开始之前修改```Service.env```文件配置

```
WEBSITE_HOST=127.0.0.1 # 网页服务地址
WEBSITE_PORT=:80       # 网页服务端口

DB_HOST=127.0.0.1      # 数据库地址
DB_PORT=27017          # 数据库端口
DB_USER=root           # 数据库用户名
DB_PASSWD=12345678     # 数据库密码
DB_POOL_SIZE=10        # 数据库连接池大小
DB_DATA_BASE=azullia   # 数据库名称

GIN_MODE=debug         # 框架运行模式

JWT_KEY = key          # JWT加密密钥(需要自行修改)
HASH_KEY = secret key  # HASH256加密密钥
```

根据需要修改完成后，进行一下步骤(任选其一部署方案即可)

---

### 两种部署方式

**Docker & Docker Compose**

**Docker部署方案：**

> 注意⚠️：已有MongoDB数据库环境的前提下可直接使用此方案，如果没有相关环境，请自行搭建或使用Docker Compose部署方式；

在文件根目录执行：

```shell
docker build -t azullia:latest .
```

进行镜像创建

运行容器：

```shell
docker run -itd -p 8000:80 -p 27017:27017 --name azullia-service azullia:latest
```

**Docker Compose部署方案：**

> 注意⚠️：Ubuntu环境下可能需要单独安装Docker Compose

> 注意⚠️：使用此部署方案会创建一个Docker虚拟网络环境和一个数据库运行容器;

```shell
docker-compose -f ./docker-compose.yml -p azullia up -d
```

---

### 检测服务是否正常启动

```shell
curl --request GET  \
     --url 'http://127.0.0.1:8000/about'
```

**返回以下json时代表服务已经很成功跑起来了！**

> 提示🔔: x.x.x 是代表目前框架运行版本

```json
{
  "code": 0,
  "msg": "success",
  "path": "/about",
  "data": {
    "version": "x.x.x"
  }
}
```

---

### 目录结构

 ```
AzuLLia
├── Dockerfile
├── README.md
├── Service.env
├── docker-compose.yml
├── docs
│   ├── Logo.png
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── service
│   ├── api.go
│   ├── database.go
│   ├── file.go
│   ├── function.go
│   ├── router.go
│   └── type.go
├── staticFile
│   ├── README.md
│   └── favicon.ico
└── uploadFile
    └── README.md

 ```

### TODO

1. [x] 编写 dockerFile
2. [ ] ~~自动化安装程序~~
3. [x] 优化数据库增删改查
4. [x] 实现基本接口功能
5. [ ] 整理数据库结构
6. [x] 编写项目接口文档
7. [x] 绘制项目Logo
8. [ ] 创建动态加载插件
9. [x] 添加JWT用户登录认证
10. [ ] 构建热更新功能
11. [x] 使用Swagger构建项目文档
12. [ ] 添加OAuth服务支持