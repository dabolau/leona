# 蕾欧娜

蕾欧娜是一个后端程序

# 创建网络

```bash
docker network create mongo-network
```

# 创建存储卷

```bash
docker volume create mongo-db
docker volume create mongo-config
```

# 启动数据库

```bash
docker run -d \
-p 27017:27017 \
--network mongo-network \
--name mongo \
-e MONGO_INITDB_ROOT_USERNAME=username \
-e MONGO_INITDB_ROOT_PASSWORD=password \
-v mongo-db:/data/db \
-v mongo-config:/data/configdb \
mongo:latest
```

# 编译程序

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o leona
```

# 构建镜像

```bash
docker build -t leona:latest .
```

# 保存镜像

```bash
docker save -o leona.tar leona:latest
```

# 载入镜像

```bash
docker load -i leona.tar
```

# 启动程序

```bash
docker run -d \
-p 8081:8081 \
--network mongo-network \
--name leona \
leona:latest
```

# 标记镜像

```bash
docker tag leona:latest dabolau/leona:latest
```

# 推送镜像到远程仓库

```bash
docker push dabolau/leona:latest
```

# 同时启动数据库和程序

编辑配置文件`docker-compose.yaml`

```bash
# 版本信息
version: '3'
# 容器服务，包含多个容器
services:
  # 容器服务
  leona:
    # 容器名称
    image: leona:latest
    # 端口[宿主机:容器]
    ports:
      - "8081:8081"
    # 容器网络
    networks:
      - mongo-network
    environment:
      - MONGO_ADDRESS=mongo:27017
      - MONGO_USERNAME=username
      - MONGO_PASSWORD=password
  # 容器服务
  mongo:
    # 容器名称
    image: mongo:latest
    # 端口[宿主机:容器]
    ports:
      - "27017:27017"
    # 容器网络
    networks:
      - mongo-network
    # 环境变量
    environment:
      - MONGO_INITDB_ROOT_USERNAME=username
      - MONGO_INITDB_ROOT_PASSWORD=password
    # 数据卷[宿主机:容器]
    volumes:
      - mongo-config:/data/configdb
      - mongo-data:/data/db
# 数据卷
volumes:
  # 命名数据卷名称与上面容器的宿主机数据卷名称一致
  mongo-config:
  mongo-data: # 容器网络
networks:
  # 命名网络名称与上面容器的网络名称一致
  mongo-network:
```

执行命令启动程序

```bash
docker-compose up -d
```