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

修改完成后进行一下步骤(任选其一运行即可)

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